/*
Задача 1: Создайте сервер на Go, который будет слушать входящие HTTP запросы.
Сервер должен возвращать простой JSON с сообщением "Hello World!".
*/

package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := &Message{"Hello, World!"}
	data, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
