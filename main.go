package useragentrest

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

// SetupRouter - setup router Handlers
func SetupRouter() *gin.Engine {
	route := gin.Default()
	route.POST("/ua", UaParser)
	return route
}

func main() {
	router := SetupRouter()
	router.Run()
}
