package mathutil

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceTo(t *testing.T) {
	cases := []struct {
		From Vector2D
		To   Vector2D
		Dist float64
	}{
		{
			Vector2D{X: 1, Y: 3},
			Vector2D{X: 5, Y: 2},
			math.Sqrt(17),
		},
		{
			Vector2D{X: 0, Y: 0},
			Vector2D{X: 0, Y: 0},
			0,
		},
		{
			Vector2D{X: 0, Y: 0},
			Vector2D{X: 4, Y: 8},
			math.Sqrt(80),
		},
	}

	for _, c := range cases {
		dist := c.From.DistanceTo(c.To)
		assert.Equal(t, dist, c.Dist)
	}
}
