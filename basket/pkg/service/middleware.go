package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BasketService) BasketService

type loggingMiddleware struct {
	logger log.Logger
	next   BasketService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BasketService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BasketService) BasketService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Add(ctx context.Context, id int32) (err error) {
	defer func() {
		l.logger.Log("method", "Add", "id", id, "err", err)
	}()
	return l.next.Add(ctx, id)
}
func (l loggingMiddleware) List(ctx context.Context) (i0 []BasketLineWithProductData, e1 error) {
	defer func() {
		l.logger.Log("method", "List", "i0", i0, "e1", e1)
	}()
	return l.next.List(ctx)
}
