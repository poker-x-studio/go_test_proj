package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	fmt.Println("index()")
	c.JSON(200, gin.H{
		"message": "index()",
	})
}

func user_login(c *gin.Context) {
	fmt.Println("user_login()")
	c.JSON(200, gin.H{
		"message": "user_login()",
	})
}

func delete_user(c *gin.Context) {
	fmt.Println("delete_user()")
	c.JSON(200, gin.H{
		"message": "delete_user()",
	})
}

func usersNew(c *gin.Context) {
	fmt.Println("usersNew()")
	c.JSON(200, gin.H{
		"message": "usersNew()",
	})
}

func ups(c *gin.Context) {
	fmt.Println("ups()")
	c.JSON(200, gin.H{
		"message": "ups()",
	})
}

func users(c *gin.Context) {
	fmt.Println("users()")
	c.JSON(200, gin.H{
		"message": "users()",
	})
}

func usersName(c *gin.Context) {
	fmt.Println("usersName()")
	c.JSON(200, gin.H{
		"message": "usersName()",
	})
}

func usersFiles(c *gin.Context) {
	fmt.Println("usersFiles()")
	c.JSON(200, gin.H{
		"message": "usersFiles()",
	})
}
