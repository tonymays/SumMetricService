package server

import (
	"encoding/json"
	"net/http"
)


// ---- jsonErr ----
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// ---- HandleOptionsRequest ----
func HandleOptionsRequest(w http.ResponseWriter, r *http.Request) {
	// establish what the options endpoints can handle
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token, API-Key")
//	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD,GET,DELETE,POST,PATCH")
	w.WriteHeader(http.StatusOK)
}

// ---- SerResponseHeaders ----
func SetResponseHeaders(w http.ResponseWriter) http.ResponseWriter {
	// establish a generic response header
	// a better header place placed here to show greater possibility
	//func SetResponseHeaders( w http.ResponseWriter, authToken string, apiKey string ) http.ResponseWriter {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token, API-Key")
//	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD,GET,DELETE,POST,PATCH")


	// just a for instance, I could establish anything I want to go back as a header
	// and authtoken for instance or an apikey ... I rem it out here for effect
	/*
	if authToken != "" {
		w.Header().Add("Auth-Token", authToken)
	}
	if apiKey != "" {
		w.Header().Add("API-Key", apiKey)
	}
	*/
	return w
}

// ---- throw ----
func throw(w http.ResponseWriter, callErr error) {
	// toss back a standard error message
	w = SetResponseHeaders(w)
	w.WriteHeader(http.StatusForbidden)
	err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusForbidden, Text: callErr.Error()})
	if err != nil {
		panic(err)
	}
}

