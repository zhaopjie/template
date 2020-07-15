/**
 * @Author: duke
 * @Description:
 * @File:  request_log
 * @Version: 1.0.0
 * @Date: 2020/6/19 4:01 下午
 */

package middleware

import (
	"time"

	"lottery/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		d := time.Since(t)

		fields := []zapcore.Field{
			logger.String("host", c.Request.Host),
			logger.String("method", c.Request.Method),
			logger.Int("status_code", c.Writer.Status()),
			logger.Float64("cost", float64(d)/float64(time.Millisecond)), // ms
			logger.String("url", c.Request.URL.Path),
			logger.String("query", c.Request.URL.RawQuery),
			logger.String("form", c.Request.Form.Encode()),
			logger.String("ip", c.ClientIP()),
		}
		logger.GetLogger().Info("request-log", fields...)
	}
}
