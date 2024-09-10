package main

import (
	"errors"
	"fmt"
	"math"
)

// Интерфейс Shape + метод Area для вычисления площади.
type Shape interface {
	Area() float64
}

// Структура Круг.
type Circle struct {
	Radius float64
}

// Реализация метода интерфейса Area для круга.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Структура прямоугольник.
type Rectangle struct {
	Width, Height float64
}

// Реализация метода интерфейса Area для прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура треугольник.
type Triangle struct {
	Base, Height float64
}

// Реализация метода интерфейса Area для прямоугольника.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Функция, которая вычисляет площадь, если объект реализует интерфейс Shape.
func calculateArea(s any) (float64, error) {
	// Проверяем, реализует ли объект интерфейс Shape.
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	// Возвращаем площадь фигуры.
	return shape.Area(), nil
}

// Использование интерфейса для различных фигур.
func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}
	areas := []any{circle, rectangle, triangle, "not a shape"}

	for _, obj := range areas {
		area, err := calculateArea(obj)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Площадь: %.2f\n", area)
		}
	}
}
