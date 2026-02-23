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

func Redis_Set(key string , value string , expTime int) error {
	if(Redis_Client == nil) {
		return fmt.Errorf("redis client not initialized");
	}

	err := Redis_Client.Set(CTX , key , value , time.Duration(expTime) * time.Second).Err();
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

func Redis_Publish(key string , msg string) error {
	err := Redis_Client.Publish(CTX , key , msg).Err();
	return err;
}

func Redis_Random(key string) (string, error) {
	data, err := Redis_Client.SRandMember(CTX , key).Result();
	if err != nil {
		return "" , err;
	}

	return data , nil;
}
