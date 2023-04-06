package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"shopper"
	"shopper/pkg/repo"
	"time"
)

const (
	salt      = "mn12nc89GHJKbm1ik8GTDK14d"
	tokenTTL  = 12 * time.Hour
	signedKey = "hM1Bjas5TFn1g4FTKg1hf89NMn1caf1"
)

type AuthService struct {
	repo repo.Authorization
}

func NewAuthService(r repo.Authorization) *AuthService {
	return &AuthService{
		repo: r,
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) CreateUser(user shopper.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	fmt.Println(generatePasswordHash("admin"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signedKey))
}

func (s *AuthService) GetUser(userId int) (shopper.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("incorrect signing method")
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
