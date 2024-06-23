package dynamicarray_test

import (
	dynamicarray "neetcode_roadmap/internal/arrays_hashing/prerequisites/dynamic_array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	assert := assert.New(t)

	t.Run(`should create an Array with default lenght`, func(t *testing.T) {
		da := dynamicarray.New[int]()
		assert.NotNil(da)
		assert.Equal(da.Length(), 10)
	})

	t.Run(`should add a new value to the array`, func(t *testing.T) {
		da := dynamicarray.New[string]()
		da.Add("hello")
		assert.NotNil(da)
		assert.Equal(da.Length(), 10)
		assert.Equal(da.Items(), []string{"hello"})
	})
}
