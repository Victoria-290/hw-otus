package chessboard

import (
	"testing"
)

// Табличные тесты для функции GenerateChessBoard.
func TestGenerateChessBoard(t *testing.T) {
	tests := []struct {
		size        int
		expectError bool
	}{
		{size: 5, expectError: false},
		{size: -1, expectError: true},
		{size: 0, expectError: true},
	}

	for _, test := range tests {
		err := GenerateChessBoard(test.size)
		if test.expectError && err == nil {
			t.Errorf("Expected error but got none, size: %d", test.size)
		} else if !test.expectError && err != nil {
			t.Errorf("Did not expect error but got: %v, size: %d", err, test.size)
		}
	}
}
