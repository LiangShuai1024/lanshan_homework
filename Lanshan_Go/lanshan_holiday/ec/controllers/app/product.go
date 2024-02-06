package app

import (
	"demo/ec/common/request"
	"demo/ec/common/response"
	"demo/ec/services"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var form request.Create
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ProductService.Create(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

