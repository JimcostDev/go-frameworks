package frameworks

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Echo() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from Echo")
	})

	fmt.Println("Echo running on :3002")
	e.Start(":3002")
}
