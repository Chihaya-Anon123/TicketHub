package model

type ProjectMember struct {
	BaseModel
	MemberID    uint   `gorm:"uniqueIndex:uk_member_id_project_id;not null" json:"member_id"`
	MemberName  string `json:"member_name"`
	ProjectID   uint   `gorm:"uniqueIndex:uk_member_id_project_id;not null" json:"project_id"`
	ProjectName string `json:"project_name"`
	MemberRole  uint8  `json:"member_role"` //1:owner,2: admin,3: member

	Member  User    `gorm:"foreignKey:MemberID;references:ID"`
	Project Project `gorm:"foreignKey:ProjectID;references:ID"`
}
