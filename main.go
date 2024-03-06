package main

import "github.com/gin-gonic/gin"

type Userstruct struct{

	Name string
	Login string
	Password string
}

var UserSlice =[]Userstruct{}

func main() {

r := gin.Default()
r.POST("/getuser")
r.GET("/list",ListUser)
r.Run(":1212")

}
func ListUser(c *gin.Context)  {
	c.JSON(200,"hello")
}
func Signin(c *gin.Context)  {
	var rrr Userstruct
	c.ShouldBindJSON(&rrr)
	
	if rrr.Login==""||rrr.Name==""{
		c.JSON(404,"errreo")
	}else{
		UserSlice = append(UserSlice, rrr)
		c.JSON(200,"success")
	}

}