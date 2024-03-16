package api

import (
	"fmt"

	"github.com/erkylima/posts-service/internal/posts/domains"
	"github.com/erkylima/posts-service/internal/posts/services"
	"github.com/erkylima/posts-service/pkg/adapters/database"
	"github.com/erkylima/posts-service/pkg/adapters/handlers"
)

func (a *Api) InitPostsRouter() {

	repo, err := database.NewMongoConnection[domains.Post]("mongodb://admin:admin@localhost:27017", "posts_service", "posts")
	if err != nil {
		fmt.Println(err)
	}
	postService := services.NewPostService(repo)
	handler := handlers.NewPostsHandler(postService)
	a.Group.POST("/posts", handler.CreateHandler)
	a.Group.GET("/posts", handler.ListHandler)
}
