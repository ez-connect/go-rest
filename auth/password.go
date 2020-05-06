package auth

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"time"
)

// Returns SHA256 of a string
func GetSHA256(value string) string {
	// sum := sha256.Sum256([]byte(value))
	// return fmt.Sprintf("%x", sum)
	h := sha256.New()

	if _, err := h.Write([]byte(value)); err != nil {
		fmt.Println(err)
		return value
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Total: (SHA-256 of nano secs = 256 bits) + (32 random bytes) = 64 bytes
func GenerateAccessToken() string {
	var accessToken string
	uid := bytes.Buffer{}
	err := binary.Write(&uid, binary.LittleEndian, time.Now().UnixNano())
	if err != nil {
		fmt.Println(err)
		accessToken = uid.String()
	}

	random := make([]byte, 32)
	if _, err := rand.Read(random); err == nil {
		h := sha256.New()
		if _, err := h.Write(uid.Bytes()); err == nil {
			accessToken = fmt.Sprintf("%x", h.Sum(random))
		}
	}

	return accessToken
}

// Generate a random password with a specific length.
// It returns an empty string if there is any error.
func GeneratePassword(length int) string {
	buff := make([]byte, length)
	if _, err := rand.Read(buff); err != nil {
		return ""
	}

	str := base64.StdEncoding.EncodeToString(buff)
	// Base 64 can be longer than len
	return str[:length]
}
