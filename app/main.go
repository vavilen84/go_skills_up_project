package main

import (
	"github.com/vavilen84/go_skills_up_project/handlers"
	"github.com/vavilen84/go_skills_up_project/store"
	"log"
)

func main() {
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
