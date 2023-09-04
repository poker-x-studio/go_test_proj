package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middleware_1() gin.HandlerFunc {
	fmt.Println("middleware_1(),1")
	return func(c *gin.Context) {
		fmt.Println("middleware_1(),2,", c.ClientIP())
		c.Next()
		//fmt.Println("middleware_1(),3")
	}
}

func middleware_2() gin.HandlerFunc {
	fmt.Println("middleware_2(),1")
	return func(c *gin.Context) {
		fmt.Println("middleware_2(),2,", c.ClientIP())
		c.Next()
		//fmt.Println("middleware_2(),3")
	}
}

func middleware_3() gin.HandlerFunc {
	fmt.Println("middleware_3(),1")
	return func(c *gin.Context) {
		fmt.Println("middleware_3(),2,", c.ClientIP())
		c.Next()
		//fmt.Println("middleware_3(),3")
	}
}
