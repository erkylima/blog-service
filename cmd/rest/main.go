package main

import (
	"context"
	"fmt"
	"time"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/services"
	"github.com/erkylima/posts-service/pkg/database"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := database.NewMongoConnection[domains.Page](ctx, "mongodb://admin:admin@localhost:27017", "posts_service", "pages")
	if err != nil {
		fmt.Printf("%s", err.Error())
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
		panic(errs)
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
	list, errP := service.ListPages([]domains.Filter{
		{
			Key:   "url",
			Value: "Tem URL",
		},
	})
	if errP != nil {
		panic(errP)
	}
	fmt.Println(list)
}
