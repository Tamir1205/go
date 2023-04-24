package item

import (
	"github.com/Tamir1205/midterm1/internal/storage/items"
	"github.com/gin-gonic/gin"
	"strings"
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
	//name := ctx.Query("name")
	filters := CreateFilter(ctx)

	items, err := h.service.FindItem(ctx, filters)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, items)
}

func CreateFilter(c *gin.Context) []items.Filter {
	var filters []items.Filter
	if c.Query("name") != "" {
		filters = append(filters, items.Filter{Key: "name", Value: c.Query("name")})
	}
	if c.Query("price") != "" && strings.Contains(c.Query("price"), "-") {
		filters = append(filters, items.Filter{Key: "price", Value: c.Query("price")})
	}
	if c.Query("rating") != "" && strings.Contains(c.Query("rating"), "-") {
		filters = append(filters, items.Filter{Key: "rating", Value: c.Query("rating")})
	}
	return filters
}
