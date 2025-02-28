package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func NewGinRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.LoadHTMLGlob("web/templates/admin.html")

	router.GET("/check-health", func(c *gin.Context) {
		log.Println("router.GET")
		c.JSON(200, gin.H{"message": "rabotaet"})
	})

	return router
}

func (h *Handlers) PromoRoutes(domain *gin.RouterGroup) {
	apiGroup := domain.Group("/api")
	{
		apiGroup.GET("/promocode/:code", h.ApplyPromocode)
		apiGroup.GET("/rewards", h.GetRewards)
		apiGroup.GET("/players", h.GetPlayers)
	}

	admin := domain.Group("/admin")
	{
		admin.GET("/", h.AdminPage)
		admin.POST("/promocode", h.CreatePromocode)
	}
}
