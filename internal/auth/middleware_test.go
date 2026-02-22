package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func createTestToken(secret, username string, exp time.Time) string {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      now.Unix(),
		"exp":      exp.Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestMiddleware(t *testing.T) {
	secret := "test-secret-key"

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{
			name:           "missing authorization header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid format - no Bearer prefix",
			authHeader:     "token123",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid format - wrong prefix",
			authHeader:     "Basic token123",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid format - empty token",
			authHeader:     "Bearer ",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid format - null token",
			authHeader:     "Bearer null",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "invalid token",
			authHeader:     "Bearer invalid-token",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			recorder := httptest.NewRecorder()
			middleware := Middleware(secret)(handler)
			middleware.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}

func TestMiddlewareValidToken(t *testing.T) {
	secret := "test-secret-key"
	token := createTestToken(secret, "testuser", time.Now().Add(1*time.Hour))

	var capturedUsername string
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedUsername = UsernameFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()
	middleware := Middleware(secret)(handler)
	middleware.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "testuser", capturedUsername)
}

func TestMiddlewareWrongSecret(t *testing.T) {
	token := createTestToken("wrong-secret", "testuser", time.Now().Add(1*time.Hour))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()
	middleware := Middleware("correct-secret")(handler)
	middleware.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}

func TestMiddlewareExpiredToken(t *testing.T) {
	secret := "test-secret-key"
	expiredToken := createTestToken(secret, "testuser", time.Now().Add(-1*time.Hour))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+expiredToken)

	recorder := httptest.NewRecorder()
	middleware := Middleware(secret)(handler)
	middleware.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}

func TestUsernameFromContext(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			name:     "empty context",
			ctx:      context.Background(),
			expected: "",
		},
		{
			name:     "context with username",
			ctx:      context.WithValue(context.Background(), usernameKey, "testuser"),
			expected: "testuser",
		},
		{
			name:     "context with wrong type",
			ctx:      context.WithValue(context.Background(), usernameKey, 123),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UsernameFromContext(tt.ctx)
			assert.Equal(t, tt.expected, result)
		})
	}
}
