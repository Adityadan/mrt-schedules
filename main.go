package main

import (
	"github.com/Adityadan/mrt-schedules/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitRoute()
}

func InitRoute(){
	var (
		router = gin.Default()
		api = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")
} 