package postgres

import (
	"github.com/stretchr/testify/suite"
)

// Re-do the previous exercise but now using testify suite
// https://pkg.go.dev/github.com/stretchr/testify/suite
// 1. Create a new struct PostgresTestSuite that embeds suite.Suite
// 2. Implement the SetupSuite method to start the PostgreSQL container
// 3. Implement the TearDownSuite method to purge the PostgreSQL container
// 4. Implement a test method TestConnection that tests the connection to the database
// 5. Run the test suite using suite.Run in a TestPostgresTestSuite function

type PostgresTestSuite struct {
	suite.Suite
}
