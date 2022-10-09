package main

import (
	"net/http"

	"os"

	"BrunoProgramer2.github.io/go-http-server/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	http.ListenAndServe(port, pkg.SetupRouter())
}
