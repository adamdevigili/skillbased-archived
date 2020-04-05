package middleware

import (
	"github.com/adamdevigili/balancer.team/pkg/constants"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = uuid.New().String()
			}
			c.Set(constants.RequestIDKey, rid)
			res.Header().Set(echo.HeaderXRequestID, rid)

			return next(c)
		}
	}
}
