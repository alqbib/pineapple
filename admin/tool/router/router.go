package router

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"opensearch-tool/tool/config"
	"opensearch-tool/tool/middle"
)

func Start(cfg *config.Config) {
	if strings.ToLower(cfg.RunMode) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	corsMiddleware := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	})
	router.Use(corsMiddleware)
	router.POST("/login", middle.Login)
	auth := router.Group("/")
	auth.Use(middle.CheckUser())
	{
		auth.GET("/index", middle.Index)
		auth.GET("/refresh", middle.Refresh)
	}

	router.Run(":" + cfg.Port)
}
