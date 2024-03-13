package main

import (
	"github.com/erkylima/posts-service/internal/adapters/repository"
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/pkg/database"
)

func main() {
	conn, err := database.NewMongoConnection("mongodb://admin:admin@localhost:27017")
	if err != nil {
		panic(err)
	}
	pageRepository := repository.NewPageRepository()
	page := &domains.Page{
		Title: "test",
		Slug:  "test",
		Body:  "test",
	}
	page, _ = pageRepository.Create(page)
	conn.Push(page)
}
