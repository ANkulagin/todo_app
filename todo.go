package todo_app

// Пакет todo_app определяет структуры данных, представляющие сущности в приложении для управления задачами.

// Структура TodoList представляет собой модель задачи в списке.
type TodoList struct {
	Id          int    `json:"id" db:"id"`                          // Идентификатор списка задач.
	Title       string `json:"title" db:"title" binding:"required"` // Название списка задач.
	Description string `json:"description" db:"description"`        // Описание списка задач.
}

// Структура UsersList представляет собой связь между пользователями и списками задач.
type UsersList struct {
	Id     int // Идентификатор связи.
	UserId int // Идентификатор пользователя.
	ListId int // Идентификатор списка задач.
}

// Структура TodoItem представляет собой модель отдельной задачи.
type TodoItem struct {
	Id          int    `json:"id" db:"id"`                          // Идентификатор задачи.
	Title       string `json:"title" db:"title" binding:"required"` // Заголовок задачи.
	Description string `json:"description" db:"description"`        // Описание задачи.
	Done        bool   `json:"done" db:"done"`                      // Флаг завершенности задачи.
}

// Структура ListsItem представляет собой связь между списками задач и отдельными задачами.
type ListsItem struct {
	Id     int // Идентификатор связи.
	ListId int // Идентификатор списка задач.
	ItemId int // Идентификатор задачи в списке.
}
