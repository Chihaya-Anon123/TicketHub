package api

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/response"
	"github.com/Chihaya-Anon123/TicketHub/internal/service"
	"github.com/gin-gonic/gin"
)

// 增加项目成员
func AddProjectMember(c *gin.Context) {
	var input service.AddProjectMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.AddProjectMember(input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "success add member", output)
}
