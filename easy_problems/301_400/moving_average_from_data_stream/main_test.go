package movingaveragefromdatastream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	a := Constructor(3)
	assert.Equal(t, float64(1.0), a.Next(1))
	assert.Equal(t, float64(5.5), a.Next(10))
	assert.Equal(t, float64(4.66667), a.Next(3))
	assert.Equal(t, float64(6.0), a.Next(6))
}
