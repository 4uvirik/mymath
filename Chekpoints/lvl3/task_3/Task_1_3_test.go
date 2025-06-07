package main

import (
	"testing"
)

func TestSumDigits(t *testing.T) {
	// Настройка тестовых данных
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"отрицательные числа", -53, 8},
		{"число 0", 0, 0},
		{"большие числа", 999999, 54},
	}
	// вызов тестируемого кода
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sumDigits(tt.input)
			// Проверка результатов
			if actual != tt.expected {
				t.Errorf("Для числа %d, получаем %d, но ожидали %d", tt.input, actual, tt.expected)
			}
		})
	}
}
