package internal

import (
	"fmt"
	"net/http"
)

type WebGee struct {
	router map[string]HandlerFunc //定义一个路由映射表，将对于路径的请求和处理函数存储起来
}

// 定义构造函数
func NewWebGee() *WebGee {
	// 当声明一个 map、slice 或 channel 类型的变量时，它只是声明了一个变量，并没有分配内存来存储数据。因此，在使用之前，必须初始化它们。
	// 在 NewWebGee 构造函数中，使用 make(map[string]HandlerFunc) 来初始化 router 字段。
	// 这确保了当 WebGee 实例被创建时，router 字段已经是一个可以存储数据的 map。
	return &WebGee{router: make(map[string]HandlerFunc)}
}
func (wg *WebGee) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	wg.router[key] = handler
}
func (wg *WebGee) GET(pattern string, handler HandlerFunc) {
	wg.addRoute("GET", pattern, handler)
}
func (wg *WebGee) POST(pattern string, handler HandlerFunc) {
	wg.addRoute("POST", pattern, handler)
}
func (wg *WebGee) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, wg)
}

// 在开启监听的时候，需要放入地址和一个接口，这个接口必须实现serverHttp方法，当检测到请求的时候就会调用这个方法
func (wg *WebGee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	fmt.Fprintf(w, "监听到：%s\n", key)
	if handler, ok := wg.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
