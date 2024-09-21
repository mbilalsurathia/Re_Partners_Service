package service

import "sort"

// CalculatePacks Function
func CalculatePacks(order int, packSizes []int) map[int]int {
	// Sort pack sizes in descending order (largest first)
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	// Store result: map of pack size -> quantity of that pack used
	result := make(map[int]int)

	// Iterate through pack sizes until the order is fulfilled
	for _, size := range packSizes {
		if order >= size {
			count := order / size
			result[size] = count
			order %= size
		}
	}

	// If any remaining items can't be filled exactly, fill with the smallest pack
	if order > 0 {
		smallestPack := packSizes[len(packSizes)-1]
		result[smallestPack] += 1
	}

	// Optimize the result to minimize the total number of packs
	for size, count := range result {
		if count > 1 {
			for _, otherSize := range packSizes {
				if otherSize > size && count*size%otherSize == 0 {
					result[otherSize] += count * size / otherSize
					delete(result, size)
					break
				}
			}
		}
	}

	return result
}
