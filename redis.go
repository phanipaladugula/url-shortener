package main

import(
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx = context.Background()

func initRedis(){
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password:"",
		DB: 0,
	})

	_,err:=redisClient.Ping(ctx).Result()
	if err != nil{
		fmt.Println("Could not connect to redis ", err)
	}else{
		fmt.Println("Connected to Redis Successfully")
	}
}

func SetCachedURL(shortCode string, originalURL string){
	err:=redisClient.Set(ctx,shortCode,originalURL,24*time.Hour).Err()
	if err != nil{
		fmt.Println("Failed to cache URL ",err)
	}
}

func GetCachedURL(shortCode string)(string,error){
	return redisClient.Get(ctx,shortCode).Result()
}