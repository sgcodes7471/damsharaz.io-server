package handlers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
)

type PingResponse struct {
	Code int
	Success bool
	Ping string
}

func Ping(w http.ResponseWriter , r *http.Request) {

	var res = PingResponse{
		Code : http.StatusOK,
		Success : true ,
		Ping : "PING" ,
	}

	w.Header().Set("Content-Type" , "application/json")

	err := json.NewEncoder(w).Encode(res)

	if(err != nil) {
		pkg.Log("Error in internals/modules/api.Ping : " + err.Error() , "ERROR");
		return
	}

	pkg.Log("GET /PING " + strconv.Itoa(http.StatusOK) , "INFO");
}