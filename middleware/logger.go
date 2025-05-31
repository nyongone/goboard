package middleware

import (
	"go-board-api/internal/logger"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

type responseWriter struct {
  http.ResponseWriter
  status	int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
  return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
  return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
  rw.status = code
  rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    defer func() {
      if err := recover(); err != nil {
        logger.Panic("Panic Occured", zap.Any("error", err), zap.Binary("stack", debug.Stack()))
      }
    }()

    wr := wrapResponseWriter(w)
    next.ServeHTTP(wr, r)
    logger.Info("Capture HTTP Request", zap.Int("status", wr.status),
                    zap.String("method", r.Method),
                    zap.String("path", r.URL.EscapedPath()),
                    zap.String("remote_addr", r.RemoteAddr))
  })
}