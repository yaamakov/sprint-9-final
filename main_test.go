package main

import (
	"testing"
)

// Test generateRandomElements function
func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез для размера 0, получен срез длиной %d", len(result))
	}

	result = generateRandomElements(-5)
	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез для отрицательного размера, получен срез с длиной %d", len(result))
	}

	size := 100
	result = generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Ожидался срез длиной %d, получен срез длиной %d", size, len(result))
	}

	for i, val := range result {
		if val <= 0 {
			t.Errorf("Ожидалось положительное значение по индексу %d, получено %d", i, val)
		}
	}
}

// Test maximum function
func TestMaximum(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("Ожидалось 0 для пустого среза, получено %d", result)
	}

	result = maximum([]int{5})
	if result != 5 {
		t.Errorf("Ожидалось 5 для среза с одним элементом, получено %d", result)
	}

	result = maximum([]int{1, 4, 2, 8, 6})
	if result != 8 {
		t.Errorf("Ожидалось 8 для среза [1, 4, 2, 8, 6], получено %d", result)
	}

	result = maximum([]int{-1, 4, -2, 8, -6})
	if result != 8 {
		t.Errorf("Ожидалось 8 для среза [-1, 4, -2, 8, -6], получено %d", result)
	}

	result = maximum([]int{-1, -4, -2, -8, -6})
	if result != -1 {
		t.Errorf("Ожидалось -1 для среза [-1, -4, -2, -8, -6], получено %d", result)
	}
}

// Test maxChunks function
func TestMaxChunks(t *testing.T) {
	result := maxChunks([]int{})
	if result != 0 {
		t.Errorf("Ожидалось 0 для пустого среза, получено %d", result)
	}

	result = maxChunks([]int{5})
	if result != 5 {
		t.Errorf("Ожидалось 5 для среза с одним элементом, получено %d", result)
	}

	result = maxChunks([]int{1, 4, 2, 8, 6})
	if result != 8 {
		t.Errorf("Ожидалось 8 для среза [1, 4, 2, 8, 6], получено %d", result)
	}

	result = maxChunks([]int{-1, 4, -2, 8, -6})
	if result != 8 {
		t.Errorf("Ожидалось 8 для среза [-1, 4, -2, 8, -6], получено %d", result)
	}

	result = maxChunks([]int{-1, -4, -2, -8, -6})
	if result != -1 {
		t.Errorf("Ожидалось -1 для среза [-1, -4, -2, -8, -6], получено %d", result)
	}
}
