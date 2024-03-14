package main

import (
	"fmt"
	"time"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/services"
	"github.com/erkylima/posts-service/pkg/database"
)

func main() {
	conn, err := database.NewMongoConnection("mongodb://admin:admin@localhost:27017")
	if err != nil {
		panic(err)
	}
	service := services.NewPageService(conn)
	page := &domains.Page{
		Title: "Agora a coisa ficou séria",
		Slug:  "testado",
		Body:  "Novo Body",
		Url:   "Tem URL",
		Type:  "tipográfico",
		// Id:    uuid.NewString(),
		Tags: []domains.Tag{
			{
				Slug: "tag1",
				Name: "Tag 1",
			},
			{
				Slug: "tag2",
				Name: "Tag 2",
			},
		},
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	_, errs := service.CreatePage(page)
	if errs != nil {
		panic(err)
	}
	page = &domains.Page{
		Title: "Agora a coisa ficou normal",
		Slug:  "testada",
		Body:  "Novo Conteudo",
		Url:   "Tem URLS",
		Type:  "tipográfico",
		Tags: []domains.Tag{
			{
				Slug: "tag1",
				Name: "Tag 1",
			},
			{
				Slug: "tag2",
				Name: "Tag 2",
			},
		},
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	service.DeletePage("testado")
	fmt.Println(page)
}
