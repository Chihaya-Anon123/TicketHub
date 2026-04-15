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
