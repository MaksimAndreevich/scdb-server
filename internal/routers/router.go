package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/scdb/server/internal/config"
	"gitlab.com/scdb/server/internal/controllers"
	"gitlab.com/scdb/server/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.WEBLocalUrl, config.AppConfig.WEBProdUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Recovery восстанвление работы приложения после ошибок
	router.Use(gin.Recovery())

	// Публичные маршруты
	public := router.Group("/api")

	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
		public.GET("/organizations/total", controllers.GetTotalOrganizations)

	}

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/organizations", controllers.GetOrganizations)
		protected.GET("/organizations/:id", controllers.GetOrganisationById)

	}

	return router
}
