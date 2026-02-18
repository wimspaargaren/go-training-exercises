package postgres

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgresTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}

type PostgresTestSuite struct {
	suite.Suite
	pool           *dockertest.Pool
	resource       *dockertest.Resource
	todoRepository *TodoRepository
}

func (s *PostgresTestSuite) SetupSuite() {
	var err error
	s.pool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	err = s.pool.Client.Ping()
	s.Require().NoError(err)

	s.resource, err = s.pool.Run("postgres", "18.0", []string{"POSTGRES_PASSWORD=postgres", "POSTGRES_DB=todo_db", "POSTGRES_USER=postgres"})
	s.Require().NoError(err)

	dsn := "host=localhost port=" + s.resource.GetPort("5432/tcp") + " user=postgres dbname=todo_db password=postgres sslmode=disable"
	for range 10 {
		fmt.Printf("Try to connect to db: %s\n", dsn)
		_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			TranslateError: true,
		})
		if connErr, ok := errors.AsType[*pgconn.ConnectError](err); ok {
			fmt.Printf("Connection error: %s docker container might be still starting\n", connErr)
			time.Sleep(time.Second)
			continue
		}
		if err == nil {
			break
		}
		s.T().Fatal("unexpected error during connection to the database:", err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	s.Require().NoError(err)
	s.todoRepository = NewTodoRepository(db)
	err = db.AutoMigrate(&Todo{})
	s.Require().NoError(err)
}

func (s *PostgresTestSuite) TearDownSuite() {
	if err := s.pool.Purge(s.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

// Write tests to verify that the TodoRepository methods work as expected.
// The TodoRepository should be accessible via s.todoRepository and has
// the following methods:
// - Create(ctx context.Context, todo *Todo) error
// - List(ctx context.Context) ([]*Todo, error)
// - GetByID(ctx context.Context, id uuid.UUID) (*Todo, error)
