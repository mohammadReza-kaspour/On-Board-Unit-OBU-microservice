package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mohammadReza-kaspour/On-Board-Unit-OBU-microservice/shared"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws" , wsHandler)
	log.Println("OBU receiver listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandler(w http.ResponseWriter , r *http.Request){
	conn , err := upgrader.Upgrade(w , r , nil)
	if err != nil {
		log.Fatal("upgrade error:" , err)
	}
	defer conn.Close()

	log.Println("OBU connected:" , conn.RemoteAddr())

	for {
		var data shared.ObuData
		if err := conn.ReadJSON(&data); err != nil {
			log.Printf("OBU %s disconnected: %v\n", conn.RemoteAddr(), err)
		}

		log.Printf("received from %s: %#v", conn.RemoteAddr(), data)
	}
}