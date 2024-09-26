package main

import (
	"fmt"
	"strings"
	"unicode"
)

// countWords принимает строку и возвращает карту с количеством упоминаний каждого слова.
func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	// Раздлим текста на слова по пробелам.
	words := strings.Fields(text)
	for _, word := range words {
		// Очистим слова от пунктуации.
		cleanedWord := cleanWord(word)
		// Приведем слова к нижнему регистру для учёта разных регистров.
		cleanedWord = strings.ToLower(cleanedWord)
		// Подсчёт.
		wordCount[cleanedWord]++
	}
	return wordCount
}

// cleanWord удалим пунктуацию с начала и конца слова.
func cleanWord(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func main() {
	text := "Hello, world! Hello, Go developers. Go, go, go!"
	fmt.Println(countWords(text))
}
