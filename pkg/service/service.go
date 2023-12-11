package service

// Пакет service определяет интерфейсы и структуру для взаимодействия с бизнес-логикой приложения.

// Импортируем пакет repository для использования интерфейсов репозитория.
import (
	"github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/repository"
)

// Интерфейс Authorization предоставляет методы для работы с авторизацией пользователей.
type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

// Интерфейс TodoList предоставляет методы для работы со списками дел.
type TodoList interface {
	// Здесь можно добавить методы для работы со списками дел, такие как создание, получение, обновление, удаление и другие.
}

// Интерфейс TodoItem предоставляет методы для работы с элементами в списках дел.
type TodoItem interface {
	// Здесь можно добавить методы для работы с элементами в списках дел, такие как создание, получение, обновление, удаление и другие.
}

// Структура Service объединяет интерфейсы для работы с базой данных в единую сущность.
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService создает новый экземпляр Service, используя переданный репозиторий.
func NewService(repos *repository.Repository) *Service {
	// Инициализируем Service с реализацией интерфейса Authorization для работы с пользователями в базе данных.
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		// Здесь можно добавить инициализацию других интерфейсов, если они будут реализованы.
	}
}
