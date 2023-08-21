package base_data

type TeamMember struct {
	UserId int64  `json:"userId"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
}

func (p *TeamMember) Clear() {
	p.UserId = 0
	p.Name = ""
	p.Icon = ""
}

type Team struct {
	Id      int64        `json:"id"`
	Name    string       `json:"name"`
	Leader  int64        `json:"leader"`
	Members []TeamMember `json:"members"`
}

func (p *Team) Clear() {
	p.Id = 0
	p.Name = ""
	p.Leader = 0
	for idx, _ := range p.Members {
		p.Members[idx].Clear()
	}
	p.Members = p.Members[:0]
}
