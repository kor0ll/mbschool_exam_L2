package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(s string) (string, error) {

	if _, err := strconv.Atoi(s); err == nil {
		return "", errors.New("некорректная строка")
	}
	var prev rune
	var result strings.Builder
	var escaped bool

	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			num := int(char - '0')
			repeated := strings.Repeat(string(prev), num-1)
			result.WriteString(repeated)
		} else {
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				result.WriteRune(char)
			}
			prev = char
		}
	}
	return result.String(), nil

}

func main() {
	s := "a4b2cd\\\\5"
	fmt.Println(Unpack(s))
}
