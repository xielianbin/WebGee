package internal

import (
	"fmt"
	"net/http"
)

// HandlerFunc定义了gee使用的请求处理程序
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 引擎实现了ServeHTTP的接口
type Engine struct {
	router map[string]HandlerFunc
}

// New是gee引擎的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET定义了添加GET请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST定义了添加POST请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run定义了启动http服务器的方法
// 在开启监听的时候，需要放入地址和一个接口，这个接口必须实现serverHttp方法，当检测到请求的时候就会调用这个方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
