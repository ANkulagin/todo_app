package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	// Получаем значение заголовка Authorization из запроса
	header := c.GetHeader(authorizationHeader)

	// Проверяем, не является ли заголовок Authorization пустым
	if header == "" {
		// Если пустой, отправляем ошибку "401 Unauthorized" с сообщением "empty auth header"
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// Разбиваем значение заголовка на части с использованием пробела в качестве разделителя
	headerParts := strings.Split(header, " ")

	// Проверяем, что заголовок разделен на две части (например, "Bearer token")
	if len(headerParts) != 2 {
		// Если не соответствует ожидаемому формату, отправляем ошибку "401 Unauthorized" с сообщением "invalid auth header"
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
