package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Logger1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(w, r)
		duration := time.Since(start)
		var sugar zap.SugaredLogger
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		sugar = *logger.Sugar()
		sugar.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
		)
	})
}
