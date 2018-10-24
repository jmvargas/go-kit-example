package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/jmvargas/go-kit-example/catalog/pkg/endpoint"
	pb "github.com/jmvargas/go-kit-example/catalog/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetRequest)
	return endpoint.GetRequest{Id: req.Id}, nil
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetResponse)
	return &pb.GetReply{
		Product: &pb.Product{
			Id:    res.Product.Id,
			Title: res.Product.Title,
			Price: res.Product.Price,
		},
	}, nil
}
func (g *grpcServer) Get(ctx context1.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

// makeListHandler creates the handler logic
func makeListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListEndpoint, decodeListRequest, encodeListResponse, options...)
}

// decodeListResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeListRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoint.ListRequest{}, nil
}

// encodeListResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeListResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ListResponse)
	products := []*pb.Product{}
	for _, product := range res.Products {
		products = append(products, &pb.Product{
			Id:    product.Id,
			Title: product.Title,
			Price: product.Price,
		})
	}
	return pb.ListReply{Products: products}, nil
}
func (g *grpcServer) List(ctx context1.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	_, rep, err := g.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListReply), nil
}
