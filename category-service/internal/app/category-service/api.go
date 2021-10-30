package category_service

import (
	desc "github.com/ozonmp/omp-grpc-template/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedSampleServiceServer
}

func NewSampleService() desc.SampleServiceServer {
	return &Implementation{}
}
