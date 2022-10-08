package main

import (
	"fmt"
	"net/http"

	"BrunoProgramer2.github.io/go-http-server/pkg"
)

func main() {
	fmt.Print("Server on http://localhost:3000")
	http.ListenAndServe(":3000", pkg.SetupRouter())
}
