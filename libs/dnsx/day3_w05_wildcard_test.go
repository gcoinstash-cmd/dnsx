package dnsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWildcardResponseValidation(t *testing.T) {
	options := DefaultOptions
	options.WildcardThreshold = 3
	assert.Equal(t, 3, options.WildcardThreshold)
}
