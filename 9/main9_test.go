package main

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	// Создаем каналы
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	// Запускаем конвейер в отдельной горутине
	go Pipeline(inputChan, outputChan)

	// Записываем числа в inputChan
	go func() {
		defer close(inputChan)
		inputChan <- 2
		inputChan <- 3
	}()

	// Ожидаемые результаты
	expected := []float64{8, 27} // 2^3 = 8, 3^3 = 27
	index := 0

	// Читаем результаты из outputChan и проверяем их
	for result := range outputChan {
		if result != expected[index] {
			t.Errorf("Expected %.2f, got %.2f", expected[index], result)
		}
		index++
	}

	// Проверяем, что все результаты были получены
	if index != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), index)
	}
}

func TestPipeline_EmptyInput(t *testing.T) {
	// Создаем каналы
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	// Запускаем конвейер в отдельной горутине
	go Pipeline(inputChan, outputChan)

	// Закрываем inputChan без отправки данных
	close(inputChan)

	// Проверяем, что outputChan тоже закрыт и не содержит данных
	for result := range outputChan {
		t.Errorf("Expected no results, got %.2f", result)
	}
}
