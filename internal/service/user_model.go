package service

import "time"

type UserModel struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"type:varchar(255);uniqueIndex"`
	Name      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName is optional, but useful if your table is named differently
func (UserModel) TableName() string {
	return "users"
}
