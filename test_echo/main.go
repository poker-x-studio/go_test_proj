package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	user_group := e.Group("/user", middleware_1)
	{
		user_group.Add("GET", "/login", user_login, middleware_2)
		user_group.Add("GET", "/delete", delete_user)
	}
	e.Use(middleware_3)

	// Routes
	e.GET("/", index)
	e.GET("/ups", ups)
	e.GET("/Ups", ups)
	e.GET("/user/new", usersNew)
	e.GET("/user/:name", usersName)
	e.GET("/user/1/files/*", usersFiles)
	e.GET("/user/1/files/*", usersNew)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(data))
	//os.WriteFile("routes.json", data, 0644)

	//
	addr := fmt.Sprintf("%s:%s", "localhost", "8001")
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Start server
	e.Logger.Fatal(e.StartServer(server))
}
