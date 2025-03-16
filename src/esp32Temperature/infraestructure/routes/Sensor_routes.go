package routes

import (
	"act2/src/esp32Temperature/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/sensors")

	postSensorController := dependencies.GetPostSensorController()

	routes.POST("/", postSensorController.Execute) 

}
