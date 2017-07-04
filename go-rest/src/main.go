package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct{
	Id int `json: "ID"`
	Title string `json: "Title"`
	Desc string `json: "desc"`
	Content string `json: "Content"`
}

type Articles []Article

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the index")
	fmt.Println("Endpoint Hit: index")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", index)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))	
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Id: 1, Title: "Hello", Desc:"Article Description", Content:"Article Content"},
		Article{Id: 2, Title: "Article2", Desc:"Article thow", Content: "Article content 2"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprint(w, "Key: " + key)
	
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}