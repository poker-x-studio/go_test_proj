package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// 中间件
func middleware_1(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("middleware_1(),1")
	return func(c echo.Context) error {
		fmt.Println("middleware_1(),2")
		return next(c)
	}
}

func middleware_2(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("middleware_2(),1")
	return func(c echo.Context) error {
		fmt.Println("middleware_2(),2")
		return next(c)
	}
}

func middleware_3(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("middleware_3(),1")
	return func(c echo.Context) error {
		fmt.Println("middleware_3(),2")
		return next(c)
	}
}
