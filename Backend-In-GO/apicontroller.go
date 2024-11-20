package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getScores(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("score")

	if name == "dailyscores" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData := []map[string]interface{}{}
		temp := dailyScorePrioQueue.front
		for temp != nil {
			jsonData = append(jsonData, map[string]interface{}{
				"name":  temp.item.username,
				"score": temp.item.score,
			},
			)
			temp = temp.next
		}

		err := json.NewEncoder(w).Encode(jsonData)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	} else if name == "halloffame" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData := []map[string]interface{}{}
		temp := hallOfFamePrioQueue.front
		for temp != nil {
			jsonData = append(jsonData, map[string]interface{}{
				"name":  temp.item.username,
				"score": temp.item.score,
			},
			)
			temp = temp.next
		}

		err := json.NewEncoder(w).Encode(jsonData)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}

func createAccount(w http.ResponseWriter, r *http.Request) {
	//temp struct to hold the json object values
	var data map[string]interface{}
	//always assume nothing can go wrong hurrrrr
	json.NewDecoder(r.Body).Decode(&data)

	err := userDBButItsATree.createUser(data["username"].(string), data["password"].(string), 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, err.Error())
	} else {
		//created successfully
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Account Created!")
	}
}

func updateScore(w http.ResponseWriter, r *http.Request) {

}
