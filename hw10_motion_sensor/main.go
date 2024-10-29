package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// Генерирует данные с сенсора и передаёт их в канал.
func sensorDataGenerator(sensorChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeout := time.After(1 * time.Minute)

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
}

// Обрабатывает данные: вычисляет среднее значение каждых 10 значений.
func processSensorData(sensorChan <-chan int, processedDataChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make([]int, 0, 10)

	for data := range sensorChan {
		buffer = append(buffer, data)

		// Если накоплено 10 значений, вычисляем среднее и отправляем в канал обработанных данных.
		if len(buffer) == 10 {
			sum := 0
			for _, value := range buffer {
				sum += value
			}
			average := float64(sum) / float64(len(buffer))
			processedDataChan <- average
			buffer = buffer[:0] // очищаем буфер.
		}
	}
	close(processedDataChan)
}

func main() {
	sensorChan := make(chan int)
	processedDataChan := make(chan float64)
	var wg sync.WaitGroup

	// Запуск горутины для генерации данных.
	wg.Add(1)
	go sensorDataGenerator(sensorChan, &wg)

	// Запуск горутины для обработки данных.
	wg.Add(1)
	go processSensorData(sensorChan, processedDataChan, &wg)

	// Получение обработанных данных в главной горутине и вывод.
	go func() {
		for avg := range processedDataChan {
			fmt.Printf("Среднее значение: %.2f\n", avg)
		}
	}()

	wg.Wait()
}
