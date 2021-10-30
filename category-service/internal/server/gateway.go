package server

import (
	"context"
	"net/http"
	"os"
	"strings"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	desc "github.com/ozonmp/omp-grpc-template/pkg/sample-service"
)

func createGatewayServer(grpcAddr, gatewayAddr string, allowedOrigins []string) *http.Server {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
		),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}

	mux := runtime.NewServeMux()
	if err := desc.RegisterSampleServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatal().Err(err).Msg("Failed registration handler")
	}

	mux.HandlePath("GET", "/swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		swaggerBytes, err := os.ReadFile("swagger/api/omp_grpc_template/omp_grpc_template.swagger.json")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(swaggerBytes)
	})

	docsServer := http.FileServer(http.Dir("./swagger/dist"))

	gatewayServer := &http.Server{
		Addr: gatewayAddr,
		Handler: cors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/docs" || strings.HasPrefix(r.URL.Path, "/docs/") {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, "/docs")
				docsServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}), allowedOrigins),
	}

	return gatewayServer
}

func cors(h http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providedOrigin := r.Header.Get("Origin")
		matches := false
		for _, allowedOrigin := range allowedOrigins {
			if providedOrigin == allowedOrigin {
				matches = true
				break
			}
		}

		if matches {
			w.Header().Set("Access-Control-Allow-Origin", providedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
