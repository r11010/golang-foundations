/*
Задача 2: Дополните сервер так, чтобы он обрабатывал POST запросы и принимал в теле запроса объект в формате JSON.
Сервер должен возвращать ответ с этим же объектом.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var msg Message

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error occurred: %s", err)
	}

	err = json.Unmarshal(bytes, &msg)
	if err != nil {
		fmt.Printf("Error occurred: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
	fmt.Printf("%s", err)
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	http.ListenAndServe(":8080", nil)
}
