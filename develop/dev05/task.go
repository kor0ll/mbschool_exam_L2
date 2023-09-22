package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура для строк из файла, key является номером строки
type Line struct {
	value string
	key   int
}

// Функция удаляет повторяющиеся строки и возвращает слайс Lines с уникальными строками
func DeleteDublicates(lines []Line) []Line {
	newLines := []Line{}
	m := make(map[int]struct{})

	for _, line := range lines {
		if _, ok := m[line.key]; !ok {
			m[line.key] = struct{}{}
			newLines = append(newLines, line)
		}
	}
	return newLines
}

func main() {
	// настраиваем флаги
	AfterFlag := flag.Int("A", 0, "Печатать +N строк после совпадения")
	BeforeFlag := flag.Int("B", 0, "Печатать +N строк до совпадения")
	ContextFlag := flag.Int("C", 0, "Печатать +N строк вокруг совпадения")
	countStringsFlag := flag.Bool("c", false, "Вывести количество найденных строк")
	ignoreCaseFlag := flag.Bool("i", false, "Игнорировать регистр при поиске")
	invertSearchFlag := flag.Bool("v", false, "Вместо совпадения исключать найденные строки")
	printLineNumFlag := flag.Bool("n", false, "Печатать номер строки, в которой найден паттерн")

	flag.Parse()

	substr := os.Args[1]

	//получаем путь к файлу
	inputPath := os.Args[len(os.Args)-1]

	//открываем файл
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Не удалось открыть файл: ", err.Error())
		return
	}
	defer file.Close()

	//слайс содержащий все строки, нужен для флагов A B С
	allLines := make(map[int]Line)
	//слайс с обработанными данными
	exportLines := []Line{}
	//текущий номер строки, нужен при -n
	lineNumber := 0

	//читаем информацию из файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber++
		line := Line{scanner.Text(), lineNumber}
		allLines[lineNumber] = line

		//при флаге -i переводим строку и шаблон в нижний регистр
		if *ignoreCaseFlag {
			lowerLine := strings.ToLower(line.value)
			lowerSubstr := strings.ToLower(substr)

			//проверяем входит ли шаблон в строку с учетом флага -v
			if (strings.Contains(lowerLine, lowerSubstr) && !*invertSearchFlag) || (!strings.Contains(lowerLine, lowerSubstr) && *invertSearchFlag) {
				exportLines = append(exportLines, line)
			}
		} else {
			//если флаг -i отсутствует, проверяем с учетом регистра
			if (strings.Contains(line.value, substr) && !*invertSearchFlag) || (!strings.Contains(line.value, substr) && *invertSearchFlag) {
				exportLines = append(exportLines, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Не удалось прочитать файл: ", err.Error())
		return
	}

	//добавление строк при выставленном -A
	if *AfterFlag != 0 {
		result := []Line{}
		for _, line := range exportLines {
			result = append(result, line)
			for i := line.key + 1; i <= line.key+*AfterFlag; i++ {
				if val, ok := allLines[i]; ok {
					result = append(result, val)
				}
			}
		}
		exportLines = result
	}

	//добавление строк при выставленном -B
	if *BeforeFlag != 0 {
		result := []Line{}
		for _, line := range exportLines {
			for i := line.key - *BeforeFlag; i < line.key; i++ {
				if val, ok := allLines[i]; ok {
					result = append(result, val)
				}
			}
			result = append(result, line)
		}
		exportLines = result
	}

	//добавление строк при выставленном -C
	if *ContextFlag != 0 {
		result := []Line{}
		for _, line := range exportLines {
			for i := line.key - *ContextFlag; i <= line.key+*ContextFlag; i++ {
				if val, ok := allLines[i]; ok {
					result = append(result, val)
				}
			}
		}
		exportLines = result
	}

	//удаляем копии, которые могли возникнуть в слайсе при флагах -A -B или -C
	exportLines = DeleteDublicates(exportLines)

	//выводим результат с учетом флага -n
	for _, line := range exportLines {
		if *printLineNumFlag {
			fmt.Printf("%d: %s \n", line.key, line.value)
		} else {
			fmt.Println(line.value)
		}

	}
	//при выставленном флаге -с выводим еще и количество строк
	if *countStringsFlag {
		fmt.Println("Количество найденных строк:", len(exportLines))
	}
}
