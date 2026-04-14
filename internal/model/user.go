package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Status   uint8  `json:"status"` //0 represents forbiden, 1 represents normal
}
