package methodData

// call system message
type CheckLoginInput struct {
	UserId int64 `json:"userId"`
}

func (p *CheckLoginInput) Clear() {
	p.UserId = 0
}

type CheckLoginOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
	IsLogin bool   `json:"isLogin"`
}

func (p *CheckLoginOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
	p.IsLogin = false
}
