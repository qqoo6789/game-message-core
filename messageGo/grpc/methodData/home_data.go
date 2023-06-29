package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// call main service save user home data
type MainServiceActionSaveHomeDataInput struct {
	UserId int64                  `json:"userId"`
	Data   base_data.GrpcHomeData `json:"data"`
}

func (p *MainServiceActionSaveHomeDataInput) Clear() {
	p.UserId = 0
	p.Data.Clear()
}

type MainServiceActionSaveHomeDataOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *MainServiceActionSaveHomeDataOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}

// call main service update  user home data last save time
type MainServiceActionUpHomeLastSaveTimeInput struct {
	UserId     int64 `json:"userId"`
	LastSaveMs int64 `json:"lastSaveMs"`
}

func (p *MainServiceActionUpHomeLastSaveTimeInput) Clear() {
	p.UserId = 0
	p.LastSaveMs = 0
}

type MainServiceActionUpHomeLastSaveTimeOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *MainServiceActionUpHomeLastSaveTimeOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}

// call main service get user home data
type MainServiceActionGetHomeDataInput struct {
	UserId int64 `json:"userId"`
}

func (p *MainServiceActionGetHomeDataInput) Clear() {
	p.UserId = 0
}

type MainServiceActionGetHomeDataOutput struct {
	Success bool                   `json:"success"`
	ErrMsg  string                 `json:"errMsg"`
	UserId  int64                  `json:"userId"`
	Data    base_data.GrpcHomeData `json:"data"`
}

func (p *MainServiceActionGetHomeDataOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
	p.UserId = 0
	p.Data.Clear()
}

// call main service save animal list data
type MultiUpdateAnimalBaseDataInput struct {
	UserId  int64                          `json:"userId"`
	Animals []base_data.GrpcAnimalBaseData `json:"animals"`
}

func (p *MultiUpdateAnimalBaseDataInput) Clear() {
	p.UserId = 0
	if p.Animals != nil {
		p.Animals = p.Animals[:0]
	}
}

type MultiUpdateAnimalBaseDataOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *MultiUpdateAnimalBaseDataOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}

//  GetUserAnimalList  world to main service
type GetUserAnimalListInput struct {
	UserId int64 `json:"userId"`
}

func (p *GetUserAnimalListInput) Clear() {
	p.UserId = 0
}

type GetUserAnimalListOutput struct {
	Success bool                           `json:"success"`
	ErrMsg  string                         `json:"errMsg"`
	UserId  int64                          `json:"userId"`
	Animals []base_data.GrpcAnimalBaseData `json:"animals"`
}

//  capture animal  scene world to main service
type CaptureAnimalInput struct {
	UserId        int64                        `json:"userId"`
	FreedAnimalId int64                        `json:"freedAnimalId"`
	CaptureAnimal base_data.GrpcAnimalBaseData `json:"animal"`
}

func (p *CaptureAnimalInput) Clear() {
	p.UserId = 0
	p.FreedAnimalId = 0
	p.CaptureAnimal.Clear()
}

type CaptureAnimalOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *CaptureAnimalOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}
