package system

import (
	"fmt"

	"gin-basice/global"
	"gin-basice/model/system"
	"gin-basice/utils"
)

type UserService struct {
}

// Login 用户登陆
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.AdpDb {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.AdpDb.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}
