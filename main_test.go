package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	empty := generateRandomElements(0)
	if empty != nil {
		t.Errorf("For size 0, should be returned nil")
	}
	size10 := generateRandomElements(10)
	if len(size10) != 10 {
		t.Errorf("Expected length of 88, received: %d", len(size10))
	}
}
