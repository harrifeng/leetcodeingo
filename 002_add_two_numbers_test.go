package leetcodeingo

import (
	"testing"

	"github.com/harrifeng/leetcodeingo/util"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := util.CreateListFromSlice([]int{2, 4, 3})
	l2 := util.CreateListFromSlice([]int{5, 6, 4})
	r1 := util.CreateListFromSlice([]int{7, 0, 8})

	// util.DisplayList(l1)
	// fmt.Println("---------")
	// util.DisplayList(l2)
	// fmt.Println("---------")
	// util.DisplayList(r1)
	// fmt.Println("---------")
	// util.DisplayList(addTwoNumbers(l1, l2))

	if !util.ListEqual(r1, addTwoNumbers(l1, l2)) {
		t.Errorf("Error occured")
	}
}
