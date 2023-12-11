package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Константы, представляющие имена таблиц в базе данных.
const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Структура Config представляет конфигурацию подключения к базе данных PostgreSQL.
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает новое подключение к базе данных PostgreSQL на основе переданной конфигурации.
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Открываем новое подключение к базе данных PostgreSQL, используя параметры из конфигурации.
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Проверяем, что подключение к базе данных установлено успешно.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Возвращаем созданное подключение к базе данных.
	return db, nil
}
