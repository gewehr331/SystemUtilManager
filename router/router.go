package router

import (
	"SystemUtilManager/handler"
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter()

func init() {

	Router.HandleFunc("/agent/{id:[0-9]+}", handler.Agents)

	Router.HandleFunc("/get_res", handler.GetResOfScan)
	Router.HandleFunc("/admin", handler.AdminPanel)
	Router.HandleFunc("/synchronization", handler.Synchronization)
	http.Handle("/", Router)
}
