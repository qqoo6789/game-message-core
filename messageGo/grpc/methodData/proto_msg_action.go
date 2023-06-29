package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// agent service forward client message request
type PullClientMessageInput struct {
	MsgVersion     int64  `json:"msgVersion"`     // 消息版本号 值为毫秒时间戳
	AgentAppId     string `json:"agentAppId"`     // 网关 appId
	UserId         int64  `json:"userId"`         //
	SocketId       string `json:"socketId"`       //
	SceneServiceId string `json:"sceneServiceId"` //
	MsgId          int32  `json:"msgId"`
	MsgBody        []byte `json:"msgBody"` // proto message marshal bytes
}

func (p *PullClientMessageInput) Clear() {
	p.MsgVersion = 0
	p.AgentAppId = ""
	p.UserId = 0
	p.SocketId = ""
	p.SceneServiceId = ""
	p.MsgId = 0
	if p.MsgBody != nil {
		p.MsgBody = p.MsgBody[:0]
	}
}

// agent service forward client message response
type PullClientMessageOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *PullClientMessageOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}

// other service broadcast proto message to client request
type BroadCastToClientInput struct {
	MsgVersion   int64  `json:"msgVersion"`   // 消息版本号 值为毫秒时间戳
	ServiceAppId string `json:"serviceAppId"` // 网关 appId
	UserId       int64  `json:"userId"`       //
	SocketId     string `json:"socketId"`     //
	MsgId        int32  `json:"msgId"`
	MsgBody      []byte `json:"msgBody"` // proto message marshal bytes
}

func (p *BroadCastToClientInput) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.SocketId = ""
	p.MsgId = 0
	if p.MsgBody != nil {
		p.MsgBody = p.MsgBody[:0]
	}
}

// other service broadcast proto message to client response
type BroadCastToClientOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *BroadCastToClientOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}

// other service multiple broadcast proto message to client request
type MultipleBroadCastToClientInput struct {
	MsgVersion   int64                          `json:"msgVersion"`   // 消息版本号 值为毫秒时间戳
	ServiceAppId string                         `json:"serviceAppId"` // from service app id
	MsgDataList  []base_data.MultiClientMsgData `json:"msgDataList"`
}

func (p *MultipleBroadCastToClientInput) Clear() {
	p.MsgVersion = 0
	p.ServiceAppId = ""
	if p.MsgDataList != nil {
		p.MsgDataList = p.MsgDataList[:0]
	}
}

// other service multiple broadcast proto message to client response
type MultipleBroadCastToClientOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *MultipleBroadCastToClientOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}
