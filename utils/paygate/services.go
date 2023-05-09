package paygate

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PayService interface {
	GeneratePayReq(ctx context.Context, name string, email string, orderid string, total int64) (string, string, error)
	GetNotification(ctx echo.Context, orderid string) string
}
