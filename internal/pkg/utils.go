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
			"roomId" : roomId ,
			"admin" : name ,
			"time" : time.Now().Unix(),
		},
	);

	signedToken , err := token.SignedString(secretKey);

	if(err != nil) {
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
	CHARSET := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	LENGTH := 7
	
	rand.Seed(time.Now().UnixNano());

	result := make([]byte , LENGTH);
	
	for i := range result {
		result[i] = CHARSET[rand.Intn(len(CHARSET))];
	}

	roomId := string(result);

	return roomId
}