package main

import (
	"fmt"

	"github.com/Victoria-290/hw-otus/hw06_testing/pkg/book"
	"github.com/Victoria-290/hw-otus/hw06_testing/pkg/chessboard"
	"github.com/Victoria-290/hw-otus/hw06_testing/pkg/shape"
)

func main() {
	// Пример работы с книгами
	b1 := book.NewBook(1, "Go", "John Doe", 2021, 300, 4.5)
	b2 := book.NewBook(2, "Python", "Jane Doe", 2019, 400, 4.7)

	fmt.Println("Book 1 year:", b1.GetYear())
	fmt.Println("Book 2 year:", b2.GetYear())

	// Пример работы с фигурами
	circle := shape.Circle{Radius: 10}
	fmt.Println("Circle area:", circle.Area())

	// Пример работы с доской
	var size int

	fmt.Print("Введите размер доски: ")
	fmt.Scanf("%d", &size)

	err := chessboard.GenerateChessBoard(size)
	if err != nil {
		fmt.Println(err)
	}
}
