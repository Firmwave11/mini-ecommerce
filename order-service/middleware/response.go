package middleware

import (
	"context"
	"net/http"

	"order-service/utils/response"
)

type Code string

const (
	InvalidBearer Code = "010104"
	InvalidToken  Code = "010105"
	UserInactive  Code = "010106"
	Marshal       Code = "010190"
)

var (
	message = map[Code]string{
		InvalidBearer: "bearer.invalid",
		InvalidToken:  "token.invalid",
		UserInactive:  "user.inactive",
		Marshal:       "system.error",
	}

	reason = map[Code]string{
		InvalidBearer: "Invalid Token",
		InvalidToken:  "Unauthenticated",
		UserInactive:  "User Inactive",
		Marshal:       "Server error",
	}
)

func (c *client) errorResponse(ctx context.Context, w http.ResponseWriter, statusCode int, systemCode, message, reason string, err error) {
	response.Errors(ctx, statusCode, systemCode, message, reason, err).Middleware(w)
}
