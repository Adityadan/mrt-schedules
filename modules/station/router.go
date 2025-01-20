package station

import (
	"net/http"

	"github.com/Adityadan/mrt-schedules/common/client/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {

	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(ctx *gin.Context) {
		GetAllStation(ctx, stationService )
	})
}

func GetAllStation(c *gin.Context, service Service) {

	datas, err := service.GetAllStation()
	if  err != nil {
		c.JSON(http.StatusBadRequest,
				response.APIResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Successfully get all stations",
			Data:    datas,
		},
	)

	
}