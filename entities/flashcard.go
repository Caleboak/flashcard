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
	Id       string                 `json:"Id"`
	Type     string                 `json:"Type"`
	Category string                 `json:"Category"`
	Question map[string]interface{} `json:"Question"`
	Options  map[string]interface{} `json:"Options"`
	Answer   map[string]interface{} `json:"Answer"`
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
	Id       string `json:"Id"`
	Type     string `json:"Type"`
	Category string `json:"Category"`
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
}

type Info struct {
	Id       string `json:"Id"`
	Type     string `json:"Type"`
	Category string `json:"Category"`
	Details  string `json:"Details"`
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
