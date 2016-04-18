package main

import (
	"os"

	"github.com/harrifeng/leetcodeingo/util"
)

func main() {
	a := []int{1, 2, 3}

	head := util.CreateListFromSlice(a)
	util.DisplayList(head)
	os.Exit(0)
}
