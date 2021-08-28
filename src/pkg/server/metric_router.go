package server

import (
	"github.com/gorilla/mux"
	"net/http"
	//"pkg"
	"pkg/configuration"
)

// ---- metricRouter ----
type metricRouter struct {
	config		configuration.Configuration
}

// ---- NewMetricRouter ----
func NewMetricRouter(config configuration.Configuration, router *mux.Router) *mux.Router {
	// fill in the metricRouter structure ... note we could have mongo clients,
	// db service clients, all kinds of things can now be passed in here
	metricRouter := metricRouter{config}

	/*
		I always add options because I never know if the client that is making
		the call requires them to be in place like React and Vuejs via Axios
	*/

	// Setup OPTIONS Method for all endpoints
	router.HandleFunc("/metric/{key}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/metric/{key}/sum", HandleOptionsRequest).Methods("OPTIONS")

	// Setup endpoint receivers
	router.HandleFunc("/metric/{key}", metricRouter.postMetric).Methods("POST")
	router.HandleFunc("/metric/{key}/sum", metricRouter.sumMetric).Methods("GET")

	/*
		return our router now that it is setup as a dependency of our server
		to which the server is a dependency of the core app

		This function is housed here to keep the endpoint definitions
		associated with their receivers
	*/
	return router
}

func (rcvr *metricRouter) postMetric(w http.ResponseWriter, r *http.Request) {

}

func (rcvr *metricRouter) sumMetric(w http.ResponseWriter, r *http.Request) {

}
