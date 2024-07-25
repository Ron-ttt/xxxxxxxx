package middleware

import (
	"compress/gzip"
	"context"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	cookies "github.com/Ron-ttt/xxxxxxxx/internal/app/cookie"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/utils"
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

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipDecode(r *http.Request) (io.ReadCloser, error) {
	if r.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		return gz, nil
	}

	return r.Body, nil
}

func GzipMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			var err error
			r.Body, err = gzipDecode(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			h.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		r.Body, err = gzipDecode(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Encoding", "gzip")
		h.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

type ContextKey string
type ToHand struct {
	Value  string
	IsAuth bool
}

// *    1. generate random names for users
// *    2. use read/writeEncrypted funcs
// TODO 3. list users func implementation
// TODO 4. no cry!!!
// ! JUST CRYYYYYYYYY
// *    5. create secret key
const secretKey2 = "13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b"

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var th ToHand
		secretKey, err := hex.DecodeString(secretKey2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		namecookie := "username"
		value, err := cookies.ReadEncrypted(r, namecookie, secretKey)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				username := utils.RandString(5)
				cookie := http.Cookie{
					Name:     namecookie,
					Value:    username,
					Path:     "/",
					HttpOnly: true,
					Secure:   false,
				}
				err1 := cookies.WriteEncrypted(w, cookie, secretKey)
				th.IsAuth = false
				th.Value = username
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			th.IsAuth = true
			th.Value = value
		}
		var key ContextKey = "Name"
		ctx := context.WithValue(r.Context(), key, th)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
