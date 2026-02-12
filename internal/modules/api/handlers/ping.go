package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type PingResponse struct {
	Code int
	Success bool
	Ping string
}

func Ping(w http.ResponseWriter , r *http.Request) {
	// fmt.Fprintf(w , "PING")

	var res = PingResponse{
		Code : http.StatusOK,
		Success : true ,
		Ping : "PING" ,
	}

	w.Header().Set("Content-Type" , "application/json")

	err := json.NewEncoder(w).Encode(res)

	if(err != nil) {
		fmt.Println("Some error in internals/modules/api.Ping")
		fmt.Println(err)
		return
	}
}