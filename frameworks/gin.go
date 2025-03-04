package frameworks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gin() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Gin")
	})

	fmt.Println("Gin running on :3001")
	router.Run(":3001")
}
