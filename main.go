package main

import (
	"dataCollectionTree/app/controller"
	"dataCollectionTree/app/repository"
	"dataCollectionTree/config"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	conf := config.GetConfig()

	toDorp := repository.NewToDoRepo()
	controller.NewTodo(serveMux, toDorp)

	log.Println("listening on port:", conf.HttpConfig.HostPort)
	if err := http.ListenAndServe(conf.HttpConfig.HostPort, serveMux); err != nil {
		log.Println("server server", err)
	}
}
