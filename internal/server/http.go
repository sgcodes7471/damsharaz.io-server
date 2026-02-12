package server

import(
	"net/http" 
	"fmt"
	"github.com/go-chi/chi"
	"sgcodes7471/damsharaz.io-server/internal/modules/api/routes"
	// "sgcodes7471/damsharaz.io-server/internal/modules/api/handlers"
)

func HTTPServer() {

	var r *chi.Mux = chi.NewRouter();

	routes.Room_Routes(r);

	// http.HandleFunc(
	// 	"/ping" ,
	// 	handlers.Ping ,
	// ); 

	PORT := ":5000";
	
	fmt.Println("HTTP Server starting on port" + PORT);
	err := http.ListenAndServe(PORT , r);
	
	if(err != nil) {
		fmt.Println("Some Error occured in starting the server : ");
	} 

}