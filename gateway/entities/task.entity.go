package entities

import "gorm.io/gorm"

type TaskEntity struct {
	gorm.Model
	Title        string
	Description  string
	TaskStatusID uint
	Completed    bool
}
