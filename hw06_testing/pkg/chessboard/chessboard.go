package chessboard

import "fmt"

// GenerateChessBoard генерирует шахматную доску заданного размера.
func GenerateChessBoard(size int) error {
	if size <= 0 {
		return fmt.Errorf("ошибка. Размер доски должен быть положительным числом")
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return nil
}
