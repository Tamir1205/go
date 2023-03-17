package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRouter(router *gin.RouterGroup) {
	router.POST("/register", h.Register)
	router.POST("/signin", h.SingIn)
}

func (h *Handler) Register(ctx *gin.Context) {
	var dto CreateUserDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.service.CreateUser(ctx, dto)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": userId})
}

func (h *Handler) SingIn(c *gin.Context) {
	var request SignInRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SignIn(c, request); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}
