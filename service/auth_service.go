package service

import (
	"errors"
	"final-project/model"
	"final-project/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	tokenTTL = 12 * time.Hour    // время жизни токена
	tokenKey = "your-secret-key" // ключ для подписи токена
)

type AuthService struct {
	authRepo repository.AuthorizationRepo
}

func NewAuthService(authRepo repository.AuthorizationRepo) *AuthService {
	return &AuthService{authRepo}
}

func (r *AuthService) RegisterUser(user model.User) error {

	err := r.authRepo.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := r.authRepo.Login(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(tokenKey))
}

func (r *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи токена соответствует нашему методу
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(tokenKey), nil
	})

	// Если токен не валиден или истекло время жизни
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Получаем и возвращаем имя пользователя
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	return claims["username"].(string), nil
}
