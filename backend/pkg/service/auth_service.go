package service

import (
	"context"
	"fmt"
	"ruang-arah/backend/config"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/repository"
	"ruang-arah/backend/security"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Role   string `json:"role"`
	UserId int32  `json:"user_id"`
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepo,
	}
}

func (service *AuthServiceImpl) GenerateToken(request web.LoginRequest) (web.LoginResponse, error) {
	userDomain, err := service.authRepository.FindUser(
		request.Email, security.GeneratePasswordHash(request.Password),
	)
	helper.PanicIfError(err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TOKEN_EXPIRES).Unix(),
			Issuer:    config.APP_NAME,
			IssuedAt:  time.Now().Unix(),
		},
		Role:   userDomain.Role,
		UserId: userDomain.Id,
	})

	tokenString, err := token.SignedString([]byte(config.SECRET_KEY))
	helper.PanicIfError(err)

	return helper.ToLoginResponse(userDomain, tokenString), nil
}

func (service *AuthServiceImpl) Register(request web.RegisterRequest, email string) (web.RegisterResponse, error) {
	hashed := security.GeneratePasswordHash(request.Password)
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")

	userDomain := domain.UserDomain{
		Username:  request.Username,
		Email:     "",
		Password:  hashed,
		Role:      "guest",
		CreatedAt: time.Now().In(timeLoc),
		UpdatedAt: time.Now().In(timeLoc),
	}

	authRepo, err := service.authRepository.Save(userDomain, email)
	helper.PanicIfError(err)

	return helper.ToRegisterResponse(authRepo, email), nil
}

func ParseToken(ctx context.Context, token string) (int32, string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	newCtx := context.WithValue(ctx, "role", claims.Role)
	ctx = context.WithValue(newCtx, "user_id", claims.UserId)
	ctx = context.WithValue(newCtx, "props", claims)

	return claims.UserId, claims.Role, nil
}
