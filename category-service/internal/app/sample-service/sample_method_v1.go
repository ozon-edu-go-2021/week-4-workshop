package sample_service

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/ozonmp/omp-grpc-template/pkg/sample-service"
)

func (i *Implementation) SampleMethodV1(
	ctx context.Context,
	req *desc.SampleMethodV1Request,
) (*desc.SampleMethodV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("SampleMethodV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.SampleMethodV1Response{
		Value: &desc.Template{
			Id:   req.GetId(),
			Name: "Sample",
		},
	}, nil
}
