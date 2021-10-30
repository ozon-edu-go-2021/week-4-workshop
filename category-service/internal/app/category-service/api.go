package category_service

import (
	desc "github.com/ozonmp/week-3-workshop/category-service/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedSampleServiceServer
}

func NewSampleService() desc.SampleServiceServer {
	return &Implementation{}
}
