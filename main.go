//https://tutorialedge.net/golang/creating-restful-api-with-golang/

package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
)

type ArticleData struct{
  Id string `json:"Id"`
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range ArticleDatas {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func handleRequestsUsingMux(){
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    //put
    myRouter.HandleFunc("/article", updateArticleData).Methods("PUT")
    //delete
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    //post or create
    myRouter.HandleFunc("/article-echo", createNewArticleDataAndEcho).Methods("POST")
    myRouter.HandleFunc("/article-append", createNewArticleDataAndApped).Methods("POST")

    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewArticleDataAndEcho(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprint(w, "%+v", string(reqBody))
}

func createNewArticleDataAndApped(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var articleData ArticleData
    json.Unmarshal(reqBody, &articleData)
    //update global ArticleDatas and append the new article
    ArticleDatas = append(ArticleDatas, articleData)

    json.NewEncoder(w).Encode(ArticleDatas)
}

func updateArticleData(w http.ResponseWriter, r *http.Request){
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var articleData ArticleData
    json.Unmarshal(reqBody, &articleData)

    id := articleData.Id
    // we then need to loop through all our articles
    for index, article := range ArticleDatas {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the
            // article
            ArticleDatas = append(ArticleDatas[:index], ArticleDatas[index+1:]...)

            //attach the new articleData
            ArticleDatas = append(ArticleDatas, articleData)
            fmt.Println("Updated")
        }
    }
}

func deleteArticle(w http.ResponseWriter, r *http.Request){
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    fmt.Println("this is vars %v", vars)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    fmt.Println("this is id %v", id)
    // we then need to loop through all our articles
    for index, article := range ArticleDatas {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the
            // article
            ArticleDatas = append(ArticleDatas[:index], ArticleDatas[index+1:]...)

            fmt.Println("Deleted")
        }
    }
}

func initArticles(){
  ArticleDatas = []ArticleData{
        ArticleData{Id: "0", Title: "Hello World", Desc: "Article Description 1", Content: "Content"},
        ArticleData{Id: "1", Title: "Hello Again", Desc: "Article Description 2", Content: "Content"},
    }
}

func main() {
    initArticles()
    //handleRequests()
    handleRequestsUsingMux()
}
