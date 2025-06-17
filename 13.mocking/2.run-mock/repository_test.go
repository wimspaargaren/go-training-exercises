package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInfoRun(t *testing.T) {
	dataStore := NewMockDataStore(t)

	dataStore.EXPECT().Store(&Info{X: 42, Y: 42}).Return(nil).Run(func(info *Info) {
		fmt.Println("INFO", info)
		// Add assert statements for the values of X and Y.
		// Assign the value 1 to the ID as normally would have been done
		// by a real database.
	}).Return(nil)
	infoHandler := NewInfoHandler(dataStore)
	result, err := infoHandler.CreateNewInfo(42, 42)
	assert.NoError(t, err)
	assert.Equal(t, &Info{
		ID: 1,
		X:  42,
		Y:  42,
	}, result)
}
