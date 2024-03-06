package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
r.POST("/ger")
r.Run(":1212")

}