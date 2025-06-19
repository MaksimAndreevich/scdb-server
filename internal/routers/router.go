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

	// Инициализируем демо-лимитер
	middleware.InitDemoLimiter()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.WEBUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "X-Demo-Image-ID", "X-Demo-Requests-Used", "X-Demo-Requests-Limit", "X-Demo-Requests-Remaining"},
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
		public.GET("/stats", controllers.GetStats)
		public.GET("/demo/stats", controllers.GetDemoStats)

	}

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(middleware.DemoLimitMiddleware())
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/organizations", controllers.GetOrganizations)
		protected.GET("/organizations/:id", controllers.GetOrganisationById)
	}

	return router
}
