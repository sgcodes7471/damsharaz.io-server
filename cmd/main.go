package main

import(
	"github.com/joho/godotenv"
	"sgcodes7471/damsharaz.io-server/internal/server" 
	"sgcodes7471/damsharaz.io-server/internal/pkg" 
	"sgcodes7471/damsharaz.io-server/internal/db" 
)

func main() {
	err := godotenv.Load(".env");
	
	if(err != nil) {
		pkg.Log("FAILED TO LOAD ENVs" , "ERROR");
		return;
	}

	db.Redis_Init();

	server.WSServer()
	server.HTTPServer()
}