package usecase

import (
	"reminder/domain"
)

type todoUsecase struct {
	repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}

func (u *todoUsecase) Create(todo *domain.Todo) error {

	return u.repo.Create(todo)
}

func (u *todoUsecase) Get(uuid string) (*domain.Todo, error) {

	return u.repo.Get(uuid)
}

func (u *todoUsecase) GetAll() ([]domain.Todo, error) {

	return u.repo.GetAll()
}

func (u *todoUsecase) Update(uuid string, todo *domain.Todo) error {
	return u.repo.Update(uuid, todo)
}

func (u *todoUsecase) Delete(uuid string) error {
	return u.repo.Delete(uuid)
}
