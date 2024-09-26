package main

import "fmt"

// binarySearch выполняет двоичный поиск элемента target в отсортированном массиве nums.
// Возвращает индекс найденного элемента или -1, если элемент не найден.
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// Если элемент найден
		if nums[mid] == target {
			return mid
		}

		// Если искомый элемент больше среднего, игнорируем левую часть
		if nums[mid] < target {
			left = mid + 1
		} else {
			// Если искомый элемент меньше среднего, игнорируем правую часть
			right = mid - 1
		}
	}
	// Элемент не найден
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 5

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден.\n", target)
	}
}
