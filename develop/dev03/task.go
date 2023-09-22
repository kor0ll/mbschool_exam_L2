package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура для строк из файла, key является колонкой для сортировки
type Line struct {
	value string
	key   string
}

// Конструктор для структуры Line
func NewLine(line string, key int) *Line {
	words := strings.Split(line, " ")
	return &Line{line, words[key]}
}

// Фукнция сравнивает две строки, учитывая числовое сравнение и сортировку в обратном порядке
func CompareStrings(s1 string, s2 string, byNumeric bool, reverse bool) bool {
	if byNumeric {
		num1, _ := strconv.Atoi(s1)
		num2, _ := strconv.Atoi(s2)
		if reverse {
			return num1 > num2
		} else {
			return num1 < num2
		}
	} else {
		if reverse {
			return s1 < s2
		} else {
			return s1 > s2
		}
	}
}

// Функция проверяет, является ли массив структур Lines отсортированным
func isSorted(lines []Line, reverse bool) bool {
	for i := 1; i < len(lines); i++ {
		if reverse && lines[i].key > lines[i-1].key || lines[i].key < lines[i-1].key {
			return false
		}
	}
	return true
}

// Функция удаляет повторяющиеся строки и возвращает слайс Lines с уникальными строками
func DeleteDublicates(lines []Line) []Line {
	newLines := []Line{}
	m := make(map[string]struct{})

	for _, line := range lines {
		if _, ok := m[line.value]; !ok {
			m[line.value] = struct{}{}
			newLines = append(newLines, line)
		}
	}

	return newLines
}

func main() {
	//определяем флаги утилиты
	byNumeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	uniqueStrings := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	keyColumn := flag.Int("k", 0, "Колонка для сортировки (0 по умолчанию)")
	sortedCheck := flag.Bool("c", false, "Проверка на сортировку данных")
	flag.Parse()

	//получаем путь к файлу
	inputPath := os.Args[len(os.Args)-1]

	lines := []Line{}
	//считываем строки из файла и заполняем lines
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Не удалось открыть файл: ", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, *NewLine(scanner.Text(), *keyColumn))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Не удалось прочитать файл: ", err.Error())
		return
	}

	//выполняем проверку на числовые значения, если указана сортировка по числам
	if *byNumeric {
		for _, v := range lines {
			_, err := strconv.Atoi(v.key)
			if err != nil {
				fmt.Println("Не удалось определить номер колонки: ", err.Error())
				return
			}
		}
	}
	//выполняем проверку на отсортированность
	if *sortedCheck {
		if isSorted(lines, *reverse) {
			fmt.Println("Файл отсортирован!")
		} else {
			fmt.Println("Файл не отсортирован!")
		}
		return
	}
	//удаляем повторяющиеся строки если указан нужный флаг
	if *uniqueStrings {
		lines = DeleteDublicates(lines)
	}

	//сортируем подготовленные данные
	sort.Slice(lines, func(i, j int) bool {
		return CompareStrings(lines[i].key, lines[j].key, *byNumeric, *reverse)
	})

	// создаем и записываем отсортированные данные в файл
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for _, line := range lines {
		_, err := outputFile.WriteString(line.value + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
