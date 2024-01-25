package v1

import (
	"net/http"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	var s Server
	http.ListenAndServe(":8080", s)

	http.ListenAndServeTLS(":8080", "", "", s)

	s.Start(":8080")
}
