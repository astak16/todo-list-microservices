package middleware

import (
	"github.com/gin-gonic/gin"
)

type ProtoClient struct {
	Key   string
	Value interface{}
}

func InitMiddleware(clients []ProtoClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		for _, client := range clients {
			c.Keys[client.Key] = client.Value
		}
		c.Next()
	}
}
