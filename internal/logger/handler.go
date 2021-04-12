// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package logger

import (
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = ksuid.New().String()
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))

		// optimization to avoid writing to the logs if
		// the log level is debug level or higher.
		if logrus.GetLevel() < logrus.InfoLevel {
			return
		}

		end := time.Now()
		log.WithFields(logrus.Fields{
			"http.method":  r.Method,
			"http.request": r.RequestURI,
			"http.remote":  r.RemoteAddr,
			"http.latency": end.Sub(start),
		}).Debug()
	})
}
