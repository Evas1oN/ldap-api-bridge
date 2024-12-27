package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/klambri/ldap-api-bridge/api/handlers"
	"github.com/klambri/ldap-api-bridge/internal/configuration"
)

func Run(config *configuration.KlambriConfig) {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.POST("/", handlers.DeleteUser)
		users.DELETE("/", handlers.DeleteUser)
		users.PUT("/", handlers.DeleteUser)
	}

	r.Run(fmt.Sprintf(":%v", config.Server.Port))
}
