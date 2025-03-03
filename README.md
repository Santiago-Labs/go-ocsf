# go-ocsf

[![CI](https://github.com/Santiago-Labs/go-ocsf/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Santiago-Labs/go-ocsf/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Santiago-Labs/go-ocsf.svg)](https://pkg.go.dev/github.com/Santiago-Labs/go-ocsf)
[![License](https://img.shields.io/github/license/Santiago-Labs/go-ocsf)](LICENSE)

`go-ocsf` is a Go library and CLI tool for converting security events from your security tools (e.g., Snyk) into the [Open Cybersecurity Schema Framework (OCSF)](https://schema.ocsf.io/) format, with output options in JSON or Parquet formats. Data can be stored locally or seamlessly uploaded to AWS S3.

## Features

- ⚙️ Converts security event data into OCSF-compliant format
- 📦 Supports JSON and Parquet output formats
- ☁️ Direct integration with AWS S3 for cloud storage
- 🖥️ Use as a CLI tool or Go library

## Installation

```bash
go get github.com/Santiago-Labs/go-ocsf
```

## Quick Start

Set environment variables required for your data source (e.g., Snyk):

```bash
export SNYK_API_KEY="your-snyk-api-key"
export SNYK_ORGANIZATION_ID="your-snyk-org-id"
```

Run the CLI to convert data and store locally as Parquet:

```bash
go run main.go --parquet
```

Store data directly in AWS S3:

```bash
export AWS_ACCESS_KEY_ID="your-aws-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-aws-secret-access-key"
export AWS_REGION="your-aws-region"

go run main.go --parquet --bucket-name="your-s3-bucket-name"
```

## Library Usage

You can embed the functionality directly in your Go code:

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/syncers"
)

func main() {
	ctx := context.Background()

	snykClient, err := snyk.NewClient(ctx, os.Getenv("SNYK_API_KEY"), os.Getenv("SNYK_ORGANIZATION_ID"))
	if err != nil {
		log.Fatal(err)
	}

	storage, err := datastore.NewLocalParquetDatastore()
	if err != nil {
		log.Fatal(err)
	}

	syncer, err := syncers.NewSnykOCSFSyncer(ctx, snykClient, storage)
	if err != nil {
		log.Fatal(err)
	}

	if err := syncer.Sync(ctx); err != nil {
		log.Fatal(err)
	}
}
```

## Supported Integrations

- Snyk
- AWS Inspector (coming soon)
- AWS GuardDuty (coming soon)
- Crowdstrike Falcon (coming soon)
- Google Workspace Logs (coming soon)
- Tenable (coming soon)
- AWS CloudTrail (coming soon)

## Contributing

We welcome contributions to improve or expand functionality.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add my feature'`)
4. Push to your branch (`git push origin feature/my-feature`)
5. Open a pull request

## License

`go-ocsf` is licensed under the [AGPL-3.0 License](LICENSE).
