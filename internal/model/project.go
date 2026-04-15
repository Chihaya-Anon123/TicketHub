package model

type Project struct {
	BaseModel
	Name        string `gorm:"type:varchar(50);uniqueIndex:uk_owner_project_name;not null" json:"name"`
	Description string `gorm:"type:varchar(500)" json:"description"`
	OwnerID     uint   `gorm:"index;uniqueIndex:uk_owner_project_name;not null" json:"owner_id"`
	Status      uint8  `gorm:"not null;default:1" json:"status"` //1：新建 2：进行中 3：暂停 4：已完成
	Owner       User   `gorm:"foreignKey:OwnerID;references:ID"`
}
