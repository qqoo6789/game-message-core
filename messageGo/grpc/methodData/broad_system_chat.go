package methodData

// call system message
type BroadCastSystemMessageInput struct {
	Sender   string `json:"sender"`
	Content  string `json:"content"`
	AliveSec uint64 `json:"aliveSec"`
}

type BroadCastSystemMessageOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}
