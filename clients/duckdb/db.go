package duckdb

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/marcboeker/go-duckdb/v2"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	db   *sqlx.DB
	once sync.Once

	reservedKeywords = map[string]bool{
		"type": true,
	}
)

func NewS3Client(ctx context.Context, bucketName string) (*sqlx.DB, error) {
	var err error
	once.Do(func() {

		var execDir string
		execDir, err = os.Executable()
		if err != nil {
			return
		}
		execDir = path.Dir(execDir)

		duckDBDir := path.Join(execDir, "duckdb_data")
		err = os.MkdirAll(duckDBDir, 0755)
		if err != nil {
			return
		}

		dbPath := path.Join(duckDBDir, "duckdb.db")
		db, err = sqlx.Connect("duckdb", dbPath)
		if err != nil {
			return
		}
		db.Mapper = reflectx.NewMapperTagFunc("json", nil, nil)

		_, err = db.Exec(fmt.Sprintf(`
			SET home_directory='%s'; 
			SET extension_directory='%s';

			FORCE INSTALL aws FROM core_nightly;
			FORCE INSTALL httpfs FROM core_nightly;
			FORCE INSTALL iceberg FROM core_nightly;

			CREATE SECRET (
				TYPE s3,
				PROVIDER credential_chain
			);

		`, duckDBDir, duckDBDir))
		if err != nil {
			return
		}

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return
		}

		creds, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return
		}

		_, err = db.Exec(fmt.Sprintf(`
			ATTACH 'arn:aws:s3tables:%s:%s:bucket/%s'
			AS s3_tables_db (
				TYPE iceberg,
				ENDPOINT_TYPE s3_tables
			);

			SET search_path = 's3_tables_db';

		`, cfg.Region, creds.AccountID, bucketName))
		if err != nil {
			return
		}
	})

	return db, err
}

func NewLocalClient(ctx context.Context) (*sqlx.DB, error) {
	var err error
	once.Do(func() {

		var execDir string
		execDir, err = os.Executable()
		if err != nil {
			return
		}
		execDir = path.Dir(execDir)

		duckDBDir := path.Join(execDir, "duckdb_data")
		err = os.MkdirAll(duckDBDir, 0755)
		if err != nil {
			return
		}

		dbPath := path.Join(duckDBDir, "duckdb.db")
		db, err = sqlx.Connect("duckdb", dbPath)
		if err != nil {
			return
		}
		db.Mapper = reflectx.NewMapperTagFunc("json", nil, nil)

		_, err = db.Exec(fmt.Sprintf(`
			SET home_directory='%s'; 
			SET extension_directory='%s';

		`, duckDBDir, duckDBDir))
		if err != nil {
			return
		}
	})

	return db, err
}
