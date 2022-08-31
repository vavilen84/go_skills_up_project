package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vavilen84/go_skills_up_project/handlers"
	"github.com/vavilen84/go_skills_up_project/store"
	"log"
)

func main() {
	fmt.Println(123)
	godotenv.Load()
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
