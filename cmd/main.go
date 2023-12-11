// Определяем пакет main, который является точкой входа в приложение.
package main

// Импортируем необходимые пакеты.
import (
	"github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/handler"
	"github.com/ANkulagin/todo-app/pkg/repository"
	"github.com/ANkulagin/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Функция main - точка входа в приложение.
func main() {
	// Устанавливаем форматтер логов в JSON.
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализируем конфигурацию приложения.
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка при инициализации конфигурации: %s", err.Error())
	}

	// Загружаем переменные окружения из файла .env.
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка при загрузке переменных окружения: %s", err.Error())
	}

	// Инициализируем подключение к базе данных PostgreSQL.
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("ошибка при инициализации базы данных: %s", err.Error())
	}

	// Создаем репозиторий для взаимодействия с базой данных.
	repos := repository.NewRepository(db)

	// Создаем сервис для обработки бизнес-логики.
	services := service.NewService(repos)

	// Создаем обработчик HTTP-запросов.
	handlers := handler.NewHandler(services)

	// Создаем экземпляр сервера TODO и запускаем его.
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("ошибка при запуске HTTP-сервера: %s", err.Error())
	}
}

// Функция initConfig инициализирует конфигурацию приложения, используя файл "config.yaml" в папке "configs".
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
