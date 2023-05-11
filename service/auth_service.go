package service

import (
	"errors"
	"final-project/model"
	"final-project/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	tokenTTL = 5 * time.Minute
	tokenKey = "your-secret-key"
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
		"username":     user.Username,
		"access_token": true,
		"exp":          time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(tokenKey))
}

func (r *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(tokenKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims)["access_token"]; !ok {
		return "", errors.New("invalid token type")
	}

	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	return claims["username"].(string), nil
}

func (r *AuthService) RefreshToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(tokenKey), nil
	})
	if err != nil {
		return "", err
	}

	if _, ok := token.Claims.(jwt.MapClaims)["access_token"]; !ok {
		return "", errors.New("invalid token type")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":     username,
		"access_token": true,
		"exp":          time.Now().Add(tokenTTL).Unix(),
	})

	newTokenString, err := newToken.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}
