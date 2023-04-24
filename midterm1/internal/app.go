package internal

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/auth"
	"github.com/Tamir1205/midterm1/internal/comment"
	"github.com/Tamir1205/midterm1/internal/config"
	"github.com/Tamir1205/midterm1/internal/item"
	"github.com/Tamir1205/midterm1/internal/rating"
	commentRepository "github.com/Tamir1205/midterm1/internal/storage/comment"
	"github.com/Tamir1205/midterm1/internal/storage/items"
	ratingRepository "github.com/Tamir1205/midterm1/internal/storage/rating"
	"github.com/Tamir1205/midterm1/internal/storage/users"
	"github.com/Tamir1205/midterm1/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{config: config}
}

func (a *App) Run() error {
	p := postgres.Config{
		Url:         a.config.DB.Url,
		MaxOpenCons: a.config.DB.MaxOpenConn,
		MaxIdleCons: a.config.DB.MaxIdleConn,
	}

	client, err := postgres.NewClient(context.Background(), p)
	if err != nil {
		return err
	}

	engine := gin.New()

	userRepository := users.NewRepository(client)
	userService := auth.NewService(userRepository)
	auth.NewHandler(userService).RegisterRouter(engine.Group("/auth"))

	commentService := comment.NewService(commentRepository.NewRepository(client))
	comment.NewHandler(
		commentService,
	).RegisterRouter(engine.Group("/comment"))

	ratingService := rating.NewService(
		ratingRepository.NewRepository(client))

	rating.NewHandler(ratingService).RegisterRouter(engine.Group("/rating"))

	itemRepository := items.NewRepository(client)
	itemService := item.NewService(itemRepository, ratingService, commentService)
	item.NewHandler(itemService).RegisterRouter(engine.Group("/item"))

	return engine.Run(":" + a.config.Server.Port)
}
