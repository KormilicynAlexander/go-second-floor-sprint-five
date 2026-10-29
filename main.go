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

// generateRandomElements returns a slice of random positive integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum returns the largest value, or 0 for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks finds the maximum using CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			if len(chunk) == 0 {
				// Пустая часть не должна влиять на результат для отрицательных чисел.
				chunkMaxima[index] = data[0]
				return
			}
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
