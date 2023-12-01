package routes

import (
	"github.com/Inf85/go-project-demo/api"
	"github.com/Inf85/go-project-demo/auth"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func NewRoutes(api *api.API) *mux.Router {
	mux := mux.NewRouter()

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignUp).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	return mux
}
