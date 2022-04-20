package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	log.Println("Server listen on port", port)
	log.Fatalln(http.ListenAndServe(port, muxDispatcher))
}
