package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func NewGinRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/check-health", func(c *gin.Context) {
		log.Println("router.GET")
		c.JSON(200, gin.H{"message": "rabotaet"})
	})

	return router
}

func (handler *Handlers) PromoRoutes(domain *gin.RouterGroup) {
	domain.POST("/api/promocode", handler.CreatePromocode)
	domain.GET("/api/promocode/:code", handler.ApplyPromocode)
}
