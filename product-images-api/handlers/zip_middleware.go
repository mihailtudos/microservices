package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw *http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(*rw)

	return &WrappedResponseWriter{
		rw: *rw,
		gw: gw,
	}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		{
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				//create a gzipped response
				grw := NewWrappedResponseWriter(&w)
				grw.Header().Set("Content-Encoding", "gzip")

				next.ServeHTTP(grw, r)
				defer grw.Flush()
				return
			}

			//handle normally
			next.ServeHTTP(w, r)
		}
	})
}
