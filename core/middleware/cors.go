package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("CorsMiddleware")
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("RequestMethod:", r.Method)
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, DELETE, POST")
			w.Header().Add("Access-Control-Allow-Credentail", "true")

			return
		}
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
