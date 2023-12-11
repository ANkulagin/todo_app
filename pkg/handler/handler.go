package handler

// Пакет handler предоставляет обработчики запросов и метод для инициализации маршрутов веб-приложения.

// Импортируем необходимые пакеты.
import (
	"github.com/ANkulagin/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
) // Импортируем пакет gin для работы с HTTP-запросами и ответами.

// Структура Handler представляет собой обработчик запросов.

// Структура Handler содержит сервисы, необходимые для обработки запросов.
type Handler struct {
	services *service.Service
}

// NewHandler создает новый экземпляр обработчика с переданными сервисами.
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes инициализирует и возвращает маршруты для обработчика.
func (h *Handler) InitRoutes() *gin.Engine {
	// Создаем новый маршрутизатор Gin.
	router := gin.New()

	// Группа маршрутов для авторизации.
	auth := router.Group("/auth")
	{
		// Регистрация нового пользователя.
		auth.POST("/sign-up", h.signUp)

		// Вход пользователя в систему.
		auth.POST("/sign-in", h.signIn)
	}

	// Группа маршрутов для API.
	api := router.Group("/api")
	{
		// Группа маршрутов для списков.
		lists := api.Group("/lists")
		{
			// Создание нового списка.
			lists.POST("/", h.createList)

			// Получение всех списков.
			lists.GET("/", h.getAllLists)

			// Получение списка по идентификатору.
			lists.GET("/:id", h.getListById)

			// Обновление списка.
			lists.PUT("/:id", h.updateList)

			// Удаление списка.
			lists.DELETE("/:id", h.deleteList)

			// Группа маршрутов для элементов списка.
			items := lists.Group(":id/items")
			{
				// Создание нового элемента списка.
				items.POST("/", h.createItem)

				// Получение всех элементов списка.
				items.GET("/", h.getAllItems)

				// Получение элемента списка по идентификатору.
				items.GET("/:item_id", h.getItemById)

				// Обновление элемента списка.
				items.PUT("/:item_id", h.updateItem)

				// Удаление элемента списка.
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	// Возвращаем маршрутизатор с настроенными маршрутами.
	return router
}
