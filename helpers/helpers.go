package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return false
	}
	return true
}

func GetFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}
