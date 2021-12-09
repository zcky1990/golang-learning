package controller

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
	"github.com/gorilla/mux"
    // "com.pos.api/app/models"
)

// // Article - Our struct for all articles
type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    Articles = append(Articles, article)
    json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}