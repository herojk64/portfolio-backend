package response

import "github.com/gin-gonic/gin"

type Envelop struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, Envelop{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Envelop{
		Success: false,
		Error:   message,
	})
}
