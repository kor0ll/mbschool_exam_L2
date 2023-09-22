package main

import (
	"fmt"
	"sort"
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

// функция возвращающая отсортированную по буквам строку
func GetSortString(s string) string {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	result := strings.Join(sl, "")

	return result
}

// функция для нахождения анаграмм в массиве
func FindAnagramm(sl []string) *map[string][]string {
	m := make(map[string][]string)
	for _, word := range sl {
		//сортируем строку по буквам и распределяем по множествам
		word = strings.ToLower(word)
		sortedWord := GetSortString(word)

		if _, ok := m[sortedWord]; !ok {
			m[sortedWord] = []string{word}
		} else {
			m[sortedWord] = append(m[sortedWord], word)
		}
	}
	//удаляем из мапы множества из одного элемента
	for i := range m {
		if len(m[i]) == 1 {
			delete(m, i)
		}
	}
	//преобразуем ключи в нормальный вид (первое встретившееся слово) и сортируем множества по возрастанию
	resultMap := make(map[string][]string, len(m))
	for _, arr := range m {
		sortedValue := make([]string, len(arr))
		copy(sortedValue, arr)
		sort.Slice(sortedValue, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		resultMap[arr[0]] = sortedValue
	}
	return &resultMap
}

func main() {
	ar := []string{"пятка", "тяпка", "лимон", "милон", "пятак", "листок", "Слиток", "слизень"}

	result := FindAnagramm(ar)

	for index, value := range *result {
		fmt.Println("Index ", index, ": ", value)
	}

}
