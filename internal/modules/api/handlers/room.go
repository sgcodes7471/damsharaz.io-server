package handlers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/db"
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
		pkg.Api_Error("Error in Decoding the Request Body " + err.Error() , "POST /api/v1/room" , 400 , w);
		return;
	}

	roomId := pkg.CreateRoomId();

	token , err := pkg.CreateToken(reqBody.Name , roomId)

	if(err != nil) {
		pkg.Api_Error("Error in Creating Token : " + err.Error() , "POST /api/v1/room" , 500 , w);
		return;
	}

	err = db.Redis_Set(roomId , token , 3600);

	if(err != nil) {
		pkg.Api_Error("Error in setting the token in redis : " + err.Error() , "POST /api/v1/room" , 500 , w);
		panic(err);
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
		pkg.Api_Error("Error occured in CreateRoom() : " + err.Error() , "POST /api/v1/room" , 500 , w);
		return;
	}

	pkg.Log("POST /api/v1/room " + strconv.Itoa(http.StatusOK) , "INFO");
}
