package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	empty := generateRandomElements(0)
	if len(empty) != 0 {
		t.Errorf("For size 0, should be returned empty slise")
	}
	size10 := generateRandomElements(10)
	if len(size10) != 10 {
		t.Errorf("Expected length of 10, received: %d", len(size10))
	}

	for _, v := range size10 {
		if v < 0 || v >= 100 {
			t.Errorf("Value %d out of range [0, 100)", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	if max := maximum([]int{}); max != 0 {
		t.Errorf("For empty slice, was expected 0, received %d", max)
	}

	if max := maximum([]int{4}); max != 4 {
		t.Errorf("Expected 4, received %d", max)
	}

	if max := maximum([]int{1, 8, 2, 6, 4}); max != 8 {
		t.Errorf("Expected 8, received %d", max)
	}

	if max := maximum([]int{7, 7, 7, 7}); max != 7 {
		t.Errorf("Expected 7, received %d", max)
	}

	if max := maximum([]int{-1, -4, -3}); max != -1 {
		t.Errorf("Expected -1, received %d", max)
	}
}
func TestMaxChunks(t *testing.T) {
	if max := maxChunks([]int{}); max != 0 {
		t.Errorf("For empty slice, was expected 0, received %d", max)
	}

	if max := maxChunks([]int{1, 2, 3}); max != 3 {
		t.Errorf("Expected 3, received %d", max)
	}

	testData := make([]int, CHUNKS)
	for i := range testData {
		testData[i] = i + 1
	}
	if max := maxChunks(testData); max != CHUNKS {
		t.Errorf("Expected %d, received %d", CHUNKS, max)
	}

	largeData := make([]int, 100)
	for i := range largeData {
		largeData[i] = i + 1
	}
	if max := maxChunks(largeData); max != 100 {
		t.Errorf("Expected 100, received %d", max)
	}
}
