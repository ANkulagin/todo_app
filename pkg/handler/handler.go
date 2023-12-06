package handler

// Пакет handler предоставляет обработчики запросов и метод для инициализации маршрутов веб-приложения.

// Импортируем необходимые пакеты.
import (
	"github.com/ANkulagin/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
) // Импортируем пакет gin для работы с HTTP-запросами и ответами.

// Структура Handler представляет собой обработчик запросов.

// Структура Handler содержит указатель на сервисы приложения.
type Handler struct {
	services *service.Service
}

// Метод NewHandler создает новый экземпляр структуры Handler с переданными сервисами

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Метод InitRoutes инициализирует маршруты веб-приложения и возвращает движок gin.Engine.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // Создаем новый экземпляр движка Gin.

	// Группа маршрутов для авторизации.
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) // Регистрация маршрута для создания нового пользователя.
		auth.POST("/sign-in", h.signIn) // Регистрация маршрута для входа в систему.
	}

	// Группа маршрутов для работы с задачами и списками.
	api := router.Group("/api")
	{
		// Группа маршрутов для управления списками задач.
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)      // Регистрация маршрута для создания нового списка задач.
			lists.GET("/", h.getAllLists)      // Регистрация маршрута для получения всех списков задач.
			lists.GET("/:id", h.getListById)   // Регистрация маршрута для получения списка задач по идентификатору.
			lists.PUT("/:id", h.updateList)    // Регистрация маршрута для обновления списка задач.
			lists.DELETE("/:id", h.deleteList) // Регистрация маршрута для удаления списка задач.

			// Группа маршрутов для управления задачами внутри списка.
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem) // Регистрация маршрута для создания новой задачи в списке.
				items.GET("/", h.getAllItems) // Регистрация маршрута для получения всех задач в списке.
			}
		}

		// Группа маршрутов для управления отдельными задачами.
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)   // Регистрация маршрута для получения задачи по идентификатору.
			items.PUT("/:id", h.updateItem)    // Регистрация маршрута для обновления задачи.
			items.DELETE("/:id", h.deleteItem) // Регистрация маршрута для удаления задачи.
		}
	}

	return router // Возвращаем настроенный движок Gin с зарегистрированными маршрутами.
}
