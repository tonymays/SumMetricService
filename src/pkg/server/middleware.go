package server

import (
	"net/http"
	"pkg/configuration"
)

/*
	this go file serves as a router middleware.  the purpose of this middleware
	will be to accept a JWT Token from the header, for instance, as then
	verify it.

	I put this here to demonstrate YOUR need for it, but, it is unnecessary
	for this test microservice since no longer term data storage is
	occurring.

	Every router endpoint must check itself against this middleware or you are
	most defiinitely at risk.
*/

// ---- VerifyToken ----
func VerifyToken(next http.HandlerFunc, config configuration.Configuration) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}