package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	userDBButItsATree = bst{}
	hallOfFamePrioQueue = prioQueue{}
	dailyScorePrioQueue = prioQueue{}
	mu sync.Mutex
)

func main() {
	//init fake db

	//dont need do this as listen and serve if taking in nil will give default
	//but i want prefix
	mux := http.NewServeMux()

	//http package is concurrent by default sooooooo
	mux.HandleFunc("GET /scores", getScores)
	mux.HandleFunc("POST /login", login)
	mux.HandleFunc("POST /create-new-account", createAccount)
	mux.HandleFunc("PUT /{username}/update-score", updateScore)

	//default server will redirect to mux i think
	//need /api/ instead of just /api
	http.Handle("/api/", http.StripPrefix("/api", mux))

	fmt.Println("Server running on localhost:8080")

	//when nil will default serve mux (wtv that means heh)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("uh oh, server boom")
		fmt.Println(err)
	}
}
