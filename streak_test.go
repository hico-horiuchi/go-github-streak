package streak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContributions(t *testing.T) {
	assert := assert.New(t)

	contributions, err := GetContributions("hico-horiuchi")
	assert.Nil(err)
	assert.NotNil(contributions)
}
