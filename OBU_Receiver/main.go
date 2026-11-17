package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mohammadReza-kaspour/On-Board-Unit-OBU-microservice/shared"
	"github.com/twmb/franz-go/pkg/kgo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var kafkaClient *kgo.Client

func main() {
	kafkaClient = initKafka()
	defer kafkaClient.Close()

	http.HandleFunc("/ws", wsHandler)
	log.Println("OBU receiver listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade error:", err)
	}
	defer conn.Close()

	log.Println("OBU connected:", conn.RemoteAddr())

	for {
		var data shared.ObuData
		if err := conn.ReadJSON(&data); err != nil {
			log.Printf("OBU %s disconnected: %v\n", conn.RemoteAddr(), err)
		}

		log.Printf("received from %s: %#v", conn.RemoteAddr(), data)

		//produce to kafka topic

		payload, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		record := &kgo.Record{
			Topic: "test-topic",
			Key:   []byte(data.Id),
			Value: payload,
		}

		produceToTopic(kafkaClient, record)
	}
}
