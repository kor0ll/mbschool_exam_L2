package main

import (
	"testing"
)

// Тест функции extractKey, которая извлекает ключ из строки.
func TestExtractKey(t *testing.T) {
	tests := []struct {
		line        string
		keyColumn   int
		expectedKey string
	}{
		{"яблоко 123 груша", 1, "123"}, // Тест с числовым ключом и месяцем.
		{"яблоко Январь", 1, "Январь"}, // Тест с месяцем.
	}

	for _, test := range tests {
		key := NewLine(test.line, test.keyColumn).key
		if key != test.expectedKey {
			t.Errorf("Строка: %s, Ожидаемый ключ: %s, Получено: %s", test.line, test.expectedKey, key)
		}
	}
}

// Тест функции compareStrings, которая сравнивает строки с учетом различных флагов.
func TestCompareStrings(t *testing.T) {
	tests := []struct {
		a              string
		b              string
		numeric        bool
		reverse        bool
		expectedResult bool
	}{
		{"123", "456", true, false, true},       // Сравнение чисел
		{"яблоко", "банан", false, false, true}, // Сравнение строк
		{"яблоко", "банан", false, true, false}, // Сравнение строк в обратном порядке
	}

	for _, test := range tests {
		result := CompareStrings(test.a, test.b, test.numeric, test.reverse)
		if result != test.expectedResult {
			t.Errorf("a: %s, b: %s, Ожидаемый результат: %v, Получено: %v", test.a, test.b, test.expectedResult, result)
		}
	}
}

// Тест функции RemoveDuplicates, которая удаляет дубликаты из списка строк
func TestRemoveDuplicates(t *testing.T) {
	// Создаем список строк с дубликатами
	lines := []Line{
		{value: "яблоко", key: "яблоко"},
		{value: "банан", key: "банан"},
		{value: "яблоко", key: "яблоко"},
		{value: "вишня", key: "вишня"},
	}

	// Ожидаемый результат после удаления дубликатов
	expectedResult := []Line{
		{value: "яблоко", key: "яблоко"},
		{value: "банан", key: "банан"},
		{value: "вишня", key: "вишня"},
	}

	// Вызываем функцию удаления дубликатов
	result := DeleteDublicates(lines)

	// Проверяем, что длина результата совпадает с ожидаемой длиной
	if len(result) != len(expectedResult) {
		t.Errorf("Ожидаемая длина: %d, Получено: %d", len(expectedResult), len(result))
	}

	// Проверяем, что каждая строка в результате совпадает с ожидаемой строкой
	for i, line := range result {
		if line != expectedResult[i] {
			t.Errorf("Ожидаемый результат: %v, Получено: %v", expectedResult[i], line)
		}
	}
}
