package auth

import (
	"context"

	firebase "firebase.google.com/go/auth"
	"github.com/ez-connect/go-rest/core"
)

func GetFirebaseUser(accessToken string) (*firebase.UserRecord, error) {
	client, err := core.GetFirebase().Auth(context.Background())
	if err != nil {
		return nil, err
	}

	// Verify firebase token
	token, err := client.VerifyIDToken(context.Background(), accessToken)
	if err != nil {
		return nil, err
	}

	// Get user
	user, err := client.GetUser(context.TODO(), token.UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
