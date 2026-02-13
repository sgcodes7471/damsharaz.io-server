package pkg

import(
	"fmt"
	"os" 
	"time"
	"runtime/debug"
	"sgcodes7471/damsharaz.io-server/internal/config"
)

func Log(message string , category string) {
	file, err := os.OpenFile(
		config.LOG_FILE_NAME ,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY ,
		0644 ,
	);

	defer file.Close();

	if(err != nil) {
		return;
	}

	var newContent string = category + " [" + time.Now().Format("2006-1-2 15:4:5") + "] : " + message + "\n";

	_, err = file.WriteString(newContent);

	if(err != nil) {
		fmt.Println(newContent)
		return;
	}
}

func Recover_Panic() {
	if r := recover(); r != nil {
		stack := debug.Stack();

		Log(
			fmt.Sprintf(
				"PANIC : %v\nSTACK : %s" , r , stack ,
			) , 
			"ERROR" , 
		);
	}	
}