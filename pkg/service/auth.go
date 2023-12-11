package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/ANkulagin/todo-app"
	"github.com/ANkulagin/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Константа salt используется для усиления безопасности при хешировании паролей.
const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "adsfjklasdfl"
	tokenTTL   = 12 * time.Hour
)

// Структура AuthService представляет сервис для работы с авторизацией пользователей.
type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user-id"`
}

// NewAuthService создает новый экземпляр AuthService с переданным репозиторием.
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser создает нового пользователя, хеширует его пароль и сохраняет в базу данных через репозиторий.
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	// Хешируем пароль пользователя перед сохранением в базу данных.
	user.Password = generatePasswordHash(user.Password)
	// Вызываем метод CreateUser у репозитория для сохранения пользователя в базе данных.
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

// generatePasswordHash хеширует переданный пароль, используя sha1 и добавляя к нему соль.
func generatePasswordHash(password string) string {
	// Инициализируем новый хеш sha1.
	hash := sha1.New()

	// Записываем в хеш байтовое представление пароля.
	hash.Write([]byte(password))

	// Форматируем хеш в строку и добавляем соль.
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
