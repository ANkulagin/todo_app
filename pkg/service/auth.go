package service

import (
	"crypto/sha1"
	"errors"
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
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// Пытаемся разобрать токен с использованием метода ParseWithClaims из библиотеки jwt
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи токена является HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		// Возвращаем ключ подписи для проверки подписи токена
		return []byte(signingKey), nil
	})
	if err != nil {
		// Если произошла ошибка при разборе токена, возвращаем ошибку
		return 0, err
	}

	// Проверяем, что утверждения токена имеют тип *tokenClaims
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	// Возвращаем идентификатор пользователя из утверждений токена
	return claims.UserId, nil
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
