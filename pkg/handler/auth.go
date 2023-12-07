package handler

// Пакет handler предоставляет обработчики запросов для авторизации в приложении.

// Импортируем необходимые пакеты.
import (
	todo_app "github.com/ANkulagin/todo-app"
	"github.com/gin-gonic/gin" // Импортируем пакет gin для работы с HTTP-запросами и ответами.
	"net/http"
)

// Метод signUp обработчика запросов, который будет вызван при запросе на регистрацию.
func (h *Handler) signUp(c *gin.Context) {
	// Здесь будет реализация логики для обработки запроса на регистрацию.
	var input todo_app.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

// Метод signIn обработчика запросов, который будет вызван при запросе на вход в систему.
func (h *Handler) signIn(c *gin.Context) {
	// Здесь будет реализация логики для обработки запроса на вход в систему.

}
