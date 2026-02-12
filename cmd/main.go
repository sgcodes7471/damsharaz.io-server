package main

import(
	"fmt"
	"net/http"
	"sgcodes7471/damsharaz.io-server/internal/modules"
)

func main() {

	http.HandleFunc(
		"/" , 
		func(w http.ResponseWriter , r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		} ,
	)

	http.HandleFunc(
		"/ping" ,
		api.Ping,
	)

	PORT := ":5000"
	fmt.Println("Server starting on port" + PORT)

	err := http.ListenAndServe(PORT , nil)

	fmt.Println(err)
	if(err != nil) {
		fmt.Println("Some Error occured in starting the server : ")
	} 
}