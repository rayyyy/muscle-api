package funcs

import (
	"log"

	"golang.org/x/net/context"
)

// CheckLogin ユーザーがログインしているかを返す
func CheckLogin(token string) bool {
	app, err := FirebaseInit()
	if err != nil {
		return false
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("error firebaseのauthenticationの取得に失敗しました。: %v\n", err)
		return false
	}

	_, err = client.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Printf("error ユーザーの認証に失敗しました。: %v\n", err)
		return false
	}

	return true
}
