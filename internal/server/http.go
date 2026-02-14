package server

import(
	"os"
	"net/http" 
	"github.com/go-chi/chi"
	"sgcodes7471/damsharaz.io-server/internal/modules/api/routes"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
	"sgcodes7471/damsharaz.io-server/internal/modules/api/middlewares"
)

func HTTPServer() {

	var r *chi.Mux = chi.NewRouter();
	r.Use(middleware.Panic_Handler);
	routes.Room_Routes(r);

	r.Get("/ws" , WSServer);

	// http.HandleFunc(
	// 	"/ping" ,
	// 	handlers.Ping ,
	// ); 

	PORT := os.Getenv("PORT");
	
	pkg.Log("HTTP Server starting on port" + PORT , "INFO");

	err := http.ListenAndServe(PORT , r);
	
	if(err != nil) {
		pkg.Log("Some Error occured in starting the server : " + err.Error() , "ERROR");
	} 

}