package main

import (
	"testing"
)

// Test generateRandomElements function
func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name           string
		size           int
		expectedLength int
	}{
		{
			name:           "размер 0",
			size:           0,
			expectedLength: 0,
		},
		{
			name:           "отрицательный размер",
			size:           -5,
			expectedLength: 0,
		},
		{
			name:           "размер 1",
			size:           1,
			expectedLength: 1,
		},
		{
			name:           "размер 100",
			size:           100,
			expectedLength: 100,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateRandomElements(tc.size)

			if len(result) != tc.expectedLength {
				t.Errorf("Ожидался срез длиной %d, получен срез длиной %d", tc.expectedLength, len(result))
			}
		})
	}
}

// Test maximum function
func TestMaximum(t *testing.T) {
	testCases := []struct {
		name       string
		inputSlice []int
		expected   int
	}{
		{
			name:       "пустой срез",
			inputSlice: []int{},
			expected:   0,
		},
		{
			name:       "срез с одним элементом",
			inputSlice: []int{8},
			expected:   8,
		},
		{
			name:       "срез положительных значений",
			inputSlice: []int{1, 4, 2, 8, 6},
			expected:   8,
		},
		{
			name:       "срез отрицательных значений",
			inputSlice: []int{-1, -4, -2, -8, -6},
			expected:   -1,
		},
		{
			name:       "срез с отрицательными и положительными значениями",
			inputSlice: []int{-1, 4. - 2, 8, -6},
			expected:   8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximum(tc.inputSlice)
			if result != tc.expected {
				t.Errorf("Ожидалось %d для среза %v, получено %d", tc.expected, tc.inputSlice, result)
			}
		})
	}
}

// Test maxChunks function
func TestMaxChunks(t *testing.T) {
	testCases := []struct {
		name       string
		inputSlice []int
		expected   int
	}{
		{
			name:       "пустой срез",
			inputSlice: []int{},
			expected:   0,
		},
		{
			name:       "срез с одним элементом",
			inputSlice: []int{8},
			expected:   8,
		},
		{
			name:       "срез положительных значений",
			inputSlice: []int{1, 4, 2, 8, 6},
			expected:   8,
		},
		{
			name:       "срез отрицательных значений",
			inputSlice: []int{-1, -4, -2, -8, -6},
			expected:   -1,
		},
		{
			name:       "срез с отрицательными и положительными значениями",
			inputSlice: []int{-1, 4. - 2, 8, -6},
			expected:   8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximum(tc.inputSlice)
			if result != tc.expected {
				t.Errorf("Ожидалось %d для среза %v, получено %d", tc.expected, tc.inputSlice, result)
			}
		})
	}
}
