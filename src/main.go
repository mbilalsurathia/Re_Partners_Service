package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type RequestBody struct {
	Items     int   `json:"items"`
	PackSizes []int `json:"pack_sizes"`
}

type ResponseBody struct {
	Packs map[int]int `json:"packs"`
}

func calculatePacks(order int, packSizes []int) map[int]int {
	// Sort pack sizes in descending order (largest first)
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	// Store result: map of pack size -> quantity of that pack used
	result := make(map[int]int)

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

	return result
}

func calculate(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	packResult := calculatePacks(requestBody.Items, requestBody.PackSizes)
	c.JSON(http.StatusOK, ResponseBody{Packs: packResult})
}

func main() {
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowMethods:        []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowPrivateNetwork: false,
		AllowHeaders:        []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials:    true,
		ExposeHeaders:       []string{"Content-Length"},
		MaxAge:              12 * 3600, // Cache the preflight request for 12 hours
	}))

	// Define the API endpoint
	router.POST("/calculate", calculate)

	router.Run(":8080")
}
