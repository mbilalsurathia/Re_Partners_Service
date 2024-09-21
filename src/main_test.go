package main

import (
	"fmt"
	"pack-calculator/src/service"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	packSizes := []int{5000, 2000, 1000, 500, 250}

	// Test orders
	orders := []int{1, 250, 251, 501, 12001}
	for _, order := range orders {
		result := service.CalculatePacks(order, packSizes)
		if result != nil {
			fmt.Printf("Order for %d items: ", order)
			for size, count := range result {
				fmt.Printf("%dx%d ", count, size)
			}
			fmt.Println()
		} else {
			fmt.Printf("Order for %d items cannot be fulfilled\n", order)
		}
	}
}
