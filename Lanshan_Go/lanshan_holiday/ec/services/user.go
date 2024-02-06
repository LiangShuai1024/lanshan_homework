package services

import (
	"demo/ec/common/request"
	"demo/ec/global"
	"demo/ec/models"
	"demo/ec/utils"
	"errors"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册逻辑
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.Ec.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.Ec.DB.Create(&user).Error
	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.Ec.DB.Where("mobile = ?", params.Mobile).First(&user).Error

	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.Ec.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

//type UserUpdateService struct {
//	NickName string `form:"nickname" json:"nickname" binding:"required,min=5,max=10"`
//	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
//}

//func (userService *userService) Update(id string) (err error, user models.User) {
//
//	intId, err := strconv.Atoi(id)
//
//	//找到用户
//	err = global.Ec.DB.First(&user, intId).Error
//	if err != nil {
//		err = errors.New("用户不存在")
//	}
//	user.Name =
//	}
//	return
//}

func (userService *userService) Update(params request.Update) (err error, user models.User) {

	intId, err := strconv.Atoi(params.ID)
	err = global.Ec.DB.Exec("update users set name=?,mobile=? where id=?", params.Name, params.Mobile, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
