package dynamicarray_test

import (
	dynamicarray "neetcode_roadmap/internal/arrays_hashing/prerequisites/dynamic_array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	assert := assert.New(t)

	t.Run(`should create an Array with default lenght`, func(t *testing.T) {
		da := dynamicarray.New()
		assert.NotNil(da)
		assert.Equal(da.Length(), 10)
	})
}
