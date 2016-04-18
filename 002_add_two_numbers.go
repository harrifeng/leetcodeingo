package leetcodeingo

import "github.com/harrifeng/leetcodeingo/util"

func addTwoNumbers(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{-1, nil}
	cur := dummy

	adv := 0
	for l1 != nil || l2 != nil || adv > 0 {
		cnt := adv
		if l1 != nil {
			cnt += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cnt += l2.Val
			l2 = l2.Next
		}
		cur.Next = &util.ListNode{cnt % 10, nil}
		cur = cur.Next
		adv = cnt / 10
	}
	return dummy.Next

}
