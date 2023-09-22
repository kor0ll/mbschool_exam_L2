package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, _ := Unpack(tc.input)
			if result != tc.expected {
				t.Errorf("Для входных данных '%s' получен результат '%s', ожидался результат '%s'.", tc.input, result, tc.expected)
			}
		})
	}
}
