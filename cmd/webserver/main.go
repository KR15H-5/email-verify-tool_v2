package main

import (
	"email-verify-tool_v2/domainchecker" // Adjust with the correct import path

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoint for domain verification
	router.POST("/check", func(c *gin.Context) {
		domain := c.PostForm("domain")
		result := domainchecker.CheckDomain(domain) // Assuming this function exists and returns the verification results
		c.JSON(200, result)
	})

	router.Run(":8080")
}
