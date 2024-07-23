package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"product-service/logger"
	"product-service/utils"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

var (
	HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "myapp_http_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: utils.DefaultBuckets,
	}, []string{"code", "method", "path"})
)

// PrometheusMiddleware implements mux.MiddlewareFunc.
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Whitelist(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		sc := GetStatusCode{ResponseWriter: w}
		code := strconv.Itoa(sc.status)
		timer := prometheus.NewTimer(HttpDuration.WithLabelValues(code, r.URL.Path, r.Method))
		next.ServeHTTP(&sc, r)
		timer.ObserveDuration()
	})
}

// GetStatusCode is struct
type GetStatusCode struct {
	http.ResponseWriter
	status int
}

// WriteHeader is func test to superflous
func (w *GetStatusCode) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

/*Whitelist parts
 * @updated: Tuesday, December 3rd, 2019.
 * --
 * @param	urls	string
 * @return	mixed
 */
func Whitelist(urls string) bool {
	ws := regexp.MustCompile(`^\/(health|metrics|favicon.ico|liveness|debug[\w|\W]*)`)
	return ws.MatchString(urls)
}
