package providers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func InitServer(port string, server *gin.Engine) error {
	err := server.Run(":" + port)
	if err != nil {
		return err
	}
	return nil
}
