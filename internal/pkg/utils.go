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


func Parse_Payload(payload string) (string, string, string, error) {
	var author_cut int = -1;
	var event_cut int = -1;

	payload_len := len(payload);

	var index int = 0;
	for index + 3 < payload_len {
		if payload[index : (index + 4)] == "/r/n" {
			if author_cut < 0 {
				author_cut = index;
			} else {
				event_cut = index;
				break;
			}
		}
		index = index + 1;
	}

	if author_cut == -1 || event_cut == -1 {
		return "" , "" , "" , fmt.Errorf("invalid format for the payload");
	}

	author := payload[0 : author_cut];
	event := payload[(author_cut + 4) : event_cut];
	msg := payload[min(payload_len - 4 , event_cut + 4) : (payload_len - 4)];

	return event, author, msg, nil
}