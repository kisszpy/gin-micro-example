package micro

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_TestHello(t *testing.T) {
	println("start")
	println("end")
}
func Test_NewServiceRegister(t *testing.T) {
	NewServiceRegister("hello,world", 5991)
}
func Test_Start(t *testing.T) {
	port := 8808
	NewServiceRegister("user-service", port)
	// fixme using struct release this code
	Start(port, "/api/v1/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 2000,
		})
	})
}
