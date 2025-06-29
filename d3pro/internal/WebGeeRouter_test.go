package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() IRouter {
	r := NewWebGeeRouter()
	r.AddRouter("GET", "/", nil)
	r.AddRouter("GET", "/hello/:name", nil)
	r.AddRouter("GET", "/hello/b/c", nil)
	r.AddRouter("GET", "/hi/:name", nil)
	r.AddRouter("GET", "/assets/*filepath", nil)
	return r
}
func TestParsePattern(t *testing.T) {
	var wgr = NewWebGeeRouter()
	ok := reflect.DeepEqual(wgr.ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(wgr.ParsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(wgr.ParsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("测试字符串分割失败")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.GetRouter("GET", "/hello/geektutu")

	if n == nil {
		//t.Fatal("值为空，不能返回") //Fatal会立即退出
		t.Error("值为空，不能返回") //Error不会立即推出
	}

	if n.GetPattern() != "/hello/:name" {
		t.Error("应该匹配： /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Error("名字应该等于： 'geektutu'")
	}

	fmt.Printf("匹配路径: %s, params['name']: %s\n", n.GetPattern(), ps["name"])

}
