package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the api"))
}

func HandleLog(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var log Log
	err := decoder.Decode(&log)

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	response, err := log.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := DB.Create(&Log{UserId: log.UserId, MetaData: log.MetaData, Action: log.Action})

	if result.Error != nil {
		panic(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
