package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	// тестовые случаи для функции findAnagrams
	testCases := []struct {
		input    []string             // входные данные: массив слов
		expected *map[string][]string // ожидаемый результат: мапа с анаграммами
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: &map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			input: []string{"абв", "вба", "где", "едг"},
			expected: &map[string][]string{
				"абв": {"абв", "вба"},
				"где": {"где", "едг"},
			},
		},
		{
			input:    []string{"апельсин", "банан", "вишня"},
			expected: &map[string][]string{}, // нет анаграмм
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := FindAnagramm(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Ожидается %v, но получено %v", tc.expected, result)
			}
		})
	}
}
