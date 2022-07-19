package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hmhuan/simple-bank/token"
	"github.com/twinj/uuid"
)

const (
	authorizationHeader     = "Authorization"
	authorizationTypeBearer = "Bearer"
	authorizationPayloadKey = "authorization_payload"
)

var (
	ErrInvalidAuthorizationHeader    error = errors.New("Invalid authorization header")
	ErrNoAuthorizationHeader         error = errors.New("No authorization header")
	ErrNotSupportedAuthorizationType error = errors.New("Not supported authorization type")
)

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errorResponse(ErrNoAuthorizationHeader)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errorResponse(ErrInvalidAuthorizationHeader)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := fields[0]
		if authorizationType != authorizationTypeBearer {
			err := errorResponse(ErrNotSupportedAuthorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
