package lru

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LruController interface {
	Get(ctx *gin.Context)
	Set(ctx *gin.Context)
}

type lruController struct {
	service LruService
}

func NewLruController(service LruService) LruController {
	return &lruController{
		service: service,
	}
}

func (controller *lruController) Get(ctx *gin.Context) {
	key := ctx.Param("key")

	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Key not given in request",
		})
	}

	value, err := controller.service.Get(key)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"key":     key,
			"value":   value,
			"message": fmt.Sprintf("%v -> %v", key, value),
		})
	}
}

func (controller *lruController) Set(ctx *gin.Context) {
	var request SetRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Key and Value are required and cannot be empty",
		})
		return
	}

	controller.service.Set(request.Key, request.Value)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("%v is set to %v", request.Key, request.Value),
	})
}
