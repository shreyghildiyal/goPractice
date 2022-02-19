package practice_test

import (
	"fmt"
	"testing"

	practice "github.com/shreyghildiyal/goPractice/problems"
)

func TestNextPermutation(t *testing.T) {
	vals := [][][]int{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 4, 3},
		},
		{
			[]int{1, 2, 4, 3},
			[]int{1, 3, 2, 4},
		},
	}

	for _, inOut := range vals {
		arr := inOut[0]
		expectedOut := inOut[1]
		practice.NextPermutation(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] != expectedOut[i] {
				fmt.Println("Expected", expectedOut, " Found", arr)
				t.FailNow()
			}
		}
	}
}
