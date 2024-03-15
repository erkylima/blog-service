package api

import (
	"fmt"

	"github.com/erkylima/posts-service/internal/adapters/handlers"
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/services"
	"github.com/erkylima/posts-service/pkg/database"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Engine *gin.Engine
	Group  gin.RouterGroup
}

func NewApi() *Api {
	router := gin.Default()
	Group := router.Group("/v1")
	return &Api{
		Engine: router,
		Group:  *Group,
	}
}

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

func (g *Api) InitGin(port int) error {
	var err error
	if port != 0 {
		err = g.Engine.Run(fmt.Sprintf(":%d", port))
	} else {
		err = g.Engine.Run(":8080")
	}
	return err
}
