package funcs

import (
	"log"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// FirebaseInit firebaseを初期化
func FirebaseInit() (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: "muscle-2710.appspot.com",
	}

	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Printf("error firebaseの初期化に失敗しました。: %v", err)
		return nil, err
	}
	return app, nil
}
