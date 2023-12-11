package repository

// Пакет, который обрабатывает работу с базой данных SQL.
import (
	"fmt"
	"github.com/ANkulagin/todo-app"
	"github.com/jmoiron/sqlx"
)

// Структура AuthPostgres представляет собой реализацию интерфейса Authorization для PostgreSQL.
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres создает новый экземпляр AuthPostgres с переданным объектом базы данных.
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser создает нового пользователя в базе данных и возвращает его идентификатор.
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	// Формируем SQL-запрос для вставки нового пользователя и получения его идентификатора.
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	// Выполняем SQL-запрос, передавая параметры из объекта пользователя.
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	// Сканируем результат запроса, получая идентификатор созданного пользователя.
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	// Возвращаем идентификатор созданного пользователя.
	return id, nil
}
