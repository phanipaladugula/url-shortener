package main

import(
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func StartAnalyticsWorker(){
	reader:=kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "click-events",
		GroupID: "analytics-group",
		MaxWait: 5 * time.Second,
	})

	fmt.Println("Worker is waiting for click events")

	for {
		m,err:=reader.ReadMessage(context.Background())
		if err != nil{
			fmt.Println("Error reading message: ",err)
			time.Sleep(5*time.Second)
			continue
		}

		var data ClickData
		json.Unmarshal(m.Value,& data)

		_,err=dbPool.Exec(context.Background(),
		"INSERT INTO analytics (short_code, ip_address,user_agent) VALUES ($1, $2, $3)",
		data.ShortCode,data.IP,data.UA)

		if err != nil{
			fmt.Println("Worker failed to save to DB: ",err)
		}else{
			fmt.Println("Worker successfully logged click for ",data.ShortCode)
		}
	}
}