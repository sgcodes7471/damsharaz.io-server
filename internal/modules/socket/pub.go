package socket

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/db"
	"sgcodes7471/damsharaz.io-server/internal/types"
)

/**
	START event publish

	1. check if has a valid admin token - done
	2. choose a person from the redis members randomly - done
	3. update that connection in the redis room model in Den - done
 	4. publish to redis - done
	

	ANSWER event publish

	1. anyone can send with a format of <author>/r/nANSWERr/n<movie>
	2. simply publish to redis


	WINNER event publish

	1. only Den can send - the name of the Winner in msg
	2. simply publish-sub


	LEAVE or DISCONNECT event publish

	1. clear Conn from the redis
	2. publish LEAVE/DISCONNECTED event
*/


func Publish(payload string , conn *websocket.Conn , r *http.Request) error {
	event , author, msg , err := pkg.Parse_Payload(payload);

	if err != nil {
		return err
	}

	switch event {
	case "START" :
		token, err := r.Cookie("token");
		if err != nil {
			return fmt.Errorf("Unauthorized access");
		}

		if token["roomId"] != r.Header.Get("roomId") {
			return fmt.Errorf("Unauthorized access");
		}

		var den_client_string string;
		den_client_string , err = db.Redis_Random(roomId) 
		if err != nil {
			return err;
		} 

		var den_client types.Client_Object;
		err = json.Unmarshal(den_client_string , &den_client);

		if err != nil {
			return err;
		}

		roomObject := type.Room_Object{
			RoomId : r.Header.Get("roomId") ,
			Token : token ,
			Den : den_client ,
			Ongoing : true
		}

		var data string;
		data , err = json.Marshal(roomObject);

		if err != nil {
			return err;
		}

		if err := db.Redis_Set(roomId, data) ; err != nil {
			return err;
		}

		if err := db.Redis_Publish(roomId , payload) ; err != nil {
			return err;
		}

	default :
		return Errorf("Not a Valid Event")
	}

	return nil;
}


