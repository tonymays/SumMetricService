package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	//"pkg"
	"pkg/configuration"
)

// ---- Server struct for the MicroService ----
type Server struct {
	Router		*mux.Router
	Config		configuration.Configuration
}

// ---- NewServer ----
func NewServer(config configuration.Configuration) *Server {
	/*
		This step is also part of the dependency injection paradigm from app.go
		and will iinitialize the router(s) amd configuration settings
	*/

	// see the following for a brief understanding of this setting
	// https://www.oreilly.com/library/view/building-restful-web/9781788294287/767f97bf-7aa7-4484-b1cb-112ad5045774.xhtml
	router := mux.NewRouter().StrictSlash(true)

	// add new routers here
	router = NewMetricRouter(config, router)

	// fill in the server struct
	s := Server {
		Router: router,
		Config: config,
	}

	// return the completed server
	return &s
}

// ---- Start ----
func (rcvr *Server) Start() error {
	if rcvr.Config.HTTPS {
		return http.ListenAndServeTLS(":8443", rcvr.Config.Cert, rcvr.Config.Key, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	} else {
		return http.ListenAndServe(rcvr.Config.ServerListenPort, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	}
}