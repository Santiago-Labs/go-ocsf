module github.com/Santiago-Labs/go-ocsf

go 1.24

toolchain go1.24.2

require (
	cloud.google.com/go/logging v1.13.0
	github.com/apache/arrow/go/v15 v15.0.2
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.7
	github.com/aws/aws-sdk-go-v2/service/athena v1.50.2
	github.com/aws/aws-sdk-go-v2/service/glue v1.108.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.41.1
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.36.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.79.1
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.2.2
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.57.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/marcboeker/go-duckdb/v2 v2.1.0
	github.com/parquet-go/parquet-go v0.25.0
	github.com/samsarahq/go/oops v0.0.0-20220211150445-4b291d6feac4
	google.golang.org/api v0.228.0
	google.golang.org/genproto v0.0.0-20250324211829-b45e905df463
	google.golang.org/genproto/googleapis/api v0.0.0-20250227231956-55c901821b1e
)

require (
	cloud.google.com/go v0.118.3 // indirect
	cloud.google.com/go/auth v0.15.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/iam v1.4.0 // indirect
	cloud.google.com/go/longrunning v0.6.4 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/apache/arrow-go/v18 v18.1.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.60 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.15 // indirect
	github.com/aws/smithy-go v1.22.2 // indirect
	github.com/duckdb/duckdb-go-bindings v0.1.13 // indirect
	github.com/duckdb/duckdb-go-bindings/darwin-amd64 v0.1.8 // indirect
	github.com/duckdb/duckdb-go-bindings/darwin-arm64 v0.1.8 // indirect
	github.com/duckdb/duckdb-go-bindings/linux-amd64 v0.1.8 // indirect
	github.com/duckdb/duckdb-go-bindings/linux-arm64 v0.1.8 // indirect
	github.com/duckdb/duckdb-go-bindings/windows-amd64 v0.1.8 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/google/flatbuffers v25.1.24+incompatible // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.6 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/marcboeker/go-duckdb/arrowmapping v0.0.6 // indirect
	github.com/marcboeker/go-duckdb/mapping v0.0.6 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.59.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.59.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/exp v0.0.0-20250128182459-e0ece0dbea4c // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	golang.org/x/tools v0.29.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250313205543-e70fdf4c4cb4 // indirect
	google.golang.org/grpc v1.71.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
