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
	DbflashCardStruct := entities.DbFlashcard{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	for _, v := range DbflashCardStruct.FlashCardStruct {
		v.Matching = append(v.Matching, card)

	}

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
	DbflashCardStruct := entities.DbFlashcard{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	for _, v := range DbflashCardStruct.FlashCardStruct {
		v.TrueFalse = append(v.TrueFalse, card)

	}

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
	DbflashCardStruct := entities.DbFlashcard{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	for _, v := range DbflashCardStruct.FlashCardStruct {
		v.Multiple = append(v.Multiple, card)

	}

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
	DbflashCardStruct := entities.DbFlashcard{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	for _, v := range DbflashCardStruct.FlashCardStruct {
		v.Info = append(v.Info, card)

	}

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
	DbflashCardStruct := entities.DbFlashcard{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	for _, v := range DbflashCardStruct.FlashCardStruct {
		v.QandA = append(v.QandA, card)

	}

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}
