package main

import (
	"fmt"
	"net/http"
	"wgpro/internal"
)

func main() {
	fmt.Println("	启动WebGee")

	type H map[string]interface{} //给json数据取别名
	wg := internal.NewWebGeeEngine(internal.NewWebGeeRouter())

	wg.GET("/", func(c internal.IContext) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	wg.POST("/hello", func(c internal.IContext) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.GetPath())
	})

	wg.GET("/index", func(c internal.IContext) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := wg.Group("/v1")
	{
		v1.GET("/", func(c internal.IContext) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c internal.IContext) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.GetPath())
		})
	}
	v2 := wg.Group("/v2")
	{
		v2.GET("/hello/:name", func(c internal.IContext) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.GetPath())
		})
		v2.POST("/login", func(c internal.IContext) {
			c.JSON(http.StatusOK, H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	wg.RUN(":9999")
}
