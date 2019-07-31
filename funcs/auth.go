package funcs

import (
	"log"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// CheckLogin ユーザーがログインしているかを返す
func CheckLogin(token string) bool {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error firebaseの初期化に失敗しました。: %v", err)
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
