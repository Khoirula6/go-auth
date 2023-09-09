package main

import (
	"net/http"

	authController "github.com/Khoirula6/go-auth/controllers"
)

func main() {
	http.HandlerFunc("/", authController.Index)

}
