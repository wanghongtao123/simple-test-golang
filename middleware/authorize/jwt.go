package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-test-golang/util"
	"github.com/pkg/errors"
)

// Authorize 验证
func Authorize() gin.HandlerFunc {
    return func(c *gin.Context) {
		request := c.Request
		token := request.Header.Get("token")
		claims, err := authorize(token)
	}
}
 
// authorize 实现
func authorize(token string) (*Claims, error) {
	claims, err = util.ParseToken(token)
	return claims, err
}