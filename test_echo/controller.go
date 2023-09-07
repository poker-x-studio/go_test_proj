package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "index~")
}

func user_login(c echo.Context) error {
	fmt.Println("user_login()")
	return c.String(http.StatusOK, "user_login()")
}

func delete_user(c echo.Context) error {
	return c.String(http.StatusOK, "delete_user~")
}

func usersNew(c echo.Context) error {
	return c.String(http.StatusOK, "userNew!")
}

func ups(c echo.Context) error {
	fmt.Println("ups()")
	return c.String(http.StatusOK, "ups()")
}

func users(c echo.Context) error {
	return c.String(http.StatusOK, "user!")
}

func usersName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, fmt.Sprintf("%s, %s", "hi", name))
}

func usersFiles(c echo.Context) error {
	return c.String(http.StatusOK, "usersFiles!")
}
