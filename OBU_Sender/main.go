package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mohammadReza-kaspour/On-Board-Unit-OBU-microservice/shared"
)

const SendInterval = time.Second * 2
const wsEndpoint = "ws://localhost:8080/ws"

func main() {
	for {
		if err := run(); err != nil {
			log.Println("connection lost:", err, "- reconnecting in 3s")
		}
		time.Sleep(time.Second * 3)
	}
}
func run() error {
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	ticker := time.NewTicker(SendInterval)
	defer ticker.Stop()

	for range ticker.C {
		generatedObuData := shared.NewObuData()
		produceToWs(conn, generatedObuData)
	}

	return nil
}

func produceToWs(conn *websocket.Conn, data *shared.ObuData) error {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		return err
	}

	if err := conn.WriteJSON(data); err != nil {
		fmt.Println("write error:", err)
		return err
	}

	fmt.Printf("sent: %#v\n", data)
	return nil
}
