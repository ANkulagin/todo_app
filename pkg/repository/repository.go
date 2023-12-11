// Пакет repository определяет интерфейсы для взаимодействия с хранилищем данных.
package repository

import (
	"github.com/ANkulagin/todo-app"
	"github.com/jmoiron/sqlx"
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

// Структура Repository объединяет интерфейсы для работы с базой данных в единую сущность.
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository создает новый экземпляр Repository, используя переданное подключение к базе данных.
func NewRepository(db *sqlx.DB) *Repository {
	// Инициализируем Repository с реализацией интерфейса Authorization для работы с пользователями в базе данных.
	return &Repository{
		Authorization: NewAuthPostgres(db),
		// Здесь можно добавить инициализацию других интерфейсов, если они будут реализованы.
	}
}
