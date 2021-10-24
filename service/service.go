package service

import (
	"reflect"
	"strings"

	"privateFlashCard/entities"
)

type IFlashcardRepo interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
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
	cardType := strings.ToLower(card.Type)
	cardCategory := strings.ToLower(card.Category)
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
		for _, v := range cardQuestion {
			if reflect.ValueOf(v).IsNil() {
				return BadRequest //bad request

			}
		}
	}
	if len(cardOption) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if reflect.ValueOf(v).IsNil() {
				return BadRequest //bad request
			}
		}
	}
	if len(cardAnswer) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for _, v := range cardAnswer {
			if reflect.ValueOf(v).IsNil() {
				return BadRequest //bad request
			} else {
				upperCase := strings.ToUpper(v.(string))
				lowerCase := strings.ToLower(v.(string))
				found := true
				if _, ok := cardQuestion[upperCase]; !ok {
					found = true
				}
				if _, ok := cardQuestion[lowerCase]; !ok {
					found = true
				}
				if !found {
					return BadRequest //bad request
				}
				val := cardAnswer[v.(string)]
				if _, ok := cardOption[val.(string)]; !ok {
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
	cardTf := card.Tf

	if cardType != "truefalse" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if cardTf != true || cardTf != false {
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
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if len(cardOption) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if reflect.ValueOf(v).IsNil() {
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
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if len(cardAnswer) == 0 {
		return BadRequest //bad request
	}
	return f.Repo.CreateQandA(card)

}
