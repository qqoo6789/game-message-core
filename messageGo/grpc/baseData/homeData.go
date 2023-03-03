package base_data

// home animal base data
type GrpcAnimalBaseData struct {
	AnimalId     uint64 `json:"animalId"`
	Name         string `json:"name"`
	Cid          int32  `json:"cid"`
	Favorability int32  `json:"favorability"`
}

// home data
type GrpcHomeData struct {
	SoilJson     string               `json:"soilJson"`
	ResourceJson string               `json:"resourceJson"`
	AnimalJson   string               `json:"animalJson"`
	AnimalList   []GrpcAnimalBaseData `json:"animalList"`
}
