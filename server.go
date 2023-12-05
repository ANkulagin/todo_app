// Пакет todo_app предоставляет простую реализацию сервера для приложения по управлению задачами.

// Импорт необходимых пакетов.
package todo_app

import (
	"context"  // Пакет context определяет тип Context, который передает сроки, сигналы отмены и другие значения, связанные с запросом.
	"net/http" // Пакет http предоставляет реализации клиента и сервера HTTP.

	"time" // Пакет time предоставляет функциональность для измерения и отображения времени.
)

// Server Структура Server представляет HTTP-сервер для приложения по управлению задачами.
type Server struct {
	httpServer *http.Server // Указатель на экземпляр HTTP-сервера.
}

// Run Метод Run запускает HTTP-сервер на указанном порту.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port, // Устанавливаем адрес сервера на указанный порт.
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // Устанавливаем максимальный размер заголовка в 1 мегабайт.
		ReadTimeout:    10 * time.Second, // Устанавливаем максимальное время на чтение всего запроса.
		WriteTimeout:   10 * time.Second, // Устанавливаем максимальное время на запись ответа.
	}
	return s.httpServer.ListenAndServe() // Запускаем HTTP-сервер и возвращаем возможную ошибку.
}

// Shutdown Метод Shutdown грациозно завершает работу HTTP-сервера.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx) // Грациозно завершаем работу сервера с использованием предоставленного контекста.
}
