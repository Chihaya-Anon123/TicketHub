package api

import (
	"fmt"

	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/middleware"
	"github.com/Chihaya-Anon123/TicketHub/internal/response"
	"github.com/Chihaya-Anon123/TicketHub/internal/service"
	"github.com/gin-gonic/gin"
)

// 创建项目
func CreateProject(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "fail to get userID"))
		return
	}

	var input service.CreateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "invalid request"))
		return
	}

	output, err := service.CreateProject(userID, input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "success create project", output)
}

// 查看项目列表
func ListProjects(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "fail to get userID"))
		return
	}

	var input service.ListProjectsInput
	if err := c.ShouldBindQuery(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, fmt.Sprintf("invalid request: %v", err)))
		return
	}

	output, err := service.ListProjects(userID, input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "success list projects", output)
}
