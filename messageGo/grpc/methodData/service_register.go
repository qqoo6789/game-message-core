package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// 服务信息用于服务 注册和释放
type ServiceRegisterInput struct {
	Service    base_data.ServiceData `json:"service"`    //
	RegisterAt int64                 `json:"registerAt"` // register service current time MS
}

func (p *ServiceRegisterInput) Clear() {
	p.Service.Clear()
	p.RegisterAt = 0
}

// 服务信息用于服务 注册和释放
type ServiceRegisterOutput struct {
	Success bool `json:"success"`
	// register service current time MS
	RegisterAt int64 `json:"registerAt"`
	// service manager current time MS
	ManagerAt int64 `json:"managerAt"`
}

func (p *ServiceRegisterOutput) Clear() {
	p.Success = false
	p.RegisterAt = 0
	p.ManagerAt = 0
}
