package main

import (
	"final-project/controller"
	"final-project/initialization"
	"final-project/repository"
	"final-project/service"
	"fmt"
	"log"
	"net/http"
)

func init() {
	cfg, err := initialization.LoadConfig("./configuration")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initialization.ConnectDB(&cfg)
}

func main() {

	repos := repository.NewRepository(initialization.DB)
	server := service.NewService(repos)
	c := controller.NewController(server)

	router := c.InitRoutes()
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
