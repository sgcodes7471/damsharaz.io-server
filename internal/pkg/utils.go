package pkg

import (
	"fmt"
	"time"
	"os"
	"math/rand"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"));

func CreateToken(name string , roomId string) (string , error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256 , 
		jwt.MapClaims{
			"rommId" : roomId ,
			"admin" : name ,
			"time" : time.Now().Unix(),
		},
	);

	signedToken , err := token.SignedString(secretKey);

	if(err != nil) {
		fmt.Println("TOKEN NOT CREATED IN CreateToken()");
		fmt.Println(err);
		return  "" , err;
	}

	return signedToken , nil
}


func VerifyToken(token string) error {
	validity, err := jwt.Parse(
		token , 
		func(token *jwt.Token) (interface{} , error) {
			return secretKey , nil
		} ,
	);

	if(err != nil) {
		return err;
	}

	if !validity.Valid {
		return fmt.Errorf("invalid token");
	}

	return nil;
}


func CreateRoomId() string {
	CHARSET := "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(){}:;',./<>/*-_+=|ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	LENGTH := 7
	
	rand.Seed(time.Now().UnixNano());

	result := CHARSET[rand.Intn(LENGTH)];
	roomId := string(result);

	return roomId
}