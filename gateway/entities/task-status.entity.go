package entities

import "gorm.io/gorm"

type TaskStatusEntity struct {
	gorm.Model
	Status string
	Active bool
	Tasks  []TaskEntity
}
