package service

import (
	"fmt"
	"privateFlashCard/entities"
	"strings"
)

type IFlashcardRepo interface {
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

type FlashcardService struct {
	Repo IFlashcardRepo
}

func NewFlashcardService(f IFlashcardRepo) FlashcardService {
	return FlashcardService{
		Repo: f,
	}
}

func (f FlashcardService) CreateMatching(card entities.Matching) error {
	card.SetMatchingId()
	cardType := strings.ToLower(card.Type)         //change to lowercase for comparison
	cardCategory := strings.ToLower(card.Category) //change to lowercase for comparison
	cardQuestion := card.Question
	cardOption := card.Options
	cardAnswer := card.Answer
	if cardType != "matching" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardQuestion { //loop through the map to chack every key has a value
			if v == nil {
				return BadRequest //bad request

			}
		}
	}
	if len(cardOption) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if v == nil {
				return BadRequest //bad request
			}
		}
	}
	if len(cardAnswer) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for k, v := range cardAnswer {
			if v == nil {
				return BadRequest //bad request
			} else { //where my issue is with the interface to string
				valStr := fmt.Sprint(v) //write the interface as a string
				upperCaseVal := strings.ToUpper(valStr)
				lowerCaseVal := strings.ToLower(valStr)
				upperCaseKey := strings.ToUpper(k)
				lowerCaseKey := strings.ToLower(k)
				found := false
				//The main issue
				//compare the question with the key of the cardAnswer
				if _, ok := cardQuestion[upperCaseKey]; ok {
					found = true
				}
				if _, ok := cardQuestion[lowerCaseKey]; ok {
					found = true
				}
				if !found {
					return BadRequest //bad request
				}

				//compare the option with the value of the cardAnswer
				found = false
				if _, ok := cardOption[upperCaseVal]; ok {
					found = true
				}
				if _, ok := cardOption[lowerCaseVal]; ok {
					found = true
				}
				if !found {
					return BadRequest //bad request
				}
			}

		}
	}
	return f.Repo.CreateMatching(card)

}

func (f FlashcardService) CreateTrueFalse(card entities.TrueFalse) error {
	card.SetTrueFalseId()

	cardType := strings.ToLower(card.Type)
	cardCategory := strings.ToLower(card.Category)
	cardQuestion := card.Question
	cardAnswer := strings.ToLower(card.Answer)

	if cardType != "truefalse" {
		return BadRequest
	}

	if cardCategory != "golang" {
		return BadRequest
	}

	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if cardAnswer != "true" && cardAnswer != "false" {
		return BadRequest

	}

	return f.Repo.CreateTrueFalse(card)

}

func (f FlashcardService) CreateInfo(card entities.Info) error {
	card.SetInfoId()

	cardType := strings.ToLower(card.Type)
	cardCategory := strings.ToLower(card.Category)
	cardDetails := card.Details

	if cardType != "info" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardDetails) == 0 {
		return BadRequest
	}
	return f.Repo.CreateInfo(card)

}

func (f FlashcardService) CreateMultiple(card entities.Multiple) error {
	card.SetMultipleId()

	cardType := strings.ToLower(card.Type)
	cardCategory := strings.ToLower(card.Category)
	cardQuestion := card.Question
	cardOption := card.Options
	cardAnswer := card.Answer

	if cardType != "multiple" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 || len(cardQuestion) < 2 {
		return BadRequest //bad request
	}
	if len(cardOption) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if v == nil {
				return BadRequest //bad request
			}
		}
	}
	if len(cardAnswer) == 0 {
		return BadRequest
	} else if _, ok := cardOption[cardAnswer]; !ok {
		return BadRequest
	}
	return f.Repo.CreateMultiple(card)

}

func (f FlashcardService) CreateQandA(card entities.QandA) error {
	card.SetQandAId()

	cardType := strings.ToLower(card.Type)
	cardCategory := strings.ToLower(card.Category)
	cardQuestion := card.Question
	cardAnswer := card.Answer

	if cardType != "qanda" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 || len(cardQuestion) < 2 {
		return BadRequest //bad request
	}
	if len(cardAnswer) == 0 || len(cardQuestion) < 2 {
		return BadRequest //bad request
	}
	return f.Repo.CreateQandA(card)

}

func (f FlashcardService) GetAll() ([]entities.FlashCardStruct, error) {
	allCards, err := f.Repo.GetAll()
	if err != nil {
		return allCards, BadRequest
	}

	return allCards, nil

}

func (f FlashcardService) GetById(id string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	if len(id) == 0 {
		return deck, InvalidId
	}
	return f.Repo.GetById(id)
}

func (f FlashcardService) GetByType(cardType string) (interface{}, error) {
	validType := map[string]interface{}{
		"Info":      true,
		"QandA":     true,
		"TrueFalse": true,
		"Matching":  true,
		"Multiple":  true,
	}

	if _, ok := validType[cardType]; !ok {
		return nil, InvalidType
	}
	return f.Repo.GetByType(cardType)
}

func (f FlashcardService) UpdateById(id string) error {

	if len(id) == 0 {
		return InvalidId
	}
	return f.Repo.UpdateById(id)
}
