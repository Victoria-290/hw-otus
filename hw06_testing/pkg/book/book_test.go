package book

import "testing"

func TestNewBook(t *testing.T) {
	tests := []struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}{
		{1, "Анна Каренина", "Лев Толстой", 1913, 803, 4.6},
		{2, "Мастер и Маргарита", "Михаил Булгаков", 1967, 461, 4.8},
	}

	for _, tt := range tests {
		b := NewBook(tt.id, tt.title, tt.author, tt.year, tt.size, tt.rate)
		if b.GetYear() != tt.year {
			t.Errorf("expected year %d, got %d", tt.year, b.GetYear())
		}
		if b.GetSize() != tt.size {
			t.Errorf("expected size %d, got %d", tt.size, b.GetSize())
		}
		if b.GetRate() != tt.rate {
			t.Errorf("expected rate %f, got %f", tt.rate, b.GetRate())
		}
	}
}
