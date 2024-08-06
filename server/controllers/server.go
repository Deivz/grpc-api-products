package controllers

import (
	"context"
	"src/pb"
	"sync"

	"github.com/google/uuid"
)

type server struct {
	pb.UnimplementedProductsServiceServer
	mu            sync.Mutex
	streamClients map[pb.ProductsService_ListProductsServer]chan *pb.ProductResponse
}

func NewServer() *server {
	return &server{
		streamClients: make(map[pb.ProductsService_ListProductsServer]chan *pb.ProductResponse),
	}
}

func (s *server) ListProducts(req *pb.Empty, stream pb.ProductsService_ListProductsServer) error {
	s.mu.Lock()
	productChan := make(chan *pb.ProductResponse)
	s.streamClients[stream] = productChan
	s.mu.Unlock()

	products := GetAll()
	for _, product := range products {
		if err := stream.Send(&pb.ProductResponse{
			Uuid:        product.Uuid,
			Name:        product.Name,
			Type:        product.Type,
			Price:       product.Price,
			Description: product.Description,
		}); err != nil {
			return err
		}
	}

	for product := range productChan {
		if err := stream.Send(product); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductUuidRequest) (*pb.ProductResponse, error) {
	// Implement your logic to get a product by UUID here
	return &pb.ProductResponse{
		Uuid:        req.Uuid,
		Name:        "Product 1",
		Type:        "Type 1",
		Price:       "10.00",
		Description: "Description 1",
	}, nil
}

func (s *server) SaveProducts(ctx context.Context, req *pb.ProductRequest) (*pb.Success, error) {
	newProduct := &pb.ProductResponse{
		Uuid:        uuid.New().String(),
		Name:        req.Name,
		Type:        req.Type,
		Price:       req.Price,
		Description: req.Description,
	}

	Create(newProduct.Uuid, newProduct.Name, newProduct.Type, newProduct.Price, newProduct.Description)

	s.notifyClients(newProduct)

	return &pb.Success{
		Message: "Product saved successfully",
	}, nil
}

func (s *server) notifyClients(product *pb.ProductResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ch := range s.streamClients {
		ch <- product
	}
}
