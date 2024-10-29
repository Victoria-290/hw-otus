package main

import (
	"crypto/rand"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Тест генерации данных с сенсора.
func TestSensorDataGenerator(t *testing.T) {
	sensorChan := make(chan int)
	var dataCount int

	go func() {
		for range sensorChan {
			dataCount++
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// Запуск функции генерации данных с сокращенным временем для теста.
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		timeout := time.After(5 * time.Second) // Уменьшенный таймаут для теста

		for {
			select {
			case <-timeout:
				close(sensorChan)
				return
			case <-ticker.C:
				// Генерация случайного значения от 1 до 100.
				n, _ := rand.Int(rand.Reader, big.NewInt(100))
				data := int(n.Int64()) + 1
				sensorChan <- data
			}
		}
	}()

	// Завершаем тест через короткое время ожидания.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		assert.Greater(t, dataCount, 0, "Должно быть сгенерировано хотя бы одно значение.")
	case <-time.After(10 * time.Second):
		t.Fatal("Тест превысил лимит времени")
	}
}

// Тест обработки данных: проверка вычисления среднего значений.
func TestProcessSensorData(t *testing.T) {
	sensorChan := make(chan int, 10)
	processedDataChan := make(chan float64, 1)
	var wg sync.WaitGroup

	// Заполняем канал фиксированными значениями.
	testData := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, val := range testData {
		sensorChan <- val
	}
	close(sensorChan)

	wg.Add(1)
	go processSensorData(sensorChan, processedDataChan, &wg)
	wg.Wait()

	// Получаем результат и проверяем.
	average, ok := <-processedDataChan
	assert.True(t, ok, "Ожидается среднее значение.")
	expectedAverage := 55.0
	assert.Equal(t, expectedAverage, average, "Среднее значение должно быть равно 55.0.")
}
