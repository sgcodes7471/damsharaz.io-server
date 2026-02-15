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
		pkg.Log("Error in Decoding the Request Body " + err.Error() , "ERROR")
		pkg.Log("POST /api/v1/room " + "400" , "INFO");
		w.WriteHeader(400);
		return;
	}

	roomId := pkg.CreateRoomId();

	token , err := pkg.CreateToken(reqBody.Name , roomId)

	if(err != nil) {
		pkg.Log("Error in Creating Token : " + err.Error() , "ERROR");
		pkg.Log("POST /api/v1/room " + "500" , "WARNING");
		w.WriteHeader(500);
		return;
	}

	err = db.Redis_Set(roomId , token , 3600);

	if(err != nil) {
		pkg.Log("POST /api/v1/room " + "500" , "WARNING");
		w.WriteHeader(500);
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
		pkg.Log("Error occured in CreateRoom() : " + err.Error() , "ERROR");
		pkg.Log("POST /api/v1/room " + "500" , "WARNING");
		w.WriteHeader(500);
		return;
	}

	pkg.Log("POST /api/v1/room " + strconv.Itoa(http.StatusOK) , "INFO");
}
