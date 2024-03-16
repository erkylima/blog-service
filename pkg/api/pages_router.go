package api

import (
	"fmt"

	"github.com/erkylima/posts-service/internal/pages/domains"
	"github.com/erkylima/posts-service/internal/pages/services"
	"github.com/erkylima/posts-service/pkg/adapters/database"
	"github.com/erkylima/posts-service/pkg/adapters/handlers"
)

func (a *Api) InitPagesRouter() {

	repo, err := database.NewMongoConnection[domains.Page]("mongodb://admin:admin@localhost:27017", "posts_service", "pages")
	if err != nil {
		fmt.Println(err)
	}
	pageService := services.NewPageService(repo)
	handler := handlers.NewPagesHandler(pageService)
	a.Group.POST("/pages", handler.CreateHandler)
	a.Group.GET("/pages", handler.ListHandler)
}
