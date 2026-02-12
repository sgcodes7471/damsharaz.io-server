package handlers

import(
	"fmt"
	"encoding/json"
	"net/http"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
)

type CreateRoomResponse struct {
	Code int
	Success bool
	RoomId string
	Token string
}

type CreateRoomRequest struct {
	Name string
}

func CreateRoom(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type" , "application/json");

	var reqBody CreateRoomRequest;
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		fmt.Println("Error in Decoding the Request Body");
		fmt.Println(err);
		w.WriteHeader(400);
		return;
	}

	roomId := pkg.CreateRoomId();

	token , err := pkg.CreateToken(reqBody.Name , roomId)

	if(err != nil) {
		w.WriteHeader(500);
		return;
	}

	var res = CreateRoomResponse{
		Code : http.StatusOK ,
		Success : true ,
		RoomId : roomId ,
		Token : token ,
	}

	tokenCookie := http.Cookie{
        Name:     "token",
        Value:    token,
        Path:     "/",
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }

	http.SetCookie(w , &tokenCookie);

	err = json.NewEncoder(w).Encode(res);

	if(err != nil) {
		fmt.Println("Some Error occured in CreateRoom()");
		fmt.Println(err);
		w.WriteHeader(500);
		return;
	}
}
