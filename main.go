package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglin1995/timeout-control/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.New()

	r.Use(middleware.ContextTimeout(3*time.Second))

	r.GET("/short", Short)
	r.GET("/long", Long)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("%+v", err)
	}
}

func Short(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Long(c *gin.Context) {
	time.Sleep(4*time.Second)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}