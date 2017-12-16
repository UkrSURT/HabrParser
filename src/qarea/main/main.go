package main

import (
	"os"
	"qarea/web"
	"qarea/middleware"
    "log"
    "net/http"
    "fmt"
)

func main() {
	args := os.Args
	if len(args) != 0 {
		str := args[1]
		web.GrabArticles()
		middleware.FindWord(str)
	}
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, middleware.Shared)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

