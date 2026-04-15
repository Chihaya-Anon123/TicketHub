package dao

import (
	"errors"

	"github.com/Chihaya-Anon123/TicketHub/internal/database"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
	"gorm.io/gorm"
)

// 创建项目
func CreateProject(project *model.Project) error {
	return database.DB.Create(project).Error
}

// 通过 UserID 和 ProjectName 查询项目
func GetProjectByUserIDAndName(userID uint, projectname string) (*model.Project, error) {
	var project model.Project

	err := database.DB.Where("name = ? and owner_id = ?", projectname, userID).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
}
