package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}

// LogInfo logs an info message
func LogInfo(message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields(fields))
	entry.Info(message)
}

// LogWarn logs a warning message
func LogWarn(message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields(fields))
	entry.Warn(message)
}

// LogError logs an error message
func LogError(message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields(fields))
	entry.Error(message)
}

// LogDebug logs a debug message
func LogDebug(message string, fields map[string]interface{}) {
	entry := Logger.WithFields(logrus.Fields(fields))
	entry.Debug(message)
}

// SetLogLevel sets the log level
func SetLogLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	Logger.SetLevel(logLevel)
	return nil
}

// LoggingMiddleware logs incoming HTTP requests with path, method, and status code
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create a wrapped ResponseWriter to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		// Call the next handler
		next.ServeHTTP(wrapped, r)
		// Calculate duration
		duration := time.Since(start)
		// Log the request details
		LogInfo("HTTP Request", map[string]interface{}{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status_code": wrapped.statusCode,
			"duration":    duration.Seconds(),
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		})
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
