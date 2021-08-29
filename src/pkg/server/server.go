package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"pkg"
	"pkg/configuration"
	"pkg/db"
	"time"
)

// ---- Server struct for the MicroService ----
type Server struct {
	Router			*mux.Router
	Config			configuration.Configuration
	Db				*db.TestDataCache
	MetricService	root.MetricService
}

// ---- NewServer ----
func NewServer(config configuration.Configuration, db *db.TestDataCache, metricService root.MetricService) *Server {
	/*
		This step is also part of the dependency injection paradigm from app.go
		and will iinitialize the router(s), configuration settings and data
		cahce
	*/

	// see the following for a brief understanding of this setting
	// https://www.oreilly.com/library/view/building-restful-web/9781788294287/767f97bf-7aa7-4484-b1cb-112ad5045774.xhtml
	router := mux.NewRouter().StrictSlash(true)

	// add new routers here
	router = NewMetricRouter(config, router, db, metricService)

	// fill in the server struct
	s := Server {
		Router: router,
		Config: config,
		Db: db,
		MetricService: metricService,
	}

	// return the completed server
	return &s
}

// ---- Start ----
func (rcvr *Server) Start() {
	// not using https for the purpose of this test but if I were this is
	// how I would setup the server.Start mechanism
	if rcvr.Config.HTTPS == "on" {
		fmt.Println("Listening on port 8443")
		http.ListenAndServeTLS(":8443", rcvr.Config.Cert, rcvr.Config.Key, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	} else {
		fmt.Println("Listening on port", rcvr.Config.ServerListenPort)
		http.ListenAndServe(rcvr.Config.ServerListenPort, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	}
}

// ---- InitTestData ----
func (rcvr *Server) InitTestData() {
	if rcvr.Config.InitWithTestData() {
		now := time.Now()
		then := now.Add(time.Duration(-180) * time.Minute)
		mdm.Key = 'active_vistors'
		mdm.Value = 15
		mdm.EntryTime = then.String()
		rcvr.db.PostMetric(&mdm)
	}
}