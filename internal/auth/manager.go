package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
)

type TokenManager interface {
	GenerateAccessToken(id string, isAdmin bool, ttl time.Duration) (string, error)
	GenerateRefreshToken(id string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
}

type Manager struct {
	Salt       string
	SigningKey string
}

func NewManager(salt string, signingKey string) *Manager {
	return &Manager{Salt: salt, SigningKey: signingKey}
}

func (m *Manager) GenerateAccessToken(id string, isAdmin bool, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(ttl).Unix(),
		"sub":   id,
		"admin": isAdmin,
	})

	return token.SignedString([]byte(m.SigningKey))
}

func (m *Manager) GenerateRefreshToken(id string, ttl time.Duration) (string, error) {
	t := time.Now().Add(ttl).Unix()
	tStr := strconv.Itoa(int(t))

	tStr = base64.StdEncoding.EncodeToString([]byte(tStr))
	userId := base64.StdEncoding.EncodeToString([]byte(id))

	// TODO: base64 to sha-256
	secret := base64.StdEncoding.EncodeToString([]byte(m.SigningKey))

	arr := []string{tStr, userId, secret}
	token := strings.Join(arr, ".")

	return fmt.Sprintf(token), nil
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
