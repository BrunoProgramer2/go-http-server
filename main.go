package main

import (
	"net/http"

	"os"

	"BrunoProgramer2.github.io/go-http-server/pkg"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	http.ListenAndServe(port, pkg.SetupRouter())
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}
