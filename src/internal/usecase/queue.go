package usecase

import (
	"context"
	"go-queue/internal/redis"
	"go-queue/internal/request"

	"github.com/labstack/echo/v4"
)

// Redisに保存する
func Queue(c echo.Context, ctx context.Context, r *request.QueueRequest) bool {
	err := redis.Set(ctx, r)
	if err == nil {
		return true
	} else {
		c.Logger().Error(err)
		return false
	}
}
