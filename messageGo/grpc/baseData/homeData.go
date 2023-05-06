package base_data

// home animal base data
type GrpcAnimalBaseData struct {
	AnimalId     uint64 `json:"animalId"`
	Name         string `json:"name"`
	Cid          int32  `json:"cid"`
	Favorability int32  `json:"favorability"`
	CreateMs     uint64 `json:"createMs"`
	UpdateMs     uint64 `json:"updateMs"`
}

// home data
type GrpcHomeData struct {
	LastSaveMs      int64                `json:"lastSaveMs"`
	SoilJson        string               `json:"soilJson"`
	ResourceJson    string               `json:"resourceJson"`
	AnimalJson      string               `json:"animalJson"`
	AnimalSceneJson string               `json:"animalSceneJson"`
	AnimalList      []GrpcAnimalBaseData `json:"animalList"`
	UseDefaultData  bool                 `json:"useDefaultData"`
}
