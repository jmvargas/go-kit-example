package service

import (
	"context"
	"log"

	"github.com/jmvargas/go-kit-example/basket/pkg/io"
	"github.com/jmvargas/go-kit-example/catalog/pkg/grpc/pb"
	"google.golang.org/grpc"
)

// BasketService describes the service.
type BasketService interface {
	Add(ctx context.Context, id int32) (err error)
	List(ctx context.Context) ([]BasketLineWithProductData, error)
}

type basicBasketService struct {
	catalogService   pb.CatalogClient
	basketRepository io.BasketLineRepository
}

func (b *basicBasketService) Add(ctx context.Context, id int32) error {
	return b.basketRepository.Increment(id)
}
func (b *basicBasketService) List(ctx context.Context) ([]BasketLineWithProductData, error) {
	basket := []BasketLineWithProductData{}
	all, err := b.basketRepository.All()
	if err != nil {
		return basket, err
	}
	for _, line := range all {
		catalogReply, err := b.catalogService.Get(ctx, &pb.GetRequest{
			Id: line.ProductId,
		})
		if err != nil {
			continue
		}
		basket = append(basket, BasketLineWithProductData{
			Title:      catalogReply.Product.Title,
			Price:      catalogReply.Product.Price,
			BasketLine: line,
		})
	}
	return basket, nil
}

// NewBasicBasketService returns a naive, stateless implementation of BasketService.
func NewBasicBasketService() BasketService {
	conn, err := grpc.Dial("catalog:8082", grpc.WithInsecure())
	if err != nil {
		log.Printf("fail after try to connecto to catalog: %s", err.Error())
		return &basicBasketService{}
	}
	return &basicBasketService{
		catalogService:   pb.NewCatalogClient(conn),
		basketRepository: io.NewBasketLineRepository(),
	}
}

// New returns a BasketService with all of the expected middleware wired in.
func New(middleware []Middleware) BasketService {
	var svc BasketService = NewBasicBasketService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
