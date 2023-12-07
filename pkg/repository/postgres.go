package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Константы, представляющие названия таблиц в базе данных.
const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Config представляет конфигурацию для подключения к базе данных.
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает и возвращает новый экземпляр sqlx.DB для подключения к базе данных PostgreSQL.
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Открываем соединение с базой данных, используя параметры конфигурации.
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Проверяем соединение с базой данных.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Возвращаем экземпляр sqlx.DB и ошибку (если есть).
	return db, nil
}
