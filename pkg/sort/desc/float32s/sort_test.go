package float32s

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	Num  = 10
	Frac = 100
)

func generate() ([]float32, []int64) {
	os := make([]int64, Num)
	xs := make([]float32, Num)
	{
		for i := 0; i < Num; i++ {
			os[i] = int64(i)
			xs[i] = float32(rand.Float64() * Frac)
		}
	}
	return xs, os
}

func TestSort(t *testing.T) {
	vs, os := generate()
	for i, o := range os {
		fmt.Printf("[%v] = %v\n", i, vs[o])
	}
	Sort(vs, os[2:])
	fmt.Printf("\n")
	for i, o := range os {
		fmt.Printf("[%v] = %v\n", i, vs[o])
	}
}
