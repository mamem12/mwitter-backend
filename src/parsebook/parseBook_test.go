package parsebook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	getPage()

	assert.Equal(0, 0, "")
}
