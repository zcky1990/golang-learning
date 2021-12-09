package controller

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
	"github.com/gorilla/mux"
     // "com.pos.api/app/models"
)

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
    fmt.Println(r)
    fmt.Println(r.Header["User-Agent"])
    fmt.Println(r.RequestURI)

    //get raw query string
    keys := r.URL.RawQuery
    
    fmt.Println(keys)
   
    //get looping query string
    values := r.URL.Query()
    for k, v := range values {
        fmt.Println(k, " => ", v)
    }

    //get params from url /{id} query string
    vars := mux.Vars(r)
    fmt.Println(vars)
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