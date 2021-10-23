package main

import (
	//"net/http"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"privateFlashCard/handler"
	"privateFlashCard/repo"
	"privateFlashCard/service"
	//"fmt"
	//"github.com/gpark1005/FlashCardsTeamTwo/entities"
	//"github.com/gpark1005/FlashCardsTeamTwo/handler"
)

func main() {

	filename := "flashcard.json"

	ext := filepath.Ext(filename)

	if ext != ".json" {
		log.Fatalln("Invalid file")
	}

	repo := repo.NewFlashcardRepo(filename)
	serv := service.NewFlashcardService(repo)
	handle := handler.NewFlashcardHandler(serv)

	router := handler.ConfigureRouter(handle)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	fmt.Println("Successfully running server 8080")

	log.Fatal(server.ListenAndServe())

}
