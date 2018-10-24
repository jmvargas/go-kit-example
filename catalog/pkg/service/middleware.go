package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	io "github.com/jmvargas/go-kit-example/catalog/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(CatalogService) CatalogService

type loggingMiddleware struct {
	logger log.Logger
	next   CatalogService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a CatalogService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CatalogService) CatalogService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context, id int32) (product io.Product, err error) {
	defer func() {
		l.logger.Log("method", "Get", "id", id, "product", product, "err", err)
	}()
	return l.next.Get(ctx, id)
}
func (l loggingMiddleware) List(ctx context.Context) (products []io.Product, err error) {
	defer func() {
		l.logger.Log("method", "List", "products", products, "err", err)
	}()
	return l.next.List(ctx)
}
