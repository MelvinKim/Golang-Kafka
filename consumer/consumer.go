package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	//create a kafka connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 0)

	if err != nil {
		fmt.Println("An error occurred while trying to connect the consumer to kafka")
	}

	//set the read deadline
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	batchMessages := conn.ReadBatch(1e3, 1e9)
	bytes := make([]byte, 1e3)
	for {
		_, err := batchMessages.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
	//1e3 --> means 1000
	//1e9 --> means one gigabyte
	//This method enables us to only read one message
	// message, _ := conn.ReadMessage(1e6)
	// fmt.Printf("Printing the message value from Kafka: %v", string(message.Value))
}
