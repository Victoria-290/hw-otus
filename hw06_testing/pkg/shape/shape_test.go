package shape

import (
	"errors"
	"testing"
)

// Табличные тесты для функции calculateArea.
func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr error
	}{
		{"Circle", Circle{Radius: 5}, 78.53981633974483, nil},
		{"Rectangle", Rectangle{Width: 10, Height: 5}, 50, nil},
		{"Triangle", Triangle{Base: 8, Height: 6}, 24, nil},
		{"Not a Shape", "not a shape", 0, errors.New("переданный объект не является фигурой")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateArea(tt.input)
			if err != nil && tt.wantErr == nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.wantErr != nil && err == nil {
				t.Errorf("expected error but got none")
			}
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
