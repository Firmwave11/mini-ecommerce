package middleware

import (
	"net/http"
	"time"

	"user-service/logger"

	"github.com/sirupsen/logrus"
)

// Logger using for record any request, response and some logmessages to stdout terminal
func (c *client) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now().In(c.loc)
		req, data := logger.Initialize(r, start)

		next.ServeHTTP(rw, req)

		data.Finalize(req.Context())
		c.write(&data)
	})
}

func (c *client) write(dl *logger.DataLogger) {
	var level logrus.Level

	if dl.StatusCode >= 200 && dl.StatusCode < 400 {
		level = logrus.InfoLevel
	} else if dl.StatusCode >= 400 && dl.StatusCode < 500 {
		level = logrus.WarnLevel
	} else {
		level = logrus.ErrorLevel
	}

	c.log.WithField("incoming_log", dl).Log(level, "apps")
}
