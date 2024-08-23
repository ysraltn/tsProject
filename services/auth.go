package services

import (
	"database/sql"
	"errors"
	"tsProject/models"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepositories models.IUserRepository
}

var jwtKey = []byte("my_secret_key") // Güçlü ve gizli bir anahtar kullanmalısınız.

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func NewAuthService(userRepositories models.IUserRepository) models.IAuthService {
	return &AuthService{
		userRepositories: userRepositories,
	}
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	user, err := s.userRepositories.GetByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", errors.New("user not found")
	}
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("user not found")
	}
	if !CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}
	token, err := GenerateJWT(username, user.Role)
	if err != nil {
		return "", "", errors.New("token generation error")
	}
	return token, user.Role, nil
}

func (s *AuthService) Register(username, password string) error {
	_, err := s.userRepositories.GetByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			return err
		}
		user := models.User{
			Username: username,
			Password: hashedPassword,
			Role:     "admin",
		}
		s.userRepositories.Add(&user)
		return nil
	}
	return errors.New("username taken")
}

func CheckPasswordHash(password, hashedPassword string) bool {
	// bcrypt.CompareHashAndPassword hash'lenmiş şifre ile düz metin şifreyi karşılaştırır
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token 24 saat geçerli
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
