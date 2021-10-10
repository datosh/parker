package parker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberSquare_SquareValues(t *testing.T) {

	t.Run("square of 1", func(t *testing.T) {
		square := NumberSquare{
			[9]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		square.SquareValues()

		for _, value := range square.values {
			if value != 1 {
				assert.Equal(t, 1, value)
			}
		}
	})

	t.Run("square incrementing", func(t *testing.T) {
		square := NumberSquare{
			[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
		square.SquareValues()

		assert.Equal(t, square.values[0], 1)
		assert.Equal(t, square.values[1], 4)
		assert.Equal(t, square.values[2], 9)

		assert.Equal(t, square.values[3], 16)
		assert.Equal(t, square.values[4], 25)
		assert.Equal(t, square.values[5], 36)

		assert.Equal(t, square.values[6], 49)
		assert.Equal(t, square.values[7], 64)
		assert.Equal(t, square.values[8], 81)
	})
}

func TestNumberSquare_rowSum(t *testing.T) {
	square := NumberSquare{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	assert.Equal(t, 6, square.rowSum(0))
	assert.Equal(t, 15, square.rowSum(1))
	assert.Equal(t, 24, square.rowSum(2))
}

func TestNumberSquare_colSum(t *testing.T) {
	square := NumberSquare{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	assert.Equal(t, 12, square.colSum(0))
	assert.Equal(t, 15, square.colSum(1))
	assert.Equal(t, 18, square.colSum(2))
}

func TestNumberSquare_diagSum1(t *testing.T) {
	square := NumberSquare{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	assert.Equal(t, 15, square.diagSum1())
}

func TestNumberSquare_diagSum2(t *testing.T) {
	square := NumberSquare{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	assert.Equal(t, 15, square.diagSum2())
}

func TestNumberSquare_SameValues(t *testing.T) {
	t.Run("trivial", func(t *testing.T) {
		square := NumberSquare{
			[9]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		square.SquareValues()
		assert.Equal(t, 8, square.SameValues())
	})

	t.Run("trivial", func(t *testing.T) {
		square := NumberSquare{
			[9]int{29, 1, 47, 41, 37, 1, 23, 41, 29},
		}
		square.SquareValues()
		assert.Equal(t, 7, square.SameValues())
	})
}
