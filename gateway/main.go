package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Simonwtaylor/code-agile/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MessageType struct {
	Example string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func reader(conn *websocket.Conn) {
	for {
		var message MessageType
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(messageType)
		json.Unmarshal(p, &message)

		log.Println(message.Example)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected 🔌")

	reader(ws)
}

func getTasks(w http.ResponseWriter, r *http.Request) []models.TaskModel {
	tasks := []models.TaskModel{
		models.TaskModel{
			ID:    "123",
			Title: "Task 1",
		},
	}

	return tasks
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World 🔥")
	r := mux.NewRouter()

	r.HandleFunc("/", homePage)
	r.HandleFunc("/ws", wsEndpoint)
	r.HandleFunc("/tasks", getTasks).Methods("GET")

	// setupRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
