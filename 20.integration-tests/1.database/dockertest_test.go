package postgres

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// In this test we will learn how to leverage ory/dockertest to run a PostgreSQL container.
// 1. Checkout the ory dockertest repository to see how you can start a container:
// https://github.com/ory/dockertest
// 2. Select a PostgreSQL image from Docker Hub: https://hub.docker.com/_/postgres
// 3. Use the gorm postgres driver to connect to the database: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

func TestPostgresDB(t *testing.T) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Select an image https://hub.docker.com/_/postgres

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

	}()

	dsn := "host=localhost port=" + resource.GetPort("5432/tcp") + " user=postgres dbname=todo_db password=postgres sslmode=disable"
	for i := 0; i < 10; i++ {
		fmt.Printf("Try to connect to db: %s\n", dsn)
		_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			TranslateError: true,
		})
		var connErr *pgconn.ConnectError
		if ok := errors.As(err, &connErr); ok {
			fmt.Printf("Connection error: %s docker container might be still starting\n", connErr)
			time.Sleep(time.Second)
			continue
		}
		if err == nil {
			break
		}
		t.Fatal("unexpected error during connection to the database:", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet

	// as of go1.15 testing.M returns the exit code of m.Run(), so it is safe to use defer here
}
