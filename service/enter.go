package service

import "gin-basice/service/system"

type ServiceGroup struct {
	System system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
