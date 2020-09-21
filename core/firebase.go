package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type IDTokenResp struct {
	Token        string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	IsNewUser    bool   `json:"isNewUser"`
}

///////////////////////////////////////////////////////////////////

var firebaseApp *firebase.App

func InitFirebase(credentialsFile string) error {
	var err error
	opt := option.WithCredentialsFile(credentialsFile)
	firebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	return err
}

func GetFirebase() *firebase.App {
	return firebaseApp
}

///////////////////////////////////////////////////////////////////

func GetFirebaseIDToken(apiKey, customToken string) (*IDTokenResp, error) {
	url := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s", apiKey)
	data := map[string]interface{}{
		"token":             customToken,
		"returnSecureToken": true,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	token := IDTokenResp{}
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
