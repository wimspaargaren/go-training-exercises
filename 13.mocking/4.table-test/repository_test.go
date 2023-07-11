package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInfo(t *testing.T) {
	tests := []struct {
		Name               string
		SetupMockBehaviour func(dataStore *MockDataStore)
		ExpectedOutput     *Info
		ExpectErr          bool
	}{
		{
			Name: "Successfully create info",
		},
		{
			Name: "Creating the info failed",
		},
	}
	// Refactor the current test to a table test
	// and fill in the test cases above.
	// Hint: if you forgot how table test work and look like,
	// have a look at 6.testing/3.table-tests
	_ = tests
	dataStore := NewMockDataStore(t)

	dataStore.On("Store", &Info{X: 42, Y: 42}).Return(nil)
	infoHandler := NewInfoHandler(dataStore)
	result, err := infoHandler.CreateNewInfo(42, 42)
	assert.NoError(t, err)
	assert.Equal(t, &Info{
		ID: 1,
		X:  42,
		Y:  42,
	}, result)
}
