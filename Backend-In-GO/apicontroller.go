package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getScores(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("score")
	jsonData := []map[string]interface{}{
		{
			"name":  "hehe",
			"score": 1234,
		},
	}
	jsonData = append(jsonData, map[string]interface{}{
			"name": "ok im done",
			"score": 9999,
		},
	)
	
	fmt.Println(jsonData[0])

	if name == "dailyscores" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(jsonData)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	} else if name == "halloffame" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(jsonData)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}
