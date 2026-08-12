package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// create log file
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// write logs to both stdout and file
	mw := io.MultiWriter(os.Stdout, f)
	logger := log.New(mw, "", log.LstdFlags)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		logger.Println("INFO: normal request received")
		c.JSON(http.StatusOK, gin.H{
			"message": "info",
		})
	})

	r.GET("/warn", func(c *gin.Context) {
		logger.Println("WARNING: suspicious activity detected")
		c.JSON(http.StatusOK, gin.H{
			"message": "warning",
		})
	})

	r.GET("/error", func(c *gin.Context) {
		logger.Println("ERROR: application failure occurred")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
	})

	logger.Println("INFO: server started on :8080")
	r.Run(":8080")
}
