package rating

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRouter(router *gin.RouterGroup) {
	router.POST("/add", h.AddRating)
}

func (h *Handler) AddRating(ctx *gin.Context) {
	var createDto CreateRatingDto

	if err := ctx.ShouldBindJSON(&createDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddRating(ctx, createDto)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "item rated successfully"})
}
