/**
 * @Author: duke
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2020/6/18 3:45 下午
 */

package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"lottery/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	server *http.Server
	lg     *zap.Logger
}

func NewHttpServer(lg *zap.Logger) (*HttpServer, error) {
	server := &HttpServer{
		lg:     lg,
		server: &http.Server{},
	}

	if err := server.initRouter(); err != nil {
		return nil, err
	}

	return server, nil
}

func (h *HttpServer) initRouter() error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(middleware.RequestLog(), middleware.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "data": "HELLO WORD", "time": time.Now()})
		return
	})
	r.POST("/", func(c *gin.Context) {

	})

	h.server.Handler = r
	return nil
}

// Start 启动服务
func (h *HttpServer) Start(l net.Listener) {
	log.Printf("http server started on port: %s", l.Addr().String())
	if err := h.server.Serve(l); err != nil {
		log.Println(err)
	}
	return
}

func (h *HttpServer) ShutDown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
