package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"privateFlashCard/entities"
	"privateFlashCard/repo"
	"privateFlashCard/service"

	"github.com/gorilla/mux"
)

type IFlashcardService interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
}

type FlashcardHandler struct {
	serv IFlashcardService
}

func NewFlashcardHandler(f IFlashcardService) FlashcardHandler {
	return FlashcardHandler{
		serv: f,
	}
}

func (f FlashcardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cardtype := vars["Type"]

	switch cardtype {
	case "Matching":
		matchcard := entities.Matching{}
		err := json.NewDecoder(r.Body).Decode(&matchcard)
		if err != nil {
			log.Fatalln(err)
		}
		err = f.serv.CreateMatching(matchcard)
		if err != nil {
			switch err {
			case service.BadRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			case repo.NotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return

			case repo.ServerError:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}
	case "Multiple":
		multiplecard := entities.Multiple{}
		err := json.NewDecoder(r.Body).Decode(&multiplecard)
		if err != nil {
			log.Fatalln(err)
		}
		err = f.serv.CreateMultiple(multiplecard)
		if err != nil {
			switch err {
			case service.BadRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			case repo.NotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return

			case repo.ServerError:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}
	case "TrueFalse":
		tfcard := entities.TrueFalse{}
		err := json.NewDecoder(r.Body).Decode(&tfcard)
		if err != nil {
			log.Fatalln(err)

		}
		err = f.serv.CreateTrueFalse(tfcard)
		if err != nil {
			switch err {
			case service.BadRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			case repo.NotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return

			case repo.ServerError:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}
	case "Info":
		infocard := entities.Info{}
		err := json.NewDecoder(r.Body).Decode(&infocard)
		if err != nil {
			log.Fatalln(err)
		}

		err = f.serv.CreateInfo(infocard)
		if err != nil {
			switch err {
			case service.BadRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			case repo.NotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return

			case repo.ServerError:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}
	case "QandA":
		QandAcard := entities.QandA{}
		err := json.NewDecoder(r.Body).Decode(&QandAcard)
		if err != nil {
			log.Fatalln(err)
		}

		err = f.serv.CreateQandA(QandAcard)
		if err != nil {
			switch err {
			case service.BadRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			case repo.NotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
				return

			case repo.ServerError:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Well done, Flashcard is Created"))
	w.Header().Set("Content-Type", "application/json")

}
