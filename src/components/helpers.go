package auth_gateway

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func sha256Hash(data string) string {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func generateCSRFToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func validateCSRFToken(token string, storedToken string) bool {
	return token == storedToken
}

func generateCookieHeader(header string) string {
	hash := sha256Hash(header)
	cookie := fmt.Sprintf("XSRF-TOKEN=%s; Path=/; HttpOnly; Secure", hash)
	return cookie
}

func getCSRFTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("XSRF-TOKEN")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func validateCSRFTokenInRequest(r *http.Request) error {
	storedToken, err := getCSRFTokenFromCookie(r)
	if err != nil {
		return err
	}
	tokenHeader := r.Header.Get("X-XSRF-TOKEN")
	if !validateCSRFToken(tokenHeader, storedToken) {
		return errors.New("CSRF token mismatch")
	}
	return nil
}

func generateSessionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func getCSRFToken(r *http.Request) (string, error) {
	storedToken, err := getCSRFTokenFromCookie(r)
	if err != nil {
		return "", err
	}
	tokenHeader := r.Header.Get("X-XSRF-TOKEN")
	if !validateCSRFToken(tokenHeader, storedToken) {
		return "", errors.New("CSRF token mismatch")
	}
	return tokenHeader, nil
}

func getHeaderFromRequest(r *http.Request, name string) (string, error) {
	if r.Header == nil {
		return "", errors.New("header is nil")
	}
	if value := r.Header.Get(name); value != "" {
		return value, nil
	}
	return "", errors.New("header not found")
}

func getPostFormValue(r *http.Request, name string) (string, error) {
	if r.Form == nil {
		return "", errors.New("form is nil")
	}
	if value := r.FormValue(name); value != "" {
		return value, nil
	}
	return "", errors.New("form value not found")
}

func getPostJSON(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}