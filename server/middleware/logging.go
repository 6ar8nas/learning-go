package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/6ar8nas/learning-go/server/types"
	"github.com/6ar8nas/learning-go/server/utils"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter, requestId string) *loggingResponseWriter {
	w.Header().Set(types.HeaderXRequestId, requestId)
	return &loggingResponseWriter{w, http.StatusContinue}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

var requestIdInc = utils.AutoIncrement{}

func assignRequestId(r *http.Request) (*http.Request, string) {
	id := fmt.Sprintf("%d", requestIdInc.Next())
	return r.WithContext(utils.AssignContextValue(r.Context(), types.ContextKeyRequestId, id)), id
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r, id := assignRequestId(r)
		lrw := newLoggingResponseWriter(w, id)
		log.Printf("[%s]: %s %s", id, r.Method, r.URL.Path)
		next.ServeHTTP(lrw, r)
		log.Printf("[%s]: %d %s %s %s", id, lrw.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
