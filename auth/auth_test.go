package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	accessTokens := []string{}
	for i := 0; i < 100; i++ {
		accessToken := GenerateAccessToken()
		assert.Equal(t, -1, core.IndexOf(accessTokens, accessToken))
		assert.Equal(t, 128, len(accessToken))
		accessTokens = append(accessTokens, accessToken)
	}
}

func TestGeneratePassword(t *testing.T) {
	passwords := []string{}
	for i := 0; i < 100; i++ {
		password := GeneratePassword(8)
		assert.Equal(t, -1, core.IndexOf(passwords, password))
		assert.Equal(t, 8, len(password))
		passwords = append(passwords, password)
	}
}

func TestOTP(t *testing.T) {
	for i := 1; i < 100; i++ {
		otp := GenerateOTPCode(6)
		time.Sleep(1 * time.Microsecond)
		assert.Less(t, 100000, otp)
		assert.Greater(t, 999999, otp)
	}
}
