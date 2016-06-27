package main

import (
	"github.com/gin-gonic/gin"

	. "github.com/apdaza/oasRuler/utils"
	. "github.com/apdaza/oasRuler/controllers"
)




func main() {
	r := gin.Default()

	r.Use(Cors())

	rule := r.Group("api/rules")
	{
		rule.GET("/domains", GetDomains)
		rule.GET("/domains/:id", GetDomain)
		rule.POST("/domains", PostDomain)
		rule.PUT("/domains/:id", UpdateDomain)
		rule.DELETE("/domains/:id", DeleteDomain)

		rule.GET("/rules", GetRules)
		rule.GET("/rules/:id", GetRule)
		rule.POST("/rules", PostRule)
		rule.PUT("/rules/:id", UpdateRule)
		rule.DELETE("/rules/:id", DeleteRule)

		rule.GET("/components", GetComponents)
		rule.GET("/components/:id", GetComponent)
		rule.GET("/componentsbyrule/:name", GetComponentByRule)
		rule.POST("/components", PostComponent)
		rule.PUT("/components/:id", UpdateComponent)
		rule.DELETE("/components/:id", DeleteComponent)

		rule.OPTIONS("/domains", Options)        // POST
		rule.OPTIONS("/domains/:id", Options)    // PUT, DELETE
		rule.OPTIONS("/rules", Options)          // POST
		rule.OPTIONS("/rules/:id", Options)      // PUT, DELETE
		rule.OPTIONS("/components", Options)     // POST
		rule.OPTIONS("/components/:id", Options) // PUT, DELETE
	}
	r.Run(":8080")
}
