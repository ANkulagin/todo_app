package todo_app

// Пакет todo_app определяет структуру данных, представляющую модель пользователя в приложении.

// Структура User представляет собой модель пользователя.
type User struct {
	Id       int    `json:"-" db:"id"`                   // Идентификатор пользователя (не отображается в JSON и используется в базе данных).
	Name     string `json:"name" binding:"required"`     // Имя пользователя.
	Username string `json:"username" binding:"required"` // Имя пользователя (логин).
	Password string `json:"password" binding:"required"` // Пароль пользователя.
}
