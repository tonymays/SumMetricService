package main

import (
	//"fmt"
	"pkg/configuration"
	"pkg/db"
	"pkg/server"
)

// ---- App Struct for the MicroService ----
type App struct {
	Config configuration.Configuration
	Server *server.Server
}

// ---- App.Init ----
func (rcvr *App) Init() error {
	/*
		-----------------------------------------------------------------------
			App.Init is perform as a dependency injection for the app core
			struct of this MicroService.  This API cannot run without all
			dependencies being satisfed.

			For more informatuon:
			https://blog.drewolson.org/dependency-injection-in-go
		-----------------------------------------------------------------------
	*/

	// Step 1: capture core settings
	var config configuration.Configuration
	config, err := configuration.CaptureCoreSettings()
	if err != nil {
		return err
	}
	rcvr.Config = config

	// Step 2: setup the data cache
	dataCache := db.NewTestDataCache([]*db.MetricDataModel{})

	// Step 3: Add router services here
	metricService := db.NewMetricService(config, dataCache)

	// Step 3: setup our server
	rcvr.Server = server.NewServer(config, dataCache, metricService)

	// Step 4: initialize test data
	rcvr.Server.InitTestData()

	// return nil if we are good
	return nil
}

// ---- App.Run ----
func (rcvr *App) Run() {
	// start the server if the train has everything it needs
	rcvr.Server.Start()
}