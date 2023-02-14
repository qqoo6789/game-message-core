package httpData

type AgentServiceResp struct {
	ErrorCode    int32  `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string `json:"errorMessage"`
	Url          string `json:"url"`
	Online       int32  `json:"online"`
	MaxOnline    int32  `json:"maxOnline"`
	CreatedAt    int64  `json:"createdAt"`
	UpdateAt     int64  `json:"updatedAt"`
}
