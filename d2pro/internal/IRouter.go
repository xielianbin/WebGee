package internal

type IRouter interface {
	AddRouter(method string, pattern string, handler HandlerFunc)       //注册路由
	GetRouter(method string, pattern string) (INode, map[string]string) //路由匹配
	ParsePattern(pattern string) []string
	Handle(c IContext)
}
