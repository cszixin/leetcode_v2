package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(array []int, cycle bool) *ListNode {
	dummpy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	q, last := dummpy, dummpy

	for _, value := range array {
		node := new(ListNode)
		node.Val = value
		q.Next = node
		q = q.Next
		last = node
	}
	if cycle {
		last.Next = dummpy.Next
	}
	return dummpy.Next
}

func PrintLinkList(head *ListNode) []int {
	p := head
	out := make([]int, 0)
	for p != nil {
		out = append(out, p.Val)
		p = p.Next
	}
	return out
}

// 快慢指针的思想,获取一个链表的中间节点
func MiddleLinkList(head *ListNode) int {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow.Val
}

func HasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
