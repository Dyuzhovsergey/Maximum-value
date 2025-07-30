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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Intn(size * 10) // Генерируем числа до size*10 для разнообразия
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	size := len(data)
	chunkSize := size / CHUNKS
	if chunkSize == 0 {
		chunkSize = 1
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	maxValues := make([]int, 0, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = size // Последний чанк может быть больше
		}

		if start >= size {
			break
		}

		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			localMax := maximum(chunk)

			mu.Lock()
			maxValues = append(maxValues, localMax)
			mu.Unlock()
		}(data[start:end])
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел...\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("\nИщем максимальное значение в %d потоков...\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
