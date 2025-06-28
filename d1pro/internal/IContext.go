package internal

// 上下文接口，封装请求和响应
type IContext interface {
	PostFrom(key string)
	Query(key string)
	Status(code int)
	SetHeader(key string, vlaue string)
	String(code int, format string, values ...interface{})
	JSON(code int, obj interface{})
	Data(code int, data []byte)
	HTML(code int, html string)
}
