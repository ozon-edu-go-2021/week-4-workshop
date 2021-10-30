package sample_service

import (
	desc "github.com/ozonmp/omp-grpc-template/pkg/sample-service"
)

type Implementation struct {
	desc.UnimplementedSampleServiceServer
}

func NewSampleService() desc.SampleServiceServer {
	return &Implementation{}
}
