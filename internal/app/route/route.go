package route

import "github.com/gin-gonic/gin"

func NewRoute(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "success",
		})
	})

}
