package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/6ar8nas/learning-go/server/types"
	"github.com/6ar8nas/learning-go/shared/utils"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

const HeaderXRequestId = "X-Request-Id"

func newLoggingResponseWriter(w http.ResponseWriter, requestId string) *loggingResponseWriter {
	w.Header().Set(HeaderXRequestId, requestId)
	return &loggingResponseWriter{w, http.StatusContinue}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func assignRequestId(r *http.Request, autoinc *utils.AutoIncrement) (*http.Request, string) {
	id := fmt.Sprintf("%d", autoinc.Next())
	return r.WithContext(utils.AssignContextValue(r.Context(), types.ContextKeyRequestId, id)), id
}

func Logging(next http.Handler) http.Handler {
	var requestIdInc = utils.AutoIncrement{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r, id := assignRequestId(r, &requestIdInc)
		lrw := newLoggingResponseWriter(w, id)
		log.Printf("[%s]: %s %s", id, r.Method, r.URL.Path)
		next.ServeHTTP(lrw, r)
		log.Printf("[%s]: %d %s %s %s", id, lrw.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
