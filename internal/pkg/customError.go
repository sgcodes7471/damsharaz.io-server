package pkg

import(
	"net/http"
	"strings"
)

func Api_Error(error_text string , api_route string , res_status int , w http.ResponseWriter) {
	if strings.TrimSpace(error_text) != "" {
		Log(error_text , "ERROR");
	}
	
	var issue_level string;
	if res_status < 500 {
		issue_level = "INFO";
	} else {
		issue_level = "WARNING";
	}

	api_log := strings.TrimSpace(api_route) + " " + string(res_status);
	
	Log(api_log , issue_level);

	w.WriteHeader(res_status);
}