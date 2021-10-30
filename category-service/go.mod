module github.com/ozonmp/omp-grpc-template

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.3
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozonmp/omp-grpc-template/pkg/sample-service v0.0.1
	github.com/rs/zerolog v1.24.0
	google.golang.org/genproto v0.0.0-20211021150943-2b146023228c
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozonmp/omp-grpc-template/pkg/sample-service => ./pkg/sample-service
