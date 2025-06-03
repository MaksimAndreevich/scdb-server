package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/scdb/server/internal/controllers"
	"gitlab.com/scdb/server/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Recovery восстанвление работы приложения после ошибок
	router.Use(gin.Recovery())

	// Публичные маршруты
	public := router.Group("/api")

	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/organisations", controllers.GetOrganisations)
		protected.GET("/organisations/:id", controllers.GetOrganisationById)
	}

	return router
}
