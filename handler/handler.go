package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"privateFlashCard/entities"

	"github.com/gorilla/mux"
)

type IFlashcardService interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
	GetAll() ([]entities.FlashCardStruct, error)
	GetById(string) (interface{}, error)
	GetByType(string) (interface{}, error)
	UpdateById(string) error
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
	var c map[string]interface{}
	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if cardtype, ok := c["Type"]; ok {

		switch cardtype {
		case "Matching":
			matchcard := entities.Matching{}
			err := json.Unmarshal(file, &matchcard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.CreateMatching(matchcard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "Multiple":
			multiplecard := entities.Multiple{}
			err := json.Unmarshal(file, &multiplecard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.CreateMultiple(multiplecard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "TrueFalse":
			tfcard := entities.TrueFalse{}
			err := json.Unmarshal(file, &tfcard)
			if err != nil {
				log.Fatalln(err)

			}
			err = f.serv.CreateTrueFalse(tfcard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "Info":
			infocard := entities.Info{}
			err := json.Unmarshal(file, &infocard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateInfo(infocard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "QandA":
			QandAcard := entities.QandA{}
			err := json.Unmarshal(file, &QandAcard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateQandA(QandAcard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		default:
			http.Error(w, "Type is not valid", http.StatusBadRequest)

		}
	}

}

func (f FlashcardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	allCards, err := f.serv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	allCardsBytes, err := json.MarshalIndent(allCards, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	w.Write(allCardsBytes)

	w.Header().Set("Content-Type", "application/json")

}

func (f FlashcardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"]

	card, err := f.serv.GetById(cardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardBytes, err := json.MarshalIndent(card, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cardBytes)

}

func (f FlashcardHandler) GetByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardtype := vars["Type"]

	card, err := f.serv.GetByType(cardtype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Marshaled, err := json.MarshalIndent(card, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(Marshaled))

}

func (f FlashcardHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"]

	err := f.serv.UpdateById(cardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
