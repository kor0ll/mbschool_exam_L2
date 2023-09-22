package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// настраиваем флаги и значения по умолчанию
	host := flag.String("host", "", "Хост для подключения")
	port := flag.String("port", "23", "Порт для подключения")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут для подключения")

	flag.Parse()

	// собираем адрес для подключения
	address := fmt.Sprintf("%s:%s", *host, *port)

	// устанавливаем таймаут для подключения
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к", address)

	// запускаем горутину для чтения данных из сокета и вывода на STDOUT
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Соединение разорвано")
				os.Exit(0)
			}
			fmt.Print(string(buffer[:n]))
		}
	}()

	fmt.Println("Начинаем передавать ввод в сокет и получать вывод...")

	// читаем ввод пользователя с консоли и отправляем в сокет
	buffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			os.Exit(1)
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Ошибка отправки данных:", err)
			os.Exit(1)
		}
	}
}
