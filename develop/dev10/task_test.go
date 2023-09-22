package main

import (
	"net"
	"os"
	"testing"
	"time"
)

func startTestServer(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go handleTestConnection(conn)
		}
	}()

	return listener, nil
}

func handleTestConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		conn.Write(buf[:n])
	}
}

func TestMain(m *testing.M) {
	serverAddr := "localhost:12345"
	listener, err := startTestServer(serverAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go func() {
		os.Exit(m.Run())
	}()

	time.Sleep(1 * time.Second) // Даем серверу немного времени для запуска
	os.Exit(0)
}

func TestTelnetClient(t *testing.T) {
	input := "Test Input\n"
	expectedOutput := input

	output := runTelnetClient("localhost:12345", input)

	if output != expectedOutput {
		t.Errorf("Ожидался вывод: %s, получен вывод: %s", expectedOutput, output)
	}
}

func runTelnetClient(serverAddr, input string) string {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	_, err = conn.Write([]byte(input))
	if err != nil {
		return err.Error()
	}

	buffer := make([]byte, len(input))
	_, err = conn.Read(buffer)
	if err != nil {
		return err.Error()
	}

	return string(buffer)
}
