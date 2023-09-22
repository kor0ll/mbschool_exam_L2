package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция конструктор создает сканер для чтения строк из r с разделителем delimiterRune.
func NewLineScanner(r io.Reader, delimiterRune rune) *Scanner {
	return &Scanner{
		scanner:       bufio.NewScanner(r),
		delimiterRune: delimiterRune,
	}
}

// Структура для чтения строк с разделителем.
type Scanner struct {
	scanner       *bufio.Scanner
	delimiterRune rune
}

// Функция считывает следующую строку с разделителем.
func (s *Scanner) Scan() bool {
	s.scanner.Split(s.splitFunc)
	return s.scanner.Scan()
}

// Функция возвращает текущую строку после сканирования.
func (s *Scanner) Text() string {
	return s.scanner.Text()
}

// splitFunc является функцией разделения для сканера.
func (s *Scanner) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexRune(data, s.delimiterRune); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func main() {
	// настраиваем флаги
	fields := flag.String("f", "", "выбрать поля (колонки), разделенные запятой")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// получение выбранных полей (колонок)
	selectedFields := make(map[int]bool)
	if *fields != "" {
		fieldList := strings.Split(*fields, ",")
		for _, fieldStr := range fieldList {
			fieldIdx, err := strconv.Atoi(fieldStr)
			if err != nil {
				fmt.Printf("Ошибка преобразования поля %s в число: %v\n", fieldStr, err)
				os.Exit(1)
			}
			selectedFields[fieldIdx] = true
		}
	}

	// обработка входных данных
	var reader io.Reader
	if flag.NArg() == 0 {
		reader = os.Stdin
	} else {
		filename := flag.Arg(0)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		fields := strings.Split(line, *delimiter)

		// выбор полей (колонок) и вывод
		var outputFields []string
		for idx, field := range fields {
			if selectedFields[idx+1] || len(selectedFields) == 0 {
				outputFields = append(outputFields, field)
			}
		}
		if len(outputFields) > 0 {
			fmt.Println(strings.Join(outputFields, *delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при сканировании: %v\n", err)
		os.Exit(1)
	}
}
