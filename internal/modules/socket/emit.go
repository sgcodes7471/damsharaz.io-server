package socket

import( 
	"net/http"
	"encoding/json"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/db"
	"sgcodes7471/damsharaz.io-server/internal/types"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

/**
	START event emit

	1. pull the room objet from redis
	2. check for den conn in the locally stored connections
	3. if found then choose a random movie name from the list 
	4. send Den that movie name in the format
 	5. to others send the name of Den


	LEAVE or DISCONNECT 

	1. clear from the Room_conections in memory if conn exists
	2. else publish event

*/