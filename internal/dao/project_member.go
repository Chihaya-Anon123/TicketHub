package dao

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/database"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
)

// 增加某个项目的成员
func AddProjectMember(projectMember *model.ProjectMember) error {
	return database.DB.Create(projectMember).Error
}

// 查找某个项目的所有成员
func GetProjectMembers(project *model.Project, role uint8) ([]model.User, int64, error) {
	db := database.DB.Model(&model.User{}).Where("project_id = ?", project.ID)

	if role != 0 {
		db = db.Where("member_role = ?", role)
	}

	var members []model.User
	var total int64

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Order("member_role ACS").Find(&members).Error; err != nil {
		return nil, 0, err
	}

	return members, total, nil
}
