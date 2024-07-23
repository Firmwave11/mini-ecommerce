package response

import (
	"context"
	"encoding/json"
	"net/http"

	"product-service/logger"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Reason    string `json:"reason,omitempty"`
		HasErrors bool   `json:"hasErrors,omitempty"`
	} `json:"status"`

	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
}

type Output interface {
	Write(c echo.Context) error

	Middleware(w http.ResponseWriter)
}

func (r *response) Middleware(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	json.NewEncoder(w).Encode(r)
}

func (r *response) Write(c echo.Context) error {
	return r.writer(c)
}

func (r *response) writer(c echo.Context) error {
	return c.JSON(r.StatusCode, r)
}

// return error response with json
func Errors(ctx context.Context, responseStatusCode int, statusCode, statusMessage, statusReason string, err error) Output {
	logger.Response(ctx, responseStatusCode, err.Error())

	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage
	res.Status.Reason = statusReason

	return &res
}

// return success response with json
func Success(ctx context.Context, responseStatusCode int, statusCode, statusMessage string, data interface{}) Output {
	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage
	logger.Response(ctx, responseStatusCode, res)

	res.Data = data

	return &res
}

// return success response with json
func SuccessWithReason(ctx context.Context, responseStatusCode int, statusCode, statusReason, statusMessage string, data interface{}) Output {
	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage
	res.Status.Reason = statusReason
	logger.Response(ctx, responseStatusCode, res)

	res.Data = data

	return &res
}

// return error response with json
func ErrorWithDataReason(ctx context.Context, responseStatusCode int, statusCode, statusReason, statusMessage string, data interface{}, err error) Output {
	logger.Response(ctx, responseStatusCode, err.Error())

	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage
	res.Status.Reason = statusReason
	res.Data = data

	return &res
}

// return success response with json
func SuccessWithMeta(ctx context.Context, responseStatusCode int, statusCode, statusMessage string, data interface{}, meta interface{}) Output {
	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage

	logger.Response(ctx, responseStatusCode, res)

	res.Data = data
	res.Meta = meta

	return &res
}

// return success response with json
func SuccessWithHasErrors(ctx context.Context, responseStatusCode int, statusCode, statusMessage string, hasErrors bool, errors interface{}) Output {
	res := response{}
	res.StatusCode = responseStatusCode
	res.Status.Code = statusCode
	res.Status.Message = statusMessage
	res.Status.HasErrors = hasErrors
	logger.Response(ctx, responseStatusCode, res)

	res.Errors = errors

	return &res
}
