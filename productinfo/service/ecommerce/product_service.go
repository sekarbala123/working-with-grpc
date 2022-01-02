package ecommerce

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	counter int
	ProductInfoServer
	productMap map[string]*Product
}

func (s *Server) AddProduct(ctx context.Context, in *Product) (*ProductID, error) {
	out := s.createProductID()
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*Product)
	}
	s.productMap[in.Id] = in
	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	return &ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *Server) createProductID() ProductID {
	var ProductID ProductID
	counter := s.counter + 1
	ProductID.Value = fmt.Sprintf("%02d", counter)
	return ProductID
}

func (s *Server) GetProduct(ctx context.Context, in *ProductID) (*Product, error) {
	log.Printf("Request for product with productid %v", in)
	return s.productMap[in.Value], nil
}
