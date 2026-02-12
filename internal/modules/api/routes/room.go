package routes

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"sgcodes7471/damsharaz.io-server/internal/modules/api/handlers"
)

func Room_Routes(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes);

	r.Route(
		"/api/v1/room" ,
		func (router chi.Router) {
			router.Post("/" , handlers.CreateRoom)
		} ,
	);
}