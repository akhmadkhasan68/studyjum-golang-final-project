package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type logger struct {
	Level    string `json:"lvl"`
	Time     string `json:"time"`
	Msg      string `json:"msg"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Status   int    `json:"status"`
	Latency  string `json:"latency"`
	ClientIP string `json:"client_ip"`
	TraceID  string `json:"trace_id"`
	ErrorMsg string `json:"error,omitempty"`
}

func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		byteLogger, _ := json.Marshal(
			logger{
				Level:    "info",
				Time:     time.Now().Format("2006-01-02T15:04:05.999+0700"),
				Msg:      "request completed",
				Method:   params.Method,
				Path:     params.Path,
				Status:   params.StatusCode,
				Latency:  params.Latency.String(),
				ClientIP: params.ClientIP,
				TraceID:  params.Request.Header.Get("X-Request-ID"),
				ErrorMsg: params.ErrorMessage,
			},
		)

		return fmt.Sprintf("%s\n", string(byteLogger))
	})
}
