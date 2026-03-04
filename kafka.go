package main

import(
	"context"
	"fmt"
	"encoding/json"
	"os"
	"github.com/segmentio/kafka-go"
)

type ClickData struct{
	ShortCode string `json:"short_code"`
	IP string `json:"ip"`
	UA string `json:"ua"`

}

var kafkaWriter *kafka.Writer

func initKafka(){
	addr := os.Getenv("KAFKA_URL")
    if addr == "" {
        addr = "localhost:9092"
    }
	kafkaWriter  = &kafka.Writer{
		Addr:kafka.TCP(addr),
		Topic: "click-events",
		Balancer: &kafka.LeastBytes{},
	}
	fmt.Println("Kafka Producer intialized")
}

func ProduceClickEvent(data ClickData){
	payload, _ :=json.Marshal(data)
	
	err:=kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: payload,
		})
	if err != nil{
		fmt.Println("Could not write message into Kafka: ",err)
	}
}

