package methodData

// call system message
type BroadCastSystemMessageInput struct {
	Sender   string `json:"sender"`
	Content  string `json:"content"`
	AliveSec uint64 `json:"aliveSec"`
}

func (p *BroadCastSystemMessageInput) Clear() {
	p.Sender = ""
	p.Content = ""
	p.AliveSec = 0
}

type BroadCastSystemMessageOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *BroadCastSystemMessageOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}
