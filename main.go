package main

import "github.com/gin-gonic/gin"

func main() {
r := gin.Default()
r.POST("/getuser")
r.GET("/list",ListUser)
r.Run(":1212")

}
func ListUser(c *gin.Context)  {
	c.JSON(200,"hello")
}