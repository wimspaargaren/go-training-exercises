package postgres

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
)

// In this test we will learn how to leverage ory/dockertest to run a PostgreSQL container.
// 1. Checkout the ory dockertest repository to see how you can start a container:
// https://github.com/ory/dockertest
// 2. Select a PostgreSQL image from Docker Hub: https://hub.docker.com/_/postgres
// 3. Use the gorm postgres driver to connect to the database: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

func TestPostgresDB(t *testing.T) {
	pool, err := dockertest.NewPool("")
	assert.NoError(t, err)

	err = pool.Client.Ping()
	assert.NoError(t, err)

	var resource *dockertest.Resource
	// Select an image https://hub.docker.com/_/postgres

	defer func() {
		err := pool.Purge(resource)
		assert.NoError(t, err)
	}()

	dsn := "host=localhost port=" + resource.GetPort("5432/tcp") + " user=postgres dbname=todo_db password=postgres sslmode=disable"
	for range 10 {
		fmt.Printf("Try to connect to db: %s\n", dsn)
		var err error
		// use the gorm library to open a connection to the database

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
}
