package db

import (
	"os"
	"context" 
	"fmt" 
	"time"
	"github.com/redis/go-redis/v9"
)

var Redis_Client *redis.Client;

var ctx = context.Background();

func Redis_Init() {
	if(Redis_Client != nil) {
		return;
	}

	Redis_Client = redis.NewClient(
		&redis.Options {
			Addr : os.Getenv("REDIS_ADDR") ,
			Password : os.Getenv("REDIS_PASSWORD") ,
			DB : 0 ,
		} ,
	);
}

func Redis_Close() {
	if(Redis_Client != nil) {
		return;
	}

	Redis_Client.Close();
}

func Redis_Set(key string , value string) error {
	if(Redis_Client == nil) {
		return fmt.Errorf("redis client not initialized");
	}

	err := Redis_Client.Set(ctx , key , value , 3600*time.Second).Err();
	if(err != nil) {
		panic(err);
		return nil;
	}

	return nil;
} 

func Redis_Get(key string) (string , error) {
	if(Redis_Client == nil) {
		return "" , fmt.Errorf("redis client not initialized");
	}

	value , err := Redis_Client.Get(ctx , key).Result();
	if(err != nil) {
		panic(err);
		return "" , err;
	}

	return value , nil;
}