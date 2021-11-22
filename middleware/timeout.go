package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	defaultTimeout = 3 * time.Second
)

func ContextTimeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if timeout == 0 {
			timeout = defaultTimeout
		}

		ctx, cancel := context.WithTimeout(c, timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{}, 1)
		c.Next() // 执行handler chain

		go func() {
			done <- struct{}{} // handler chain执行完毕标记
		}()

		select {
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"success": false, "error": ctx.Err().Error()})
			return
		case <-done:
			return
		}
	}
}