package main

import (
	"encoding/json"
	"fmt"
	"github.com/twitchChatVotes/Bot"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Data		string `json:"data"`
	Type		string `json:"type"`
}

type FormData struct {
	Channel		string `json:"channel"`
	Option1		string `json:"option1"`
	Option2		string `json:"option2"`
	Emote1		string `json:"emote1"`
	Emote2		string `json:"emote2"`
	Duration	int `json:"duration"`
	Votes1		int `json:"votes1"`
	Votes2  	int `json:"votes2"`
}

type Messages struct {
	Control string `json:"control"`
	X json.RawMessage
}

func main() {

	 //go bot.RunBot("#shroud", "shroud200", "shroudTHICC")
	 //bot.RunBot("#ninja", "ninjaCRINJA", "ninjaS")
	fs := http.FileServer(http.Dir("/twitch-chat-votes/public/"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleBasicWS)

	go handleMessages()

	log.Println("http server started on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Listenandserve error: ", err)
	}
}

func handleBasicWS(w http.ResponseWriter, r *http.Request){
	var conn, _ = upgrader.Upgrade(w,r,nil)
	go func(conn *websocket.Conn) {

		for {
			var m Messages
			err := conn.ReadJSON(&m)
			if err != nil {
				fmt.Println("Error durning reading:", err)
				conn.Close()
				break

			}

			fmt.Printf("Message: %#v\n", m)
			switch m.Control {
			case "formData":
				fmt.Println("formData")
				var formData FormData
				if err := json.Unmarshal([]byte(m.X), &formData); err != nil {
					// handle error
					fmt.Println("error durning unmarshaling: ", err)
				}
				// do something
				fmt.Printf("%+v\n", formData)
				go bot.RunBot(formData.Channel, formData.Emote1, formData.Emote2, formData.Duration, conn)
			case "testMessage":
				fmt.Println("test")
			case "testMessage2":
				fmt.Println("test2")
			default:
				fmt.Println("other")
			}
		}
		msg := Message{Data:"pls work", Type:"test"}
		conn.WriteJSON(msg)
	}(conn)
}

func handleConnections(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true

	for{
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages(){


	for{
		msg := <- broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

