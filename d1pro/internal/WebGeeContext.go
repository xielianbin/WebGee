package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
type WebGeeContext struct {
	// 目标
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	// 响应信息
	StatusCode int
}

func (c *WebGeeContext) PostForm(key string) string {
	//FormValue方法‌：这是Go语言net/http包中Request对象的方法，会自动解析URL查询参数和POST表单数据（包括application/x-www-form-urlencoded和multipart/form-data编码），返回第一个匹配键名的字符串值
	return c.Req.FormValue(key)
}
func (c *WebGeeContext) Query(key string) string {
	//URL.Query()会解析URL中?后的查询字符串（如?name=John&age=20），返回url.Values类型（本质是map[string][]string）89
	//Get(key)方法从映射中提取首个匹配键的值（若键不存在返回空字符串）
	// 处理GET请求 /search?q=golang
	//searchTerm := c.Req.URL.Query().Get("q")
	return c.Req.URL.Query().Get(key)
}
func (c *WebGeeContext) Status(code int) {
	c.StatusCode = code
	//设置响应状态码
	c.Writer.WriteHeader(code)
}
func (c *WebGeeContext) SetHeader(key string, value string) {
	//设置响应头的键值对
	c.Writer.Header().Set(key, value)
}
func (c *WebGeeContext) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	//直接写入字符串响应
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
func (c *WebGeeContext) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	//将变长参数values序列化为JSON输出
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
func (c *WebGeeContext) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
func (c *WebGeeContext) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
