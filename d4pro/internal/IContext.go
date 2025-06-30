package internal

import "net/http"

// 上下文接口，封装请求和响应
type IContext interface {
	PostForm(key string) string
	Query(key string) string
	Status(code int)
	SetHeader(key string, vlaue string)
	String(code int, format string, values ...interface{})
	JSON(code int, obj interface{})
	Data(code int, data []byte)
	HTML(code int, html string)
	Param(key string) string //获取参数值

	SetWriter(w http.ResponseWriter)
	GetWriter() http.ResponseWriter
	SetRequest(r *http.Request)
	GetRequest() *http.Request
	SetPath(path string)
	GetPath() string
	SetMethod(method string)
	GetMethod() string
	SetStatus(code int)
	GetStatus(code int)
	SetParam(params map[string]string)
}
