package handler

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(fc FlashcardHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcard", fc.CreateCard).Methods("POST")
	r.HandleFunc("/flashcard", fc.GetAll).Methods("GET")
	r.HandleFunc("/flashcard/Id={Id}", fc.GetById).Methods("GET")
	r.HandleFunc("/flashcard/Type={Type}", fc.GetByType).Methods("GET")

	return r

}
