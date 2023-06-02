package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.stripchat.dev/myclub/go-service/internal/config"
	"math"
	"net/http"
	"os"
	"time"

	logrusSentry "github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

// timeFormat is the format for timestamps
var timeFormat = "2006-01-02 15:04:05 -0700"

// New ...
func New(config *config.Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeFormat,
	})

	if config.GetString("Runtime.Logger.Level") == "DEBUG" {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debugln("Debug mode enabled")
	}

	if config.GetString("Runtime.Logger.SentryDsn") == "" {
		return logger
	}

	tags := map[string]string{
		"type": "go-service",
	}

	levels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.TraceLevel,
	}

	hook, err := logrusSentry.NewWithTagsSentryHook(config.GetString("Runtime.Logger.SentryDsn"), tags, levels)

	if err == nil {
		hook.Timeout = 5 * time.Second
		hook.StacktraceConfiguration.Enable = true
		hook.SetEnvironment(config.GetString("env.environment"))

		logger.Hooks.Add(hook)
	}

	return logger
}

// Logger is the logrus logger handler
func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
