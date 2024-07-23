package logger

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
)

// Initialize for init first context
func Initialize(req *http.Request, start time.Time) (*http.Request, DataLogger) {
	var (
		lock = new(Locker)
		dl   DataLogger
	)

	dl.RequestID = uuid.New().String()
	dl.Host = req.Host
	dl.Endpoint = req.URL.Path
	dl.TimeStart = start
	dl.RequestHeader = dumpRequest(req)
	dl.RequestBody = dumpBody(req)

	ctx := context.WithValue(req.Context(), logKey, lock)

	return req.WithContext(ctx), dl
}

// dumpHeader for getting all request from Header
func dumpRequest(req *http.Request) string {
	header, err := httputil.DumpRequest(req, false)
	if err != nil {
		return "dump request header failed"
	}

	trim := bytes.ReplaceAll(header, []byte("\r\n"), []byte("  "))
	return string(trim)
}

// dumpBody for getting all request from payload body
func dumpBody(req *http.Request) string {
	// exctract all payload
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return ""
	}

	// put again body to payload request
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return string(buf)
}

// Response is record response
func Response(ctx context.Context, status int, res interface{}) {
	value, ok := extract(ctx)
	if !ok {
		return
	}

	value.Set(_StatusCode, status)
	value.Set(_Response, res)
}
