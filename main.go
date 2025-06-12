package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements функция генерирует числа и записывает в созданный слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1000000) + 1
	}
	return result
}

// maximum находит максимальное число в слайсе, который передаётся в функцию, и возвращает его.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks принимает слайс и возвращает максимальное значение.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	chunkSize := len(data) / CHUNKS

	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(chunkIndex int) {
			defer wg.Done()

			startIndex := chunkIndex * chunkSize
			endIndex := startIndex + chunkSize

			if chunkIndex == CHUNKS-1 {
				endIndex = len(data)
			}

			chunkMax := maximum(data[startIndex:endIndex])
			for j := startIndex + 1; j < endIndex; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}

			maxValues[chunkIndex] = chunkMax
		}(i)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	max := maximum(data)
	elapsed := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
