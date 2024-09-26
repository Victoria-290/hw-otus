package main

import "testing"

// TestBinarySearch проверяет работу функции binarySearch.
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},  // Элемент найден
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 0},  // Элемент найден в начале
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},  // Элемент найден в конце
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, -1}, // Элемент не найден
		{[]int{}, 1, -1},                    // Пустой массив
	}

	for _, test := range tests {
		result := binarySearch(test.nums, test.target)
		if result != test.expected {
			t.Errorf("For nums %v and target %d, expected %d but got %d", test.nums, test.target, test.expected, result)
		}
	}
}
