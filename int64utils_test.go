package int64utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	assert := assert.New(t)

	assert.Equal([]int64{1}, Difference([]int64{1, 2}, []int64{2}))
	assert.Equal([]int64{1}, Difference([]int64{1, 2}, []int64{2, 9}))
	assert.Equal([]int64{1, 2}, Difference([]int64{1, 2}, []int64{3}))
}
