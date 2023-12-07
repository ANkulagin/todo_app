package service

import (
	todo_app "github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}
func (a *AuthService) CreateUser(user todo_app.User) (int, error) {
	return a.repo.CreateUser(user)
}
