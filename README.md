# go-kit-example
Proof of concept using go kit

This proof of concept has two services:

- Catalog exposes its API in HTTP and GRPC.
- Basket exposes its API using HTTP. It make requests to Catalog service using GRPC.
