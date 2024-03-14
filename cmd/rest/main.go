package main

import (
	"fmt"

	"github.com/erkylima/posts-service/internal/core/services"
	"github.com/erkylima/posts-service/pkg/database"
)

func main() {
	conn, err := database.NewMongoConnection("mongodb://admin:admin@localhost:27017")
	if err != nil {
		panic(err)
	}
	service := services.NewPageService(conn)
	page, err := service.ReadPage("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(page)
}
