package service

import (
	"strings"
	"unicode"

	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/dao"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
)

type CreateProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectOutput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     uint   `json:"owner_id"`
	Status      uint8  `json:"status"`
}

// 创建项目（项目创建者自动归为自己）
func CreateProject(userID uint, input CreateProjectInput) (*CreateProjectOutput, error) {
	//检验项目名
	if input.Name == "" {
		return nil, errs.New(code.CodeInvalidParams, "project name should not be empty")
	}
	if strings.IndexFunc(input.Name, unicode.IsSpace) != -1 {
		return nil, errs.New(code.CodeInvalidParams, "project name should not have spaces")
	}

	project := &model.Project{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     userID,
		Status:      1,
	}

	existproject, err := dao.GetProjectByUserIDAndName(userID, project.Name)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if existproject != nil {
		return nil, errs.New(code.CodeInvalidParams, "the user already has the project of the same name")
	}

	if err := dao.CreateProject(project); err != nil {
		return nil, errs.ErrDBError
	}

	return &CreateProjectOutput{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		OwnerID:     project.OwnerID,
		Status:      project.Status,
	}, nil
}

type ListProjectsInput struct {
	Page     int   `form:"page"`
	PageSize int   `form:"page_size"`
	Status   uint8 `form:"status"`
}

type ProjectItem struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
}

type ListProjectsOutput struct {
	ProjectItems []ProjectItem `json:"project_items"`
	Total        int64         `json:"total"`
	Page         int           `json:"page"`
	PageSize     int           `json:"page_size"`
}

// 获取用户创建的项目列表
func ListProjects(userID uint, input ListProjectsInput) (*ListProjectsOutput, error) {
	//检验并设置页码和页数
	if input.Page <= 0 {
		input.Page = 1
	}
	if input.PageSize <= 0 {
		input.PageSize = 10
	}
	if input.PageSize > 100 {
		input.PageSize = 100
	}

	if input.Status > 4 {
		return nil, errs.New(code.CodeInvalidParams, "invalid status")
	}

	projects, total, err := dao.GetProjectListByUserID(userID, input.Page, input.PageSize, input.Status)
	if err != nil {
		return nil, errs.ErrDBError
	}

	list := make([]ProjectItem, 0, len(projects))
	for _, project := range projects {
		list = append(list, ProjectItem{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Status:      project.Status,
		})
	}

	return &ListProjectsOutput{
		ProjectItems: list,
		Total:        total,
		Page:         input.Page,
		PageSize:     input.PageSize,
	}, nil
}
