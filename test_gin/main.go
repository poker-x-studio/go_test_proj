package main

import "github.com/gin-gonic/gin"

func main() {
	// 初始化一个http服务对象
	g := gin.Default()
	gin.SetMode(gin.DebugMode)

	// 通过use设置全局中间件
	// 设置日志中间件，主要用于打印请求日志
	g.Use(gin.Logger())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	g.Use(gin.Recovery())

	g.Use(middleware_3())
	user_group := g.Group("/user", middleware_1())
	{
		user_group.GET("/login", user_login, middleware_2())
		user_group.GET("/delete", delete_user)
	}

	// Routes
	g.GET("/", index)
	g.GET("/ups", ups)
	g.GET("/user/new", usersNew)
	g.GET("/user/:name", usersName)
	g.GET("/user/1/files/", usersFiles)
	//r.GET("/user/1/files/", usersNew)

	g.Run(":8001") // 监听并在 0.0.0.0:8080 上启动服务
}
