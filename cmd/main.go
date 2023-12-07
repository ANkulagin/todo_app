package main

// Функция main является точкой входа в приложение.

// Импортируем необходимые пакеты.
import (
	todo_app "github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/handler"
	"github.com/ANkulagin/todo-app/pkg/repository"
	"github.com/ANkulagin/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

// Функция main инициализирует репозиторий, сервисы и обработчики, а затем запускает HTTP-сервер.
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initconfig :%s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("errror loading env variables: %s", err.Error())
	}
	// Создаем экземпляр репозитория.
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)

	// Создаем экземпляр сервиса, передавая ему репозиторий.
	services := service.NewService(repos)

	// Создаем экземпляр обработчика, передавая ему сервисы.
	handlers := handler.NewHandler(services)

	// Создаем экземпляр HTTP-сервера.
	srv := new(todo_app.Server)

	// Запускаем сервер на порту 8000 с зарегистрированными маршрутами из обработчика.
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error()) // Выводим сообщение об ошибке при запуске сервера.
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
