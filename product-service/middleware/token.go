package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"product-service/logger"
	"product-service/utils/request"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Key int

const (
	bearerToken = "Bearer"
	CtxUserId   = Key(10)
)

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
type Token struct {
	TokenID       string    `json:"TokenID"`
	UserAccountID string    `json:"UserAccountID"`
	AuthToken     string    `json:"AuthToken"`
	CreatedAt     time.Time `json:"CreatedAt"`
	ExpiredAt     time.Time `json:"ExpiredAt"`
}
type RequestToken struct {
	Token string `json:"token"`
}

type ResponseToken struct {
	Status *Status `json:"status"`
	Token  *Token  `json:"data"`
}

func (c *client) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var (
			ctx           = r.Context()
			authorization = r.Header.Get("Authorization")
			err           error
			deliveryHost  = viper.GetString("client.user.host")
			url           = fmt.Sprintf("%s%s", deliveryHost, viper.GetString("client.user.checktoken"))
			resp          = &ResponseToken{}
		)

		if reflect.ValueOf(authorization).IsZero() {
			err = fmt.Errorf("authorization header not found")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidToken), message[InvalidToken], reason[InvalidToken], err)
			return
		}

		token := strings.Split(authorization, " ")
		if token[0] != bearerToken {
			err = fmt.Errorf("bearer token salah. %s", authorization)
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidBearer), message[InvalidBearer], reason[InvalidBearer], err)
			return
		}
		req := &RequestToken{
			Token: token[1],
		}

		payload, err := json.Marshal(req)
		if err != nil {
			err = fmt.Errorf("invalid marshal json")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidBearer), "Invalid Request", "Invalid Marshal", err)
			return
		}

		header := map[string]string{
			"Content-Type": "application/json",
		}

		client := request.NewHTTPClient(c.httpClient)
		reqHttp := client.Request(header, url)

		res, statusCode, err := reqHttp.Post(ctx, payload)
		if err != nil {
			logger.Messagef("error request to user service %v", err).To(ctx)
			err = fmt.Errorf("error request to user service")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidBearer), "Error User Service", "Error Request to User Service", err)
			return
		}

		err = json.Unmarshal(res, &resp)
		if err != nil {
			err = fmt.Errorf("invalid umarshal json")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidBearer), "Invalid Response", "Invalid UMarshal", err)
			return
		}

		if statusCode >= 400 {
			err = fmt.Errorf("response error")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidToken), resp.Status.Message, resp.Status.Reason, err)
			return
		}

		// store to context
		ctx = context.WithValue(
			ctx,
			CtxUserId,
			resp.Token.UserAccountID,
		)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
