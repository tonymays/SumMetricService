package main

import (
	//"errors"
	"pkg/configuration"
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

	// apture core settings
	var err error
	config, err = configuration.CaptureCoreSettings()
	if err != nil {
		// return any errors at this point
		return err
	}
	rcvr.Config = config

	// setup our server

	// return nil if we are good
	return nil
}

// ---- App.Run ----
func (rcvr *App) Run() error {
	// start the server if the train has everything it needs
	return rcvr.Server.Start()
}