package app

import (
	"demo/ec/common/request"
	"demo/ec/common/response"
	"demo/ec/services"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
// 校验入参，调用 UserService 注册逻辑
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

//	func UserUpdate(c *gin.Context) {
//		var form request.Update
//		//claims, _ := util.ParseToken(c.GetHeader("Authorization"))
//		if err := c.ShouldBindJSON(&form); err != nil {
//			response.ValidateFail(c, request.GetErrorMsg(form, err))
//			return
//		}
//
//		if err, user := services.UserService.Update(c.Keys["id"].(string)); err != nil {
//			response.BusinessFail(c, err.Error())
//		} else {
//			response.Success(c, user)
//		}
//	}
func Update(c *gin.Context) {
	var form request.Update
	form = request.Update{ID: c.Keys["id"].(string)}
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Update(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}

}
