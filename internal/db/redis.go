package db

import (
	"os"
	"context" 
	"fmt" 
	"time"
	"github.com/redis/go-redis/v9"
)

var Redis_Client *redis.Client;

var CTX = context.Background();

func Redis_Init() {
	if(Redis_Client != nil) {
		return
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
	if(Redis_Client == nil) {
		return
	}

	Redis_Client.Close();
}

func Redis_Set(key string , value string) error {
	if(Redis_Client == nil) {
		return fmt.Errorf("redis client not initialized");
	}

	err := Redis_Client.Set(CTX , key , value , 3600*time.Second).Err();
	if(err != nil) {
		return err;
	}

	return nil;
} 

func Redis_Get(key string) (string , error) {
	if(Redis_Client == nil) {
		return "" , fmt.Errorf("redis client not initialized");
	}

	value , err := Redis_Client.Get(CTX , key).Result();
	if(err != nil) {
		return "" , err;
	}

	return value , nil;
}

func Redis_Delete(key string) error {
	deleted , err := Redis_Client.Del(CTX , key).Result();
	if err != nil {
		return err;
	}

	if deleted == 0 {
		return fmt.Errorf("Attempt to Delete non-existing");
	}

	return nil;
}
