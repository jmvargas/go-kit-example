package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	io "github.com/jmvargas/go-kit-example/catalog/pkg/io"
	service "github.com/jmvargas/go-kit-example/catalog/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Id int32 `json:"id"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Product io.Product `json:"product"`
	Err     error      `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.CatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		product, err := s.Get(ctx, req.Id)
		return GetResponse{
			Err:     err,
			Product: product,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// ListRequest collects the request parameters for the List method.
type ListRequest struct{}

// ListResponse collects the response parameters for the List method.
type ListResponse struct {
	Products []io.Product `json:"products"`
	Err      error        `json:"err"`
}

// MakeListEndpoint returns an endpoint that invokes List on the service.
func MakeListEndpoint(s service.CatalogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.List(ctx)
		return ListResponse{
			Err:      err,
			Products: products,
		}, nil
	}
}

// Failed implements Failer.
func (r ListResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, id int32) (product io.Product, err error) {
	request := GetRequest{Id: id}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Product, response.(GetResponse).Err
}

// List implements Service. Primarily useful in a client.
func (e Endpoints) List(ctx context.Context) (products []io.Product, err error) {
	request := ListRequest{}
	response, err := e.ListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListResponse).Products, response.(ListResponse).Err
}
