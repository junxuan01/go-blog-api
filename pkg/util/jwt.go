package util

import (
	"time"

	"go-blog-api/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

// 封装 JWT 相关的工具函数和常量

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 Token
func GenerateToken(userID uint, username string) (string, error) {
	cfg := config.AppConfig.JWT
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(cfg.ExpireHours) * time.Hour)

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    cfg.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(cfg.Secret))
	return token, err
}

// ParseToken 解析 Token
func ParseToken(token string) (*Claims, error) {
	secret := []byte(config.AppConfig.JWT.Secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
