package user

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/katerji/ecommerce/envs"
	"strconv"
	"time"
)

type claims struct {
	jwt.RegisteredClaims
	User      User  `json:"user"`
	ExpiresAt int64 `json:"expires_at"`
}

func (s *Service) verifyToken(token string) (*User, error) {
	jwtSecret := envs.GetInstance().GetJWTSecret()
	return s.validateToken(token, jwtSecret)
}

func (s *Service) verifyRefreshToken(token string) (*User, error) {
	jwtSecret := envs.GetInstance().GetJWTRefreshSecret()
	return s.validateToken(token, jwtSecret)
}

func (s *Service) validateToken(token, jwtSecret string) (*User, error) {
	c := claims{}
	_, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("error parsing token")
	}
	return &c.User, nil

}

func (s *Service) createJwt(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTExpiry(),
	})
	jwtSecret := envs.GetInstance().GetJWTSecret()
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Service) createRefreshJwt(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user,
		"expires_at": getJWTRefreshExpiry(),
	})
	jwtSecret := envs.GetInstance().GetJWTRefreshSecret()
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getJWTExpiry() int64 {
	expiryString := envs.GetInstance().GetJWTExpiry()
	expiry, _ := strconv.Atoi(expiryString)
	return intToUnixTime(expiry)
}

func getJWTRefreshExpiry() int64 {
	expiryString := envs.GetInstance().GetJWTRefreshExpiry()
	expiry, _ := strconv.Atoi(expiryString)
	return intToUnixTime(expiry)
}

func intToUnixTime(num int) int64 {
	now := time.Now()
	duration := time.Duration(num) * time.Second
	result := now.Add(duration)
	return result.Unix()
}
