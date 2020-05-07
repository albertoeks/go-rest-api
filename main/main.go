package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Title string `json:"Title"`
	Content string `json:"content"`
}

var Tasks []Task

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Go HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tasks", getAllTasks)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func getAllTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: getAllTasks")
	json.NewEncoder(w).Encode(Tasks)
}

func main() {
	Tasks = []Task{
		Task{Title: "First task", Content: "First task content"},
		Task{Title: "Second task", Content: "Second task content"},
	}

	handleRequests()
}
