package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция загружает сайт и сохраняет его в index.html
func DownloadSite(url string) error {
	// открываем HTTP-запрос к указанному URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// создаем файл для сохранения сайта
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// копируем содержимое ответа в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Сайт успешно загружен в файл index.html!")
	return nil
}

func main() {
	// проверяем, был ли передан URL сайта
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run task.go <URL>")
		return
	}

	// получаем URL
	url := os.Args[1]

	// вызываем функцию для загрузки сайта
	err := DownloadSite(url)
	if err != nil {
		fmt.Println("Ошибка при загрузке сайта:", err.Error())
	}
}
