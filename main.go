package main

import (
	"fmt"
	//"github.com/pchmura/twitchChatVotes/Bot"
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
			m := Message{}
			err := conn.ReadJSON(&m)
			if err != nil {
				fmt.Println("Error reading JSON", err)
			}

			fmt.Printf("Message: %#v\n", m)

			switch m.Type {
			case "testMessage":
				fmt.Println("test")
			case "testMessage2":
				fmt.Println("test2")
			default:
				// freebsd, openbsd,
				// plan9, windows...
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

