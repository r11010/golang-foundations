/*
Задача 3. Изучите RESTful API построение, маршрутизацию и обработку запросов на Go.
Создайте несколько роутов (маршрутов), каждый из которых обрабатывает определенный HTTP метод (GET, POST)
и возвращает JSON ответ с информацией об эндпоинте
*/
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var database = make(map[string]Student)

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var student Student
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &student)
	if err != nil {
		return
	}
	database[student.Name] = student
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(database)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	http.HandleFunc("/database", handleGetRequest)
	http.ListenAndServe(":8080", nil)
}
