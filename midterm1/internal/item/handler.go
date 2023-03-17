package item

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRouter(router *gin.RouterGroup) {
	router.GET("/find", h.FindItem)
}

func (h *Handler) FindItem(ctx *gin.Context) {
	name := ctx.Query("name")

	item, err := h.service.FindItem(ctx, name)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, MapItemToDto(item))
}
