package app

import (
	"context"
	"lbms/app/config"
	"lbms/app/controllers"
	"lbms/app/middleware"
	"lbms/app/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	version string `yaml:"version"`
	status  string `yaml:"status"`
}

func (a *App) GetStatus() string {
	return a.status
}

func (a *App) GetVersion() string {
	return a.version
}

func New() (*App, error) {
	return &App{}, nil
}

func (a *App) Run() error {
	r := gin.Default()
	r.Use(middleware.Cors())

	c, err := config.New(".yaml")
	if err != nil {
		return err
	}

	ctx := context.Background()
	connection, err := mongo.Connect(ctx, options.Client().ApplyURI(c.DSN))
	if err != nil {
		return err
	}

	db := connection.Database("db1")

	// connection to the first book collection
	bookCollection := db.Collection("books")
	bookRepo := repositories.NewBookRepo(bookCollection)
	bookHandler := controllers.NewBookHandler(bookRepo)

	r.POST("/books", bookHandler.NewBook)
	r.GET("/books/:id", bookHandler.Get)

	if err := r.Run(":" + c.Port); err != nil {
		return err
	}
	return nil
}
