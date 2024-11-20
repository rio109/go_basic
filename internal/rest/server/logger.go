package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	InvalidFormat   = 1
	InvalidDate     = 2
	InvalidType     = 3
	InvalidSchedule = 4
	InternalError   = 100
)

const (
	green   = "\033[97;42m"
	yellow  = "\033[30;43m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	red     = "\033[97;101m"
)

func Logger(entry *log.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		url := c.Request.Host + c.Request.URL.Path
		query := c.Request.URL.RawQuery
		if len(query) > 0 {
			url = fmt.Sprintf("%s?%s", url, query)
		}
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		var codeColor string
		var level log.Level
		reset := "\033[0m"

		if statusCode >= http.StatusInternalServerError { // 5XX
			codeColor = red
			level = log.ErrorLevel
		} else if statusCode >= http.StatusBadRequest { // 4XX
			codeColor = yellow
			level = log.WarnLevel
		} else if statusCode >= http.StatusMultipleChoices { // 3XX. 이후 따로 처리해야할 수 있어 분리
			codeColor = magenta
			level = log.InfoLevel
		} else if statusCode >= http.StatusOK { // 2XX.
			codeColor = green
			level = log.InfoLevel
		} else { // 1XX. 이후 따로 처리해야할 수 있어 분리
			reset = ""
			level = log.InfoLevel
		}
		msg := fmt.Sprintf("|%s %d %s| %13v | %15s | %s %-5s%s %s",
			codeColor, statusCode, reset, latency, clientIP, blue, method, reset, url)

		entry.Log(level, msg)
	}
}
