package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/jmvargas/go-kit-example/basket/pkg/service"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Id int32 `json:"id"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Err error `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.BasketService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		err := s.Add(ctx, req.Id)
		return AddResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// ListRequest collects the request parameters for the List method.
type ListRequest struct{}

// ListResponse collects the response parameters for the List method.
type ListResponse struct {
	Lines []service.BasketLineWithProductData `json:"lines"`
	Error error                               `json:"error"`
}

// MakeListEndpoint returns an endpoint that invokes List on the service.
func MakeListEndpoint(s service.BasketService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		i0, e1 := s.List(ctx)
		return ListResponse{
			Error: e1,
			Lines: i0,
		}, nil
	}
}

// Failed implements Failer.
func (r ListResponse) Failed() error {
	return r.Error
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, id int32) (err error) {
	request := AddRequest{Id: id}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).Err
}

// List implements Service. Primarily useful in a client.
func (e Endpoints) List(ctx context.Context) (i0 []service.BasketLineWithProductData, e1 error) {
	request := ListRequest{}
	response, err := e.ListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListResponse).Lines, response.(ListResponse).Error
}
