package middleware

import (
	"github.com/gin-gonic/gin"
	"solution/internal/service"
	"time"
)

type RequestsLoggerMiddleware struct {
	apiSvc *service.ApiService
}

func NewRequestsLoggerMiddleware(apiSvc *service.ApiService) *RequestsLoggerMiddleware {
	return &RequestsLoggerMiddleware{apiSvc}
}

func (m *RequestsLoggerMiddleware) Callback(c *gin.Context) {
	start := time.Now()
	c.Next()
	m.apiSvc.LogRequest(c.Request.Method, c.FullPath(), time.Since(start))
}
