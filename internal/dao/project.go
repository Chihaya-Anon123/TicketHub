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

// 通过 ID 查找项目
func GetProjectByID(projectID uint) (*model.Project, error) {
	var project model.Project

	err := database.DB.Where("id = ?", projectID).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
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

// 查询 UserID 下的所有项目
func GetProjectListByUserID(userID uint, page, pageSize int, status uint8) ([]model.Project, int64, error) {
	var projects []model.Project
	var total int64

	db := database.DB.Model(&model.Project{}).Where("owner_id = ?", userID)

	if status != 0 {
		db = db.Where("status = ?", status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&projects).Error; err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}
