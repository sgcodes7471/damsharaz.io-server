package server

import( 
	"net/http"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/db"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{};

func WSServer(w http.ResponseWriter , r *http.Request) {
	roomId := r.URL.Query().Get("roomId");
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


	go func() {
		_, msg, err := conn.ReadMessage();

		if(err != nil) {
			pkg.Log("Error reading in /ws from connection " + name + " : " + err.Error() , "ERROR");
			pkg.Log("GET /ws " + "500" , "WARNING");
			w.WriteHeader(500);
			return
		}

		finalMsgPayload := name + "/r/n" + string(msg);

		if err := db.Redis_Client.Publish(db.CTX , roomId , finalMsgPayload).Err() ; err != nil {
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