package main

import (
	"fmt"
	"net/http"
	"wgpro/internal"
)

func main() {
	wg := internal.NewWebGee()
	wg.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "  welcome!")
	})
	wg.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, " hello welcome!")
	})
	wg.RUN(":9999")
}
