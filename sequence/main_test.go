package sequence

import "testing"
import "github.com/stretchr/testify/require"

func TestSequence(t *testing.T) {
	t.Run("case #1", func(t *testing.T) {
		case1 := []int{1, 2, 3, 8, 9, 3, 2, 1}
		require.Equal(t, 3, Max(case1))
	})

	t.Run("case #2", func(t *testing.T) {
		case2 := []int{1, 2, 1, 2, 2, 1}
		require.Equal(t, 2, Max(case2))
	})

	t.Run("case #3", func(t *testing.T) {
		case3 := []int{7, 1, 2, 3, 7, 2, 1}
		require.Equal(t, 2, Max(case3))
	})
}
