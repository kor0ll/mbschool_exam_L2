package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getHash(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v)
	}
	return sum
}

func FindAnogram(sl []string) *map[string][]string {
	m := make(map[string][]string)

	for _, val := range sl {
		val = strings.ToLower(val)
		sum := 0
		currentWord := ""
		sum = getHash(val)

		for i := range m {
			if getHash(i) == sum {
				currentWord = i
			}
		}
		if currentWord != "" {
			m[currentWord] = append(m[currentWord], val)
		} else {
			m[val] = []string{val}
		}
	}
	for i := range m {
		if len(m[i]) == 1 {
			delete(m, i)
		}
	}

	return &m
}

func main() {
	ar := []string{"пятка", "тяпка", "лимон", "милон", "пятак", "листок", "Слиток", "слизень"}

	fmt.Println(FindAnogram(ar))
}
