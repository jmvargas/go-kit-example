package service

import (
	"context"

	"github.com/jmvargas/go-kit-example/catalog/pkg/io"
)

// CatalogService describes the service.
type CatalogService interface {
	Get(ctx context.Context, id int32) (product io.Product, err error)
	List(ctx context.Context) (products []io.Product, err error)
}

type basicCatalogService struct {
	repository io.ProductRepository
}

func (b *basicCatalogService) Get(ctx context.Context, id int32) (product io.Product, err error) {
	// TODO implement the business logic of Get
	return b.repository.Get(id)
}
func (b *basicCatalogService) List(ctx context.Context) (products []io.Product, err error) {
	// TODO implement the business logic of List
	return b.repository.All()
}

// NewBasicCatalogService returns a naive, stateless implementation of CatalogService.
func NewBasicCatalogService() CatalogService {
	return &basicCatalogService{
		repository: io.NewProductRepository(),
	}
}

// New returns a CatalogService with all of the expected middleware wired in.
func New(middleware []Middleware) CatalogService {
	var svc CatalogService = NewBasicCatalogService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
