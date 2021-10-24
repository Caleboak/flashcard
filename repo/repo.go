package repo

import (
	"encoding/json"
	"io/ioutil"
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
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.Matching = append(DbflashCardStruct.Matching, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}

func (r flashcardRepo) CreateTrueFalse(card entities.TrueFalse) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.TrueFalse = append(DbflashCardStruct.TrueFalse, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}

func (r flashcardRepo) CreateMultiple(card entities.Multiple) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.Multiple = append(DbflashCardStruct.Multiple, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}

func (r flashcardRepo) CreateInfo(card entities.Info) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct.Info = append(DbflashCardStruct.Info, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}

func (r flashcardRepo) CreateQandA(card entities.QandA) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.QandA = append(DbflashCardStruct.QandA, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}
