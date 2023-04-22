package usecase

import (
	"context"
	"go-queue/internal/redis"

	"github.com/labstack/echo/v4"
)

func QueueExec(c echo.Context, ctx context.Context) bool {
	redis.AllGet(ctx)
	return true
}
