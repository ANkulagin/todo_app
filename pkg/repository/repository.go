// Пакет repository определяет интерфейсы для взаимодействия с хранилищем данных.
package repository

import "github.com/jmoiron/sqlx"

// Интерфейс Authorization определяет методы для работы с авторизацией.
type Authorization interface {
	// Добавьте методы для реализации логики авторизации.
}

// Интерфейс TodoList определяет методы для работы со списками задач.
type TodoList interface {
	// Добавьте методы для реализации логики работы со списками задач.
}

// Интерфейс TodoItem определяет методы для работы с отдельными задачами.
type TodoItem interface {
	// Добавьте методы для реализации логики работы с отдельными задачами.
}

// Структура Repository представляет собой хранилище данных, объединяющее интерфейсы.
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Метод NewRepository создает новый экземпляр структуры Repository.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
