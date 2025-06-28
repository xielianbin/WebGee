package internal

type IRouter interface {
	AddRouter(method string, pattern string, handler HandlerFunc)
	Handle(c IContext) //处理函数
	GetHandler() map[string]HandlerFunc
	SetHandler(handler map[string]HandlerFunc)
}
