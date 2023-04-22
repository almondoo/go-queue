package route

import (
	"context"
	"fmt"
	"go-queue/internal/request"
	"go-queue/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, ctx context.Context) {
	helloWorld(e, ctx)
	queue(e, ctx)
	queueExec(e, ctx)
}

func helloWorld(e *echo.Echo, ctx context.Context) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func queue(e *echo.Echo, ctx context.Context) {
	e.POST("/queue", func(c echo.Context) (err error) {
		// バリデーション
		r := new(request.QueueRequest)
		if err = c.Bind(r); err != nil {
			fmt.Println("Bindエラー")
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
		}

		// Queueにセットする
		var ok bool = usecase.Queue(c, ctx, r)

		if ok {
			return echo.NewHTTPError(http.StatusOK, "Queue Set Complete.")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "Queue Set Fail.")
		}
	})
}

// func check(e *echo.Echo, ctx context.Context) {
// 	e.GET('')
// }

func queueExec(e *echo.Echo, ctx context.Context) {
	e.POST("/queue/exec", func(c echo.Context) (err error) {
		usecase.QueueExec(c, ctx)
		return c.JSON(http.StatusOK, "ok")
	})
}
