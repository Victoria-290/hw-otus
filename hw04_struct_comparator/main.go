package main

import "fmt"

// Структура Book с неэкспортируемыми полями.
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Методы для получения полей.
func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float64 {
	return b.rate
}

// Методы для установки полей.
func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

// Перечисление для выбора режима сравнения.
type CompareField int

const (
	ByYear CompareField = iota
	BySize
	ByRate
)

// Структура для сравнения книг.
type BookComparator struct {
	field CompareField
}

// Конструктор для создания нового BookComparator.
func NewBookComparator(field CompareField) *BookComparator {
	return &BookComparator{field: field}
}

// Метод для сравнения двух книг.
func (bc *BookComparator) Compare(book1, book2 *Book) bool {
	switch bc.field {
	case ByYear:
		return book1.Year() > book2.Year()
	case BySize:
		return book1.Size() > book2.Size()
	case ByRate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}

func main() {
	book1 := Book{}
	book1.SetID(1)
	book1.SetTitle("Анна Каренина")
	book1.SetAuthor("Лев Толстой")
	book1.SetYear(1913)
	book1.SetSize(803)
	book1.SetRate(4.6)

	book2 := Book{}
	book2.SetID(2)
	book2.SetTitle("Мастер и Маргарита")
	book2.SetAuthor("Михаил Булгаков")
	book2.SetYear(1967)
	book2.SetSize(461)
	book2.SetRate(4.8)

	// Сравнение по году.
	comparator := NewBookComparator(ByYear)
	fmt.Printf("Book1 > Book2 by year: %v\n", comparator.Compare(book1, book2))

	// Сравнение по размеру.
	comparator = NewBookComparator(BySize)
	fmt.Printf("Book1 > Book2 by size: %v\n", comparator.Compare(book1, book2))

	// Сравнение по рейтингу.
	comparator = NewBookComparator(ByRate)
	fmt.Printf("Book1 > Book2 by rate: %v\n", comparator.Compare(book1, book2))
}
