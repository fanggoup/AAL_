package service

import (
	"AAL_time/modle"
	"AAL_time/package/e"
	"AAL_time/package/utils"
	"AAL_time/serializer"
	"gorm.io/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16`
}

func (service *UserService) Register() *serializer.Response{
	code := e.SUCCESS
	var user modle.User
	var count int64
	modle.DB.Model(&modle.User{}).Where("user_name=?",service.UserName).First(&user).Count(&count)
	if count == 1{
		code := e.ErrorExistUser
		return &serializer.Response{
			Status : code,
			Msg: e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	// 加密
	if err := user.SetPassword(service.Password);err != nil{
		utils.LogrusObj.Info("加密",err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	// 创建用户
	if err := modle.DB.Create(&user).Error;err != nil{
		utils.LogrusObj.Info("创建用户",err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Login() *serializer.Response{
	var user modle.User
	code := e.SUCCESS
	// 用户不存在
	if err := modle.DB.Where("user_name=?",service.UserName).First(&user).Error;err != nil{
		if err.Error() == gorm.ErrRecordNotFound.Error(){
			utils.LogrusObj.Info("用户不存在",err)
			code := e.ErrorNotExistUser
			return &serializer.Response{
				Status: code,
				Msg: e.GetMsg(code),
			}

		}
		utils.LogrusObj.Info("查找用户其他问题",err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	if ! user.CheckPassword(service.Password){
		code := e.ErrorNotCompare
		return &serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	token,err := utils.GenerateToken(user.ID,service.UserName,0)
	if err != nil{
		utils.LogrusObj.Info("派发token失败",err)
		code = e.ErrorAuthToken
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data: serializer.TokenData{User: serializer.BuildUser(user),Token: token},
		Msg: e.GetMsg(code),
	}
}