package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestGrepUtility(t *testing.T) {
	// Подготовка временного каталога и файлов с тестовыми данными
	inputFile := "input.txt"
	expectedOutputFile := "output.txt"
	// создание файлов с тестовыми данными
	err := os.WriteFile(inputFile, []byte("В данном простом примере\nМы будем искать слово шаблон\nШаблонов может быть несколько\nОн может иметь разные суффиксы\nА также ШАБЛОН может быть в разных регистрах"), 0644)
	if err != nil {
		t.Log(err.Error())
	}
	err = os.WriteFile(expectedOutputFile, []byte("Мы будем искать слово шаблон\nШаблонов может быть несколько\nА также ШАБЛОН может быть в разных регистрах\n"), 0644)
	if err != nil {
		t.Log(err.Error())
	}
	// запуск утилиты с заданными флагами и входными данными
	cmd := exec.Command("go", "run", "task.go", "шаблон", "-i", inputFile)

	output, _ := cmd.CombinedOutput()
	t.Logf("Результат выполнения утилиты:\n%s", string(output))

	// Сравнение фактического вывода с ожидаемым
	expectedOutput, _ := os.ReadFile(expectedOutputFile)

	if string(expectedOutput) != string(output) {
		t.Log("Некорректная работа утилиты")
	}
}
