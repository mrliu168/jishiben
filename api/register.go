package api

import (
	"jishiben/dao"
	"jishiben/models"
)

type UserRegister struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Username        string `form:"user_name" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}
func(server *UserRegister)Valid() *models.Response{
	if server.PasswordConfirm!=server.Password{
		return &models.Response{
			Status:40001,
			Msg:"两次输入的密码不相同",
		}
	}
	count:=0
	dao.DB.Model(&models.User{}).Where("nickname=?",server.Nickname).Count(&count)
	if count>0{
		return &models.Response{
			Status:40001,
			Msg: "用户名已注册",
		}
	}
	return nil
}
func(server *UserRegister)Register()(models.User,*models.Response){
	user:=models.User{
		Nickname: server.Nickname,
		Username: server.Username,
	}
	if err:=server.Valid();err!=nil{
		return user,err
	}
	if err:=user.SetPassword(server.Password);err!=nil{
		return user,&models.Response{
			Status:40002,
			Msg:"密码加密失败",
		}
	}
	if err := dao.DB.Create(&user).Error; err != nil {
		return user, &models.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}



