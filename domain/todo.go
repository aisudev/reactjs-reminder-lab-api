package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID       string         `gorm:"autoIncrement"`
	UUID     string         `gorm:"primaryKey" json:"uuid"`
	Name     string         `gorm:"varchar(50)" json:"name"`
	Favorite *bool          `gorm:"default:false" json:"favorite"`
	Progress *bool          `gorm:"default:false" json:"progress"`
	CreateAt *time.Time     `gorm:"autoCreateTime" json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`
}

type TodoRepository interface {
	Create(*Todo) error
	Get(string) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(string, *Todo) error
	Delete(string) error
}

type TodoUsecase interface {
	Create(*Todo) error
	Get(string) (*Todo, error)
	GetAll() ([]Todo, error)
	Update(string, *Todo) error
	Delete(string) error
}
