package funcs

import (
	"os"
)

// GetEnv ENVを取得する、ない場合はデフォルトを返す
func GetEnv(key string, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}