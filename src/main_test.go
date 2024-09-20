package main

import (
	"fmt"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	items := 501
	expected := map[int]int{500: 1, 250: 1}

	result := calculatePacks(items, packSizes)

	for k, v := range expected {
		fmt.Println(result)
		if result[k] != v {
			t.Errorf("Expected %d packs of size %d, but got %d", v, k, result[k])
		}
	}
}
