package main

import (
	"fmt"
	"net/http"
	"sync"
)

// init fake db
var (
	userDBButItsATree = bst{}
	hallOfFamePrioQueue = prioQueue{}
	dailyScorePrioQueue = prioQueue{}
	mu sync.Mutex
)

func main() {
	//http package is concurrent by default sooooooo
	http.HandleFunc("/api/scores", getScores)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("uh oh, server boom")
		fmt.Println(err)
	}
}
