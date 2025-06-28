package internal

import (
	"net/http"
)

type IEngin interface {
	RUN(addr string)
	GET(pattern string, handler HandlerFunc)
	POST(pattern string, handler HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
