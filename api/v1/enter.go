package v1

import "gin-basice/api/v1/system"

type ApiV1Group struct {
	System system.ApiGroup
}

var ApiV1GroupApp = new(ApiV1Group)
