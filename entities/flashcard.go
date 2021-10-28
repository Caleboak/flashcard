package entities

import "github.com/google/uuid"

type FlashCardStruct struct {
	Matching  []Matching
	TrueFalse []TrueFalse
	Multiple  []Multiple
	Info      []Info
	QandA     []QandA
}

type Matching struct {
	Id       string
	Type     string
	Category string
	Question map[string]interface{}
	Options  map[string]interface{}
	Answer   map[string]interface{}
}

type Multiple struct {
	Id       string
	Type     string
	Category string
	Question string
	Options  map[string]interface{}
	Answer   string
}

type TrueFalse struct {
	Id       string
	Type     string
	Category string
	Question string
	Answer   string
}

type Info struct {
	Id       string
	Type     string
	Category string
	Details  string
}

type QandA struct {
	Id       string
	Type     string
	Category string
	Question string
	Answer   string
}

func (f *Matching) SetMatchingId() {
	f.Id = uuid.New().String()
}

func (f *Multiple) SetMultipleId() {
	f.Id = uuid.New().String()
}

func (f *TrueFalse) SetTrueFalseId() {
	f.Id = uuid.New().String()
}

func (f *Info) SetInfoId() {
	f.Id = uuid.New().String()
}

func (f *QandA) SetQandAId() {
	f.Id = uuid.New().String()
}
