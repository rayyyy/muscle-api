package funcs

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

// UploadImage 画像をfirebaseにアップロードする
func UploadImage(filename string, base64Image string, prefixPath string) bool {
	app, err := FirebaseInit()
	if err != nil {
		return false
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
		return false
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	// TODO: 変数で受け取るようにする
	contentType := "image/png"
	ctx := context.Background()

	writer := bucket.Object(prefixPath + filename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	src, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	newFile, err := os.Create("tmp/" + filename)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	defer newFile.Close()

	if _, err := newFile.Write(src); err != nil {
		log.Fatalln(err)
		return false
	}

	newFile.Close()

	f, err := os.Open("tmp/" + filename)
	defer f.Close()

	if _, err = io.Copy(writer, f); err != nil {
		log.Fatalln(err)
		return false
	}

	if err := writer.Close(); err != nil {
		log.Fatalln(err)
		return false
	}

	if err := os.Remove("tmp/" + filename); err != nil {
		log.Fatalln(err)
	}
	return true
}

// UploadProfileIcon プロフィール画像アップロード
func UploadProfileIcon(filename string, base64Image string) bool {
	return UploadImage(filename, strings.Split(base64Image, "base64,")[1], "user-icon/")
}

// UploadPlanImage メンタープラン画像アップロード
func UploadPlanImage(filename string, base64Image string) bool {
	return UploadImage(filename, strings.Split(base64Image, "base64,")[1], "mentor-image/")
}
