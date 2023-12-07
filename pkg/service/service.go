package service

// Пакет service определяет интерфейсы и структуру для взаимодействия с бизнес-логикой приложения.

// Импортируем пакет repository для использования интерфейсов репозитория.
import (
	todo_app "github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/repository"
)

// Интерфейс Authorization определяет методы для работы с авторизацией.
type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
}

// Интерфейс TodoList определяет методы для работы со списками задач.
type TodoList interface {
	// Добавьте методы для реализации логики работы со списками задач.
}

// Интерфейс TodoItem определяет методы для работы с отдельными задачами.
type TodoItem interface {
	// Добавьте методы для реализации логики работы с отдельными задачами.
}

// Структура Service представляет собой общий сервис, объединяющий интерфейсы.
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Метод NewService создает новый экземпляр структуры Service, используя переданный репозиторий.
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
