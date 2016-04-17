package leetcodeingo

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a1 := []int{2, 7, 11, 15}
	t1 := 9
	r1 := []int{0, 1}

	if !reflect.DeepEqual(r1, twoSum(a1, t1)) {
		t.Errorf("Error occured")
	}
}
