package main

import (
	ua "useragentrest"
)

func main() {
	router := ua.SetupRouter()
	router.Run()
}
