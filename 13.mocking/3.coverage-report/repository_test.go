package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Run make test. This will run go test with a coverage profile,
// which will automatically openin your browser. See if you can
// increase the coverage to 100%
// Also note the functioncoverage.out which will report the code
// coverage per function.
func TestCreateInfoRun(t *testing.T) {
	dataStore := NewMockDataStore(t)

	dataStore.EXPECT().Store(&Info{X: 42, Y: 42}).Return(nil)
	infoHandler := NewInfoHandler(dataStore)
	result, err := infoHandler.CreateNewInfo(42, 42)
	assert.NoError(t, err)
	assert.Equal(t, &Info{
		X: 42,
		Y: 42,
	}, result)
}
