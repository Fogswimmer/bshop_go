package helpers

import (
	"net/http"
	"time"

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

func FormatBD(bd string) (string, error) {
	fmtBD, err := time.Parse("2006-01-02", bd)
	if err != nil {
		return "", err
	}
	return fmtBD.Format("02 Jan 2006"), nil
}
