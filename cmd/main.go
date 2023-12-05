// Пакет main является точкой входа для запуска сервера to-do приложения.

// Импортируем необходимые пакеты.
package main

import (
	"github.com/ANkulagin/todo-app"             // Импортируем пакет todo_app, который содержит логику сервера.
	"github.com/ANkulagin/todo-app/pkg/handler" // Импортируем пакет handler, содержащий обработчики запросов.
	"log"                                       // Импортируем пакет log для вывода логов.
)

// Функция main - точка входа для приложения.
func main() {
	handlers := new(handler.Handler) // Создаем экземпляр обработчика запросов.
	srv := new(todo_app.Server)      // Создаем экземпляр HTTP-сервера.

	// Запускаем сервер на порту 8000 с зарегистрированными маршрутами из обработчика.
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server : %s", err.Error()) // Выводим сообщение об ошибке при запуске сервера.
	}
}
