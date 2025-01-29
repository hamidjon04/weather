package api

import (
	"log/slog"
	"weather/api/handler"
	"weather/service"

	_ "weather/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Weather
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Router(service service.Service, log *slog.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.NewHandler(service, log)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	return router
}
