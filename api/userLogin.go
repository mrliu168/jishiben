package api

import (
	"jishiben/dao"
	"jishiben/models"
)
type UserLog struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}
func (server *UserLog)Login()(models.User,*models.Response) {
	var user models.User
	if err := dao.DB.Where("username = ?", server.Username).First(&user).Error; err != nil {
		return user, &models.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	if user.CheckPassword(server.Password) == false {
		return user, &models.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}
