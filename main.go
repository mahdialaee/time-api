package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type uTime struct {
	HumanReadable string
	UnixMilli     int64
}

func getTime(c *gin.Context) {

	var tnow = uTime{}

	tnow.HumanReadable = time.Now().Format(time.ANSIC)
	tnow.UnixMilli = time.Now().UnixMilli()

	c.IndentedJSON(http.StatusOK, tnow)

}

func main() {

	router := gin.Default()
	router.GET("/time", getTime)
	router.Run("localhost:8080")
}
