package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRouter(router *gin.RouterGroup) {
	router.POST("/add", h.AddComment)
	router.GET("/item/:itemId", h.GetCommentsByItemId)
}

func (h *Handler) AddComment(ctx *gin.Context) {
	var createDto CreateCommentDto

	if err := ctx.ShouldBindJSON(&createDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddComment(ctx, createDto)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "comment created"})
}

func (h *Handler) GetCommentsByItemId(ctx *gin.Context) {
	param := ctx.Param("itemId")

	itemId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comments, err := h.service.GetCommentsByItemId(ctx, itemId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
