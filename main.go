package main

import (
	"fmt"
	"com.pos.api/app/controller"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("Run Server on : localhost:10000")
	myRouter := mux.NewRouter().StrictSlash(true)

	controller.Articles = [] controller.Article {
        controller.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        controller.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }

    myRouter.HandleFunc("/", controller.HomePage)
	myRouter.HandleFunc("/login", controller.LoginPage)
    myRouter.HandleFunc("/articles", controller.ReturnAllArticles)
    myRouter.HandleFunc("/article", controller.CreateNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", controller.DeleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", controller.ReturnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}