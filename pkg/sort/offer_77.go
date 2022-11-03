package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList() *ListNode {

	arr := []int{4, 1, 5, 6}
	dummpy := &ListNode{}
	p := dummpy
	for _, v := range arr {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	return dummpy.Next

}
func getMid(head, tail *ListNode) *ListNode {
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return slow
}

func sortList(head *ListNode) *ListNode {
	var mysort func(*ListNode, *ListNode) *ListNode
	merge := func(left *ListNode, right *ListNode) *ListNode {
		dummpy := &ListNode{}
		p := dummpy
		l := left
		r := right
		for l != nil && r != nil {
			if l.Val <= r.Val {
				p.Next = l
				p = p.Next
				l = l.Next
			} else {
				p.Next = r
				p = p.Next
				r = r.Next
			}

		}

		if l != nil {
			p.Next = l
		}
		if r != nil {
			p.Next = r
		}
		return dummpy.Next
	}
	mysort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			//空链表
			return head
		}
		if head.Next == tail {
			//链表长度1
			head.Next = nil
			return head

		}
		mid := getMid(head, tail)

		return merge(mysort(head, mid), mysort(mid, tail))
	}
	return mysort(head, nil)

}
