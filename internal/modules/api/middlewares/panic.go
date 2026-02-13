package middleware

import(
	"net/http" 
	"sgcodes7471/damsharaz.io-server/internal/pkg"
)

func Panic_Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer pkg.Recover_Panic()
		next.ServeHTTP(w, r)
	});
}