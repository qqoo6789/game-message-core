package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// call main service get user home data
type MainServiceActionGetHomeDataInput struct {
	UserId int64 `json:"userId"`
}
type MainServiceActionGetHomeDataOutput struct {
	Success bool                   `json:"success"`
	ErrMsg  string                 `json:"errMsg"`
	UserId  int64                  `json:"userId"`
	Data    base_data.GrpcHomeData `json:"data"`
}

// call main service save animal list data
type MultiUpdateAnimalBaseDataInput struct {
	UserId  int64                          `json:"userId"`
	Animals []base_data.GrpcAnimalBaseData `json:"animals"`
}
type MultiUpdateAnimalBaseDataOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

//  GetUserAnimalList  world to main service
type GetUserAnimalListInput struct {
	UserId int64 `json:"userId"`
}
type GetUserAnimalListOutput struct {
	Success bool                           `json:"success"`
	ErrMsg  string                         `json:"errMsg"`
	UserId int64 `json:"userId"`
	Animals []base_data.GrpcAnimalBaseData `json:"animals"`
}
