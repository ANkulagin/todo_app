package handler

import (
	"github.com/ANkulagin/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}
func (h *Handler) updateList(c *gin.Context) {
	// Получаем идентификатор пользователя из контекста запроса
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Извлекаем идентификатор списка из параметра запроса
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Если идентификатор не удалось извлечь, возвращаем ошибку "400 Bad Request" с сообщением "invalid id param"
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	// Извлекаем данные для обновления списка из тела JSON-запроса
	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		// Если не удалось прочитать JSON из запроса, возвращаем ошибку "400 Bad Request" с сообщением об ошибке
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызываем метод сервиса для обновления списка
	if err := h.services.Update(userId, id, input); err != nil {
		// Если произошла ошибка при обновлении, возвращаем ошибку "500 Internal Server Error" с сообщением об ошибке
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если обновление прошло успешно, возвращаем JSON-ответ с сообщением "ok" и статусом "200 OK"
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
