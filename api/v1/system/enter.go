package system

import "gin-basice/service"

type ApiGroup struct {
	BaseApi
	CasbinApi
}

var casbinService = service.ServiceGroupApp.System.CasbinService
