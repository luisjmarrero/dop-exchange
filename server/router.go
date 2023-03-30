package server

import (
	"lmarrero/dop-exchange-api/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		ratesGroup := v1.Group("rates")
		{
			rates := new(controllers.RatesController)
			ratesGroup.GET("/", rates.GetAllRates)
			ratesGroup.GET("/:coin", rates.GetRate)
		}
	}
	return router

}
