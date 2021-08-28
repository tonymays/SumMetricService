package server

import (
	"net/http"
)

// establish what the options endpoints can handle
func HandleOptionsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token, API-Key")
//	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD,GET,DELETE,POST,PATCH")
	w.WriteHeader(http.StatusOK)
}

// establish a generic response header
func SetResponseHeaders(w http.ResponseWriter) http.ResponseWriter {
// placed here to show greater possibility
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
