// controller/graphql/graphql.go
package graphql

import (

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc{
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
		GraphiQL: true,
	})

	// 只需要通过Gin简单封装即可
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}