package internal

import "net/http"

// 定义函数类
type HandlerFunc func(http.ResponseWriter, *http.Request)
