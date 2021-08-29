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
	if rcvr.Config.InitWithTestData == "on" {
		now := time.Now()

		// write a 3 hour record
		mdm1 := db.MetricDataModel{}
		then := now.Add(time.Duration(-180) * time.Minute)
		mdm1.Key = "active_vistors"
		mdm1.Value = 15
		mdm1.EntryTime = then.String()
		rcvr.Db.PostMetric(&mdm1)

		// write a 2 hour record
		mdm2 := db.MetricDataModel{}
		then = now.Add(time.Duration(-120) * time.Minute)
		mdm2.Key = "active_vistors"
		mdm2.Value = 10
		mdm2.EntryTime = then.String()
		rcvr.Db.PostMetric(&mdm2)

		// write a 30 minute record
		mdm3 := db.MetricDataModel{}
		then = now.Add(time.Duration(-30) * time.Minute)
		mdm3.Key = "active_vistors"
		mdm3.Value = 5
		mdm3.EntryTime = then.String()
		rcvr.Db.PostMetric(&mdm3)

		// write a 15 minute record
		mdm4 := db.MetricDataModel{}
		then = now.Add(time.Duration(-15) * time.Minute)
		mdm4.Key = "active_vistors"
		mdm4.Value = 20
		mdm4.EntryTime = then.String()
		rcvr.Db.PostMetric(&mdm4)
	}
}