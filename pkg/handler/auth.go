package handler

// Пакет handler предоставляет обработчики запросов для авторизации в приложении.

// Импортируем необходимые пакеты.
import (
	"github.com/ANkulagin/todo-app"
	"github.com/gin-gonic/gin" // Импортируем пакет gin для работы с HTTP-запросами и ответами.
	"net/http"
)

// Функция signUp обрабатывает запрос на регистрацию нового пользователя.
func (h *Handler) signUp(c *gin.Context) {
	// Инициализируем переменную input для хранения данных пользователя.
	var input todo.User
	// Привязываем JSON-данные запроса к переменной input.
	if err := c.BindJSON(&input); err != nil {
		// Если возникает ошибка при привязке данных, возвращаем ошибку BadRequest.
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Создаем нового пользователя с использованием сервиса Authorization.
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		// Если возникает ошибка при создании пользователя, возвращаем ошибку InternalServerError.
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем успешный ответ с идентификатором созданного пользователя.
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Функция signIn обрабатывает запрос на вход пользователя в систему.
func (h *Handler) signIn(c *gin.Context) {
	// Здесь будет код обработки входа пользователя, который нужно добавить.
	// Пример использования функции newErrorResponse для возвращения ошибки:
	// newErrorResponse(c, http.StatusUnauthorized, "Неверные учетные данные")
}
