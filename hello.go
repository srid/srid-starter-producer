package main

import (
	"github.com/bmizerany/lpx"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func makeGinAccounts() gin.Accounts {
	return gin.Accounts{
		config.Username: config.Password,
	}
}

func main() {
	router := gin.Default()
	routerAuthed := router.Group("/", gin.BasicAuth(makeGinAccounts()))
	routerAuthed.POST("/logs", logsReceived)
	router.Run(":" + config.Port)
}

func logsReceived(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	println(string(data))
	c.JSON(200, gin.H{})
}
