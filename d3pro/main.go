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

	wg.POST("/login", func(c internal.IContext) {
		c.JSON(http.StatusOK, H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	wg.GET("/hello/:name", func(c internal.IContext) {
		fmt.Println("hello name")
		fmt.Println(c.Param("name"))
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.GetPath())
	})
	wg.GET("/assets/*filepath", func(c internal.IContext) {
		c.JSON(http.StatusOK, H{"filepath": c.Param("filepath")})
	})
	wg.RUN(":9999")
}
