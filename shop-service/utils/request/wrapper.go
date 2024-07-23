package request

import (
	"context"
	"fmt"
	"net/http"
	"shop-service/logger"
	"time"
)

// New ...
func (r *component) wrapper(ctx context.Context, payload []byte, method string) ([]byte, int, error) {
	start := time.Now()
	var tp logger.ThirdParty

	tp.Method = method
	tp.URL = r.url
	tp.RequestHeader = headerToString(&r.header)

	if payload != nil {
		tp.RequestBody = string(payload)
	}

	data, status, err := r.do(payload, method)

	tp.StatusCode = status
	if err != nil {
		tp.Response = err.Error()
	}

	switch status {
	case http.StatusOK:
		tp.Response = "request success"
	default:
		tp.Response = "response is empty"
		if data != nil {
			tp.Response = string(data)
		}
	}

	tp.ExecTime = time.Since(start).Seconds()

	/* storing third party request and response to context */
	tp.Store(ctx)

	return data, status, err
}

func headerToString(h *map[string]string) string {
	var header string

	if h == nil {
		return header
	}

	for key, v := range *h {
		header += fmt.Sprintf(` %v: %v `, key, v)
	}
	return header
}
