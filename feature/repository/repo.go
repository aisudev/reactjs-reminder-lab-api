package repository

import (
	"reminder/domain"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(todo *domain.Todo) error {

	if err := r.db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) Get(uuid string) (*domain.Todo, error) {

	todo := new(domain.Todo)

	if err := r.db.Where("uuid = ?", uuid).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {

	todo := []domain.Todo{}

	if err := r.db.Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) Update(uuid string, todo *domain.Todo) error {

	if err := r.db.Where("uuid = ?", uuid).Updates(todo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) Delete(uuid string) error {

	if err := r.db.Where("uuid = ?", uuid).Delete(&domain.Todo{}).Error; err != nil {
		return err
	}

	return nil
}
