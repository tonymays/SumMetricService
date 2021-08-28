package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"pkg"
	"pkg/configuration"
	"pkg/db"
)

// ---- metricRouter ----
type metricRouter struct {
	config			configuration.Configuration
	db				*db.TestDataCache
	metricService	root.MetricService
}

// ---- NewMetricRouter ----
func NewMetricRouter(config configuration.Configuration, router *mux.Router, db *db.TestDataCache, metricService root.MetricService) *mux.Router {
	// fill in the metricRouter structure ... note we could have mongo clients,
	// db service clients, all kinds of things can now be passed in here
	metricRouter := metricRouter{config, db, metricService}

	/*
		I always add options because I never know if the client that is making
		the call requires them to be in place like React and Vuejs via Axios
	*/

	// Setup OPTIONS Method for all endpoints
	router.HandleFunc("/metric/{key}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/metric/{key}/sum", HandleOptionsRequest).Methods("OPTIONS")

	// Setup endpoint receivers - all enpoints must check themselves into the middleware for a COVID Test.
	router.HandleFunc("/metric/{key}", VerifyToken(metricRouter.postMetric, config)).Methods("POST")
	router.HandleFunc("/metric/{key}", VerifyToken(metricRouter.getMetric, config)).Methods("GET")
	router.HandleFunc("/metric/{key}/sum", VerifyToken(metricRouter.sumMetric, config)).Methods("GET")

	/*
		return our router now that it is setup as a dependency of our server
		to which the server is a dependency of the core app

		This function is housed here to keep the endpoint definitions
		associated with their receivers
	*/
	return router
}

func (rcvr *metricRouter) postMetric(w http.ResponseWriter, r *http.Request) {
	// grab the body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	err = r.Body.Close()
	if err != nil {
		panic(err)
	}

	var m root.Metric
	err = json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}

	// grab the key
	vars := mux.Vars(r)
	m.Key = vars["key"]

	// show it to me
	fmt.Println(m)




	w = SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)


/*	vars := mux.Vars(r)




	var user root.User
	user.Uuid = vars["userUuid"]
	users, err := rcvr.dbService.FindUser(user)
	if err != nil {
 		throw(w,err)
		return
	}
	var role root.Role
	role.Uuid = vars["roleUuid"]
	roles, err := rcvr.dbService.FindRole(role)
	if err != nil {
		throw(w,err)
		return
	}
	var filter root.RoleUser
	var update root.RoleUser
	filter.UserUuid = users[0].PrivateUuid
	filter.RoleUuid = roles[0].PrivateUuid
	update.Status = "Active"
	err = rcvr.dbService.UpdateRoleUser(filter,update)
	if err == nil {
		w = SetResponseHeaders(w, "", "")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(update)
		if err != nil {
			panic(err)
		}
	} else {
		throw(w,err)
	}
*/
}

func (rcvr *metricRouter) getMetric(w http.ResponseWriter, r *http.Request) {
	w = SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
}

func (rcvr *metricRouter) sumMetric(w http.ResponseWriter, r *http.Request) {
	w = SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
}
