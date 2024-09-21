package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pack-calculator/src/service"
	"pack-calculator/src/types"
)

func Calculate(c *gin.Context) {
	var requestBody types.OrderRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	packResult := service.CalculatePacks(requestBody.Items, requestBody.PackSizes)
	c.JSON(http.StatusOK, types.PacketResponse{Packs: packResult})
}
