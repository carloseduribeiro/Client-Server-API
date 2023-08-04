package api

import (
	"context"
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
)

type ExchangeRepository interface {
	Save(ctx context.Context, dto *exchange.Exchange) error
}
