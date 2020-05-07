package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
)

// UaParser - Return parsed UA string in JSON format
func UaParser(c *gin.Context) {
	parser := uaparser.NewFromSaved()
	x, _ := ioutil.ReadAll(c.Request.Body)
	client := parser.Parse(string(x))
	c.JSON(http.StatusOK, client)
}

func main() {
	route := gin.Default()
	route.POST("/ua", UaParser)
	route.Run()
}
