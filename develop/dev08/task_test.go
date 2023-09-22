package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	dir, err := printWorkingDirectory()
	if err != nil {
		t.Fatalf("Ошибка при получении текущей директории: %v", err)
	}
	t.Logf("Текущая директория: %s", dir)
}

func TestEcho(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
		errMsg   string
	}{
		{[]string{"echo", "Привет, мир!"}, "Привет, мир!", ""},
		{[]string{"echo"}, "", "использование: echo <текст>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			output, err := echo(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}

			if output != test.expected {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, output)
			}
		})
	}
}

func TestListProcesses(t *testing.T) {
	output, err := listProcesses()
	if err != nil {
		t.Fatalf("Ошибка при выполнении команды 'ps': %v", err)
	}
	t.Logf("Вывод 'ps':\n%s", output)
}

func TestKillProcess(t *testing.T) {
	// в pid нужно указать имя процесса, который хотим убить
	pid := "someprocess"

	tests := []struct {
		args   []string
		errMsg string
	}{
		{[]string{"kill", pid}, ""},
		{[]string{"kill"}, "Использование: kill <PID>"},
		{[]string{"kill", "invalid"}, "Ошибка: неверный формат PID"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			err := killProcess(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
		})
	}
}
