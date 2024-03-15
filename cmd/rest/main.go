package main

import (
	"github.com/erkylima/posts-service/pkg/api"
)

func main() {
	api := api.NewApi()

	api.InitPagesRouter()
	err := api.InitGin(0)
	if err != nil {
		panic(err)
	}
}
