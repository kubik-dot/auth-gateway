package authgateway

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateRSAKeyPair generates a new RSA key pair
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	return privateKey, &privateKey.PublicKey, nil
}

// GenerateJWTToken generates a new JWT token with the given payload
func GenerateJWTToken(payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	return tokenString, nil
}

func GenerateSecretKey() (string, error) {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return "", fmt.Errorf("failed to generate secret key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(secret), nil
}

func GenerateCookieID() string {
	return uuid.New().String()
}

func GenerateRandomToken() string {
	return base64.StdEncoding.EncodeToString(make([]byte, 16))
}

const (
	jwtKey       = []byte("your-jwt-secret-key")
	jwtExpiration = time.Hour * 24 * 30
)