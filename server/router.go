package server

import (
	"lmarrero/dop-exchange-api/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			ratesGroup.GET("/", rates.GetAllDOPRates)
			ratesGroup.GET("/:target", rates.GetRateFromDOP)
			ratesGroup.GET("custom/:base/:target", rates.GetRateFromBase)
		}
		coins := new(controllers.CoinsController)
		v1.GET("/coins", coins.GetSupportedCoins)
	}

	// swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
