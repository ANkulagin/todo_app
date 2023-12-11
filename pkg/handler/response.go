package handler

// Пакет, который обрабатывает HTTP-запросы и формирует HTTP-ответы.
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура error представляет JSON-структуру для сообщения об ошибке.
type error struct {
	Message string `json:"message"`
}

// Функция newErrorResponse создает и возвращает JSON-ответ с сообщением об ошибке.
// Она также записывает сообщение об ошибке в журнал.
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	// Записываем сообщение об ошибке в журнал.
	logrus.Error(message)

	// Отправляем JSON-ответ с сообщением об ошибке и указанным HTTP-статусом.
	c.AbortWithStatusJSON(statusCode, error{message})
}
