package controller

import (
    "fmt"
    "net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the LoginPage!")
    fmt.Println("Endpoint Hit: LoginPage")
}