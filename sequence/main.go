package sequence

import (
	"math"
)

type Sequence struct {
	data []int
	i    int
}

func (r *Sequence) Next() {
	r.i++
}

func Max(arr []int) int {
	var sequence Sequence
	c := 1
	n := len(arr)
	for i := 0; i < int(math.Ceil(float64(n)/2)); i++ {
		if arr[i] == c {
			sequence.data = append(sequence.data, arr[i])
			c++
		}
	}

	max := 0
	for i := n - 1; i > int(math.Floor(float64(n)/2)); i-- {
		if sequence.i == len(sequence.data) {
			break
		}
		if arr[i] == sequence.data[sequence.i] {
			if arr[i] > max {
				max = arr[i]
			}
			sequence.Next()
		}
	}

	return max
}
