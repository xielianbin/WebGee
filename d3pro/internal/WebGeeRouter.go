package internal

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type WebGeeRouter struct {
	roots   map[string]INode //存储每种请求方式的Trie树根节点
	handler map[string]HandlerFunc
}

func (r *WebGeeRouter) Handle(c IContext) {
	//fmt.Println(c)
	n, params := r.GetRouter(c.GetMethod(), c.GetPath())
	//fmt.Println(n, params)
	if n != nil {
		c.SetParam(params)
		fmt.Println(c)
		key := c.GetMethod() + "-" + n.GetPattern()
		fmt.Println(key)
		r.handler[key](c)
		//fmt.Println(r, key, c)
	} else {
		c.String(http.StatusNotFound, "404, NOT FOUND: %s\n", c.GetPath())
	}
}

func (r *WebGeeRouter) GetRouter(method string, pattern string) (INode, map[string]string) {
	searchParts := r.ParsePattern(pattern)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.Search(searchParts, 0)
	//fmt.Println(n)
	if n != nil {
		parts := r.ParsePattern(n.GetPattern())
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		//fmt.Println(n, params)
		return n, params
	}
	return nil, nil
}

func (r *WebGeeRouter) ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}

func NewWebGeeRouter() *WebGeeRouter {
	return &WebGeeRouter{roots: make(map[string]INode), handler: make(map[string]HandlerFunc)}
}
func (r *WebGeeRouter) AddRouter(method, pattern string, handler HandlerFunc) {
	log.Printf("WebGeeRout AddRouter : %4s - %s", method, pattern)
	parts := r.ParsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &WebGeeNode{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handler[key] = handler
}
