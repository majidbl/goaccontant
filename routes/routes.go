package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {

	// gin.H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}
