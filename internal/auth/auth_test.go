package auth

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		salt     string
		wantErr  bool
	}{
		{
			name:     "valid password and salt",
			password: "testpassword",
			salt:     "0102030405060708090a0b0c0d0e0f10",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			salt:     "0102030405060708090a0b0c0d0e0f10",
			wantErr:  false,
		},
		{
			name:     "unicode password",
			password: "пароль123",
			salt:     "0102030405060708090a0b0c0d0e0f10",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password, tt.salt)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotEmpty(t, hash)
			_, decodeErr := hex.DecodeString(hash)
			assert.NoError(t, decodeErr, "hash should be valid hex string")
			assert.Len(t, hash, 128, "hash should be 128 hex chars (64 bytes)")
		})
	}
}

func TestHashPasswordDeterministic(t *testing.T) {
	password := "testpassword"
	salt := "0102030405060708090a0b0c0d0e0f10"

	hash1, err := HashPassword(password, salt)
	require.NoError(t, err)

	hash2, err := HashPassword(password, salt)
	require.NoError(t, err)

	assert.Equal(t, hash1, hash2, "same password and salt should produce same hash")
}

func TestHashPasswordKnownValue(t *testing.T) {
	password := "testpassword"
	salt := "0102030405060708090a0b0c0d0e0f10"
	expected := "404ba06bdb03dc9a8a9ad7ea8e1f13a58d0c4a2a600580bf9ac558147c20afd960e7300e8ce8d0874dbd6be8cf4147caf07182787e468001f06d17df9b7e42b5"

	hash, err := HashPassword(password, salt)
	require.NoError(t, err)

	assert.Equal(t, expected, hash)
}

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()
	require.NoError(t, err)
	assert.NotEmpty(t, salt)

	_, decodeErr := hex.DecodeString(salt)
	assert.NoError(t, decodeErr, "salt should be valid hex string")

	assert.Len(t, salt, 32, "salt should be 32 hex chars (16 bytes)")
}

func TestGenerateSaltUniqueness(t *testing.T) {
	salts := make(map[string]bool)
	for range 100 {
		salt, err := GenerateSalt()
		require.NoError(t, err)
		assert.False(t, salts[salt], "generated salt should be unique")
		salts[salt] = true
	}
}
