package main

import (
	"os"

	"github.com/idkwhyureadthis/auth-service/pkg/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	connUrl := os.Getenv("DB_URL")
	secretKey := os.Getenv("SECRET")
	h := handler.New(connUrl, []byte(secretKey))

	h.Start(port)
}
