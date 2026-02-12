package main

import(
	"fmt"
	"github.com/joho/godotenv"
	"sgcodes7471/damsharaz.io-server/internal/server" 
)

func main() {
	err := godotenv.Load(".env");
	
	if(err != nil) {
		fmt.Println("FAILED TO LOAD ENVs");
		return;
	}

	server.WSServer()
	server.HTTPServer()
}