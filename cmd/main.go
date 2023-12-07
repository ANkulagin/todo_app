package main

import (
	// Импортируем пакеты из сторонних библиотек и локальных пакетов.
	todo_app "github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/handler"
	"github.com/ANkulagin/todo-app/pkg/repository"
	"github.com/ANkulagin/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Функция main инициализирует репозиторий, сервисы и обработчики, а затем запускает HTTP-сервер.
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация конфигурации приложения.
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initconfig: %s", err.Error())
	}

	// Загрузка переменных окружения из файла .env.
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Создаем экземпляр репозитория, используя конфигурацию из файла и переменные окружения.
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)

	// Создаем экземпляр сервиса, передавая ему репозиторий.
	services := service.NewService(repos)

	// Создаем экземпляр обработчика, передавая ему сервисы.
	handlers := handler.NewHandler(services)

	// Создаем экземпляр HTTP-сервера.
	srv := new(todo_app.Server)

	// Запускаем сервер на порту из конфигурации с зарегистрированными маршрутами из обработчика.
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

// initConfig инициализирует конфигурацию, читая файл конфигурации из директории configs.
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
