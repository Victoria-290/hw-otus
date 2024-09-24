package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "Hello, world! Hello, Go developers. Go, go, go!",
			expected: map[string]int{"hello": 2, "world": 1, "go": 4, "developers": 1},
		},
		{
			input:    "Test test TEST!",
			expected: map[string]int{"test": 3},
		},
		{
			input:    "123 123 456",
			expected: map[string]int{"123": 2, "456": 1},
		},
		{
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := countWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}
