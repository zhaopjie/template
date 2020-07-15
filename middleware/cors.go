/**
 * @Author: duke
 * @Description:
 * @File:  cors
 * @Version: 1.0.0
 * @Date: 2020/6/22 10:26 上午
 */

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X_Requested_With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		path := strings.Split(c.Request.URL.Path, "/")
		if path[len(path)-2] == "file" {
			c.Writer.Header().Set("Content-Disposition", "attachment;filename='"+path[len(path)-1]+"'")
		}
		c.Next()
	}
}
