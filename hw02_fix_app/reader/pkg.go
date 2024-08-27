package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Victoria-290/hw-otus/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close() // Закрыла файл после завершения функции

	fileContent, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data []types.Employee

	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if limit > 0 && limit < len(data) {
		data = data[:limit] // Обрезала до нужного количества элементов
	}

	return data, nil
}
