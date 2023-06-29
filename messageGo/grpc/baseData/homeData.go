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

func (p *GrpcAnimalBaseData) Clear() {
	p.AnimalId = 0
	p.Name = ""
	p.Cid = 0
	p.Favorability = 0
	p.CreateMs = 0
	p.UpdateMs = 0
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

func (p *GrpcHomeData) Clear() {
	p.LastSaveMs = 0
	p.SoilJson = ""
	p.ResourceJson = ""
	p.AnimalJson = ""
	p.AnimalSceneJson = ""
	p.UseDefaultData = false
	if p.AnimalList != nil {
		p.AnimalList = p.AnimalList[:0]
	}
}
