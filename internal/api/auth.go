package api

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/response"
	"github.com/Chihaya-Anon123/TicketHub/internal/service"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
	var input service.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.Register(input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "register success", output)
}

// 用户登录
func Login(c *gin.Context) {
	var input service.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, code.CodeInvalidParams, "invalid request")
		return
	}

	output, err := service.Login(input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "login success", output)
}
