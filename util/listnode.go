package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListFromSlice(arr []int) *ListNode {
	dummy := &ListNode{1, nil}
	var cur = dummy

	for _, c := range arr {
		cur.Next = &ListNode{c, nil}
		cur = cur.Next
	}
	return dummy.Next
}

func DisplayList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
