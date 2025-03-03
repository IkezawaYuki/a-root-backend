package service

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/interface/dto/req"
	"IkezawaYuki/a-root-backend/interface/repository"
	"fmt"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"strings"
	"time"
)

type authService struct {
	redisClient repository.RedisRepository
}

type AuthService interface {
	IsCustomerIsLogin(tokenString string) (int, error)
	IsAdminLogin(tokenString string) (int, error)
	CheckPassword(user req.User, password string) error
	GenerateJWTAdmin(admin *model.Admin) (string, error)
	GenerateJWTCustomer(c *model.Customer) (string, error)
}

func NewAuthService(redisClient repository.RedisRepository) AuthService {
	return &authService{
		redisClient: redisClient,
	}
}

func (a *authService) IsCustomerIsLogin(tokenString string) (int, error) {
	slog.Info("IsCustomerIsLogin is invoked")
	slog.Info(tokenString)
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.AccessSecretKey), nil
	})
	if err != nil {
		slog.Info(err.Error())
		return 0, arootErr.ErrAuthorization
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, arootErr.ErrAuthorization
	}
	if !claims.VerifyAudience("customer", true) {
		return 0, arootErr.ErrAuthentication
	}

	return int(claims["sub"].(float64)), nil
}

func (a *authService) IsAdminLogin(tokenString string) (int, error) {
	slog.Info("IsAdminLogin is invoked")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.AccessSecretKey), nil
	})
	if err != nil {
		slog.Info(err.Error())
		return 0, arootErr.ErrAuthorization
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, arootErr.ErrAuthorization
	}
	if !claims.VerifyAudience("admin", true) {
		return 0, arootErr.ErrAuthentication
	}
	return int(claims["sub"].(float64)), nil
}

func (a *authService) GenerateJWTCustomer(c *model.Customer) (string, error) {
	claims := jwt.MapClaims{
		"iss":   "popple",
		"aud":   "customer",
		"sub":   c.ID,
		"email": c.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.AccessSecretKey))
}

func (a *authService) GenerateJWTAdmin(admin *model.Admin) (string, error) {
	claims := jwt.MapClaims{
		"iss":   "popple",
		"aud":   "admin",
		"sub":   admin.ID,
		"email": admin.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.AccessSecretKey))
}

func (a *authService) CheckPassword(user req.User, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		return fmt.Errorf("password is incorrect: %s, %v", err.Error(), arootErr.ErrAuthorization)
	}
	return nil
}
