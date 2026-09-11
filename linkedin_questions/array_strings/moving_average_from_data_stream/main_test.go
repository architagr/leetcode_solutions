package movingaveragefromdatastream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	// The worked example from the problem statement. The window is 3,
	// so the fourth call drops the 1: (10 + 3 + 5) / 3 = 6.
	a := Constructor(3)
	assert.InDelta(t, 1.0, a.Next(1), 1e-5)
	assert.InDelta(t, 5.5, a.Next(10), 1e-5)
	assert.InDelta(t, 14.0/3.0, a.Next(3), 1e-5)
	assert.InDelta(t, 6.0, a.Next(5), 1e-5)
}
