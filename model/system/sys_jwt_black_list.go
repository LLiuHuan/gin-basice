package system

import "gin-basice/global"

type JwtBlacklist struct {
	global.AdpModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
