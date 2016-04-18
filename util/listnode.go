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

func ListEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		// fmt.Println(l1.Val, l2.Val)
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
