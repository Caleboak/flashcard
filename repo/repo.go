package repo

import (
	"encoding/json"
	"io/ioutil"
	"net/rpc"
	"privateFlashCard/entities"
)

type flashcardRepo struct {
	filename string
}

func NewFlashcardRepo(filename string) flashcardRepo {
	return flashcardRepo{
		filename: filename,
	}
}

func (r flashcardRepo) CreateMatching(card entities.Matching) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return rpc.ServerError("unable to read info")
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return rpc.ServerError("card not created")
	}

	DbflashCardStruct.Matching = append(DbflashCardStruct.Matching, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return rpc.ServerError("unable to save card")
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateTrueFalse(card entities.TrueFalse) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return rpc.ServerError("unable to read info")
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return rpc.ServerError("card not created")
	}

	DbflashCardStruct.TrueFalse = append(DbflashCardStruct.TrueFalse, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return rpc.ServerError("unable to save card")
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateMultiple(card entities.Multiple) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return rpc.ServerError("unable to read info")
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return rpc.ServerError("card not created")
	}

	DbflashCardStruct.Multiple = append(DbflashCardStruct.Multiple, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return rpc.ServerError("unable to save card")
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateInfo(card entities.Info) error {

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return rpc.ServerError("unable to read info")
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return rpc.ServerError("card not created")
	}

	DbflashCardStruct.Info = append(DbflashCardStruct.Info, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return rpc.ServerError("unable to save card")
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateQandA(card entities.QandA) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return rpc.ServerError("unable to read info")
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return rpc.ServerError("unable to create card")
	}

	DbflashCardStruct.QandA = append(DbflashCardStruct.QandA, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return rpc.ServerError("unable to save info")
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) GetAll() ([]entities.FlashCardStruct, error) {
	deck := entities.FlashCardStruct{}
	allCards := []entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return allCards, rpc.ServerError("unable to read info")
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return allCards, rpc.ServerError("unable to create deck")
	}

	allCards = append(allCards, deck)

	return allCards, nil

}

func (r flashcardRepo) GetById(id string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	var returnDeck interface{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return returnDeck, err
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return returnDeck, err
	}
	returnDeck, err = IdCheck(id, deck)
	if err != nil {
		return returnDeck, NotFound
	}

	return returnDeck, nil
}

func (r flashcardRepo) GetByType(cardType string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	var returnDeck interface{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return returnDeck, err
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return returnDeck, err
	}
	switch cardType {
	case "Matching":
		if len(deck.Matching) != 0 {
			returnDeck = deck.Matching
			return returnDeck, nil
		}

	case "Info":
		if len(deck.Info) != 0 {
			returnDeck = deck.Info
			return returnDeck, nil
		}
	case "Multiple":
		if len(deck.Multiple) != 0 {
			returnDeck = deck.Multiple
			return returnDeck, nil
		}
	case "QandA":
		if len(deck.QandA) != 0 {
			returnDeck = deck.QandA
			return returnDeck, nil
		}
	case "TrueFalse":
		if len(deck.TrueFalse) != 0 {
			returnDeck = deck.TrueFalse
			return returnDeck, nil
		}
	}

	return returnDeck, CardEmpty
}
