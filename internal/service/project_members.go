package service

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/dao"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
)

type AddProjectMemberInput struct {
	UserID    uint  `json:"user_id"`
	ProjectID uint  `json:"project_id"`
	Role      uint8 `json:"role"`
}

type AddProjectMemberOutput struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	UserName    string `json:user_name"`
	ProjectID   uint   `json:"project_id"`
	ProjectName string `json:"project_name"`
	Role        uint8  `json:"role"`
}

// 增加项目成员
// 创始人只能有一个，因此 role 不能是 1：owner
func AddProjectMember(input AddProjectMemberInput) (*AddProjectMemberOutput, error) {
	user, err := dao.GetUserByID(input.UserID)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if user == nil {
		return nil, errs.New(code.CodeNotFound, "user not found")
	}

	project, err := dao.GetProjectByID(input.ProjectID)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if project == nil {
		return nil, errs.New(code.CodeNotFound, "project not found")
	}

	if input.Role != 2 && input.Role != 3 {
		return nil, errs.New(code.CodeInvalidParams, "invalid role")
	}

	projectMember := model.ProjectMember{
		MemberID:    user.ID,
		MemberName:  user.Username,
		ProjectID:   project.ID,
		ProjectName: project.Name,
		MemberRole:  input.Role,
	}
	if err := dao.AddProjectMember(&projectMember); err != nil {
		return nil, errs.ErrDBError
	}

	output := AddProjectMemberOutput{
		ID:          projectMember.ID,
		UserID:      projectMember.MemberID,
		UserName:    projectMember.MemberName,
		ProjectID:   projectMember.ProjectID,
		ProjectName: projectMember.ProjectName,
		Role:        projectMember.MemberRole,
	}
	return &output, err
}
