package internal

import (
	"fmt"
	"log"
	"net/http"
)

type WebGeeRouter struct {
	handler map[string]HandlerFunc
}

func (r *WebGeeRouter) GetHandler() map[string]HandlerFunc {
	return r.handler
}

func (r *WebGeeRouter) SetHandler(handler map[string]HandlerFunc) {
	r.handler = handler
}

func NewWebGeeRouter() *WebGeeRouter {
	return &WebGeeRouter{make(map[string]HandlerFunc)}
}
func (r *WebGeeRouter) AddRouter(method, pattern string, handler HandlerFunc) {
	log.Printf("WebGeeRout %4s - $s", method, pattern)
	key := method + "-" + pattern
	r.handler[key] = handler
}
func (r *WebGeeRouter) Handle(c IContext) {
	key := c.GetMethod() + "-" + c.GetPath()
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		fmt.Println(key)
		c.String(http.StatusNotFound, "404 page not found: %s\n", c.GetPath())
	}
}
