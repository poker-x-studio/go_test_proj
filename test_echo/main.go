package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/gommon/log"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Debug = true

	user_group := e.Group("/user", middleware_1)
	{
		user_group.Add("GET", "/login", user_login, middleware_2)
		user_group.Add("GET", "/delete", delete_user)
	}
	e.Use(middleware_3)

	file, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, file)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[ECHO][${time_custom}] ${remote_ip} ${method} ${uri} ${status} ${bytes_out} ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.000",
		Output:           fileAndStdoutWriter,
	}))

	// Routes
	e.GET("/", index)
	e.GET("/ups", ups)
	e.GET("/user/new", usersNew)
	e.GET("/user/:name", usersName)
	e.GET("/user/1/files/*", usersFiles)
	e.GET("/user/1/files/*", usersNew)

	header := `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${long_file}","line":"${line}"}`
	e.Logger.SetHeader(header)
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Warn("warn")
	e.Logger.Error("error")

	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(data))
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
