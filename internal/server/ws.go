package server

import( 
	"net/http"
	"encoding/json"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/db"
	"sgcodes7471/damsharaz.io-server/internal/types"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var Rooms_Connections = make(map[string][]types.Client_Object);

var upgrader = websocket.Upgrader{};

func WSServer(w http.ResponseWriter , r *http.Request) {
	roomId := r.Header.Get("roomId");
	name := r.URL.Query().Get("name");

	if roomId == "" || name == "" {
		pkg.Log("Required Query Params missing in /ws" , "ERROR");
		pkg.Log("GET /ws " + "400" , "INFO");
		w.WriteHeader(400);
		return
	}

	for _, _ch := range name {
		if _ch == '/' {
			pkg.Log("Name has '/' : FORBIDDEN" , "ERROR");
			pkg.Log("GET /ws " + "403" , "INFO");
			w.WriteHeader(403);
			return
		}
	}

	var value string;
	value , err := db.Redis_Get(roomId);

	if  err == redis.Nil {
		pkg.Log("RoomId invalid! Room NOT FOUND in Redis" , "ERROR");
		pkg.Log("GET /ws " + "400" , "INFO");
		w.WriteHeader(404);
		return
	} else if err != nil {
		pkg.Log("Error in connecting to Redis : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}

	
	var Room = types.Room_Object{
		RoomId : roomId ,
		Token : value ,
		Den : nil ,
		Ongoing : false ,
	};

	data , err := json.Marshal(Room);
	
	if err != nil {
		pkg.Log("Error in Serializing Room Object : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}

	key := roomId + "_data";
	if _ , err := db.Redis_Get(key); err == redis.Nil {
		err = db.Redis_Set(key , string(data) , 0);
		
		if err != nil {
			pkg.Log("Error in writing Room Object to Redis : " + err.Error() , "ERROR");
			pkg.Log("GET /ws " + "500" , "WARNING");
			w.WriteHeader(500);
			return
		}
	} else if err != nil {
		pkg.Log("Error in connecting to Redis : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}

	
	conn, err := upgrader.Upgrade(w, r, nil);

	if err != nil {
		pkg.Log("Websocket did not start : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}
	defer conn.Close();


	sub := db.Redis_Client.Subscribe(db.CTX , roomId);
	defer sub.Close();

	ch := sub.Channel();

	client := types.Client_Object{
		Conn : conn ,
		Name : name ,
	}

	Rooms_Connections[roomId] = append(Rooms_Connections[roomId] , client);

	data , err = json.Marshal(client);

	if err != nil {
		pkg.Log("Error in Serializing Client Object : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}

	key = roomId + "_member";
	if err := db.Redis_Client.SAdd(db.CTX , key , data).Err(); err != nil {
		pkg.Log("Error in adding member to Redis : " + err.Error() , "ERROR");
		pkg.Log("GET /ws " + "500" , "WARNING");
		w.WriteHeader(500);
		return
	}


	go func() {
		_, msg, err := conn.ReadMessage();

		if(err != nil) {
			pkg.Log("Error reading in /ws from connection " + name + " : " + err.Error() , "ERROR");
			pkg.Log("GET /ws " + "500" , "WARNING");
			w.WriteHeader(500);
			return
		}

		finalMsgPayload := name + "/r/n" + string(msg);

		if err := db.Redis_Publish(roomId , finalMsgPayload) ; err != nil {
			pkg.Log("Error publishing to Redis in /ws from connection " + name + " : " + err.Error() , "ERROR");
			pkg.Log("GET /ws " + "500" , "WARNING");
			w.WriteHeader(500);
			return
		}
	} ();


	for msg := range ch {
		if err := conn.WriteMessage(websocket.TextMessage , []byte(msg.Payload)); err != nil {
			pkg.Log("Error writing from Redis in /ws to connection " + name + " : " + err.Error() , "ERROR");
			pkg.Log("GET /ws " + "500" , "WARNING");
			w.WriteHeader(500);
			return
		}
	}
}