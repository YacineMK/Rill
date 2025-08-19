package jwt

import (
	"errors"
	"time"

	"github.com/YacineMK/Rill/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.GetConfig()

var secretKey = []byte(cfg.JWT.Secret)
var tokenDuration = cfg.JWT.DurationMinutes

type Claims struct {
	StreamID string `json:"stream_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(streamID string) (string, error) {
	claims := &Claims{
		StreamID: streamID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenDuration) * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func DecodeJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
