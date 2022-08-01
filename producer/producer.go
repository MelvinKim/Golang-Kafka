package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 0) //creating a connection to kafka

	if err != nil {
		fmt.Printf("Error when connecting the producer to Kafka: %v", err)
	}

	//set the write deadline
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	//write something
	conn.WriteMessages(kafka.Message{Value: []byte("hi there, i love Kafka")})
}
