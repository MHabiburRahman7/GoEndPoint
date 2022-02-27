package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type ArticleData struct{
  Title string `json:"Title"`
  Desc string `json:"Description"`
  Content string `json:"Content"`
}

var ArticleDatas []ArticleData

func returnAllArticles(w http.ResponseWriter, r *http.Request){
  fmt.Println("Well, hit again")
  json.NewEncoder(w).Encode(ArticleDatas)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/articles", returnAllArticles)
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleRequestsUsingMux(){
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func initArticles(){
  ArticleDatas = []ArticleData{
        ArticleData{Title: "Hello World", Desc: "Article Description 1", Content: "Content"},
        ArticleData{Title: "Hello Again", Desc: "Article Description 2", Content: "Content"},
    }
}

func main() {
    initArticles()
    //handleRequests()
    handleRequestsUsingMux()
}
