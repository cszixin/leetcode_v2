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

func IfIntersection(head1, head2 *ListNode) (int, bool) {
	np := make(map[*ListNode]struct{})
	p1, p2 := head1, head2
	for p1 != nil {
		np[p1] = struct{}{}
		p1 = p1.Next
	}
	for p2 != nil {
		if _, ok := np[p2]; ok {
			return p2.Val, true
		}
		p2 = p2.Next
	}
	return -1, false
}

// 创建两个相交的链表
func CreateIntersection() (*ListNode, *ListNode) {
	array1 := []int{3, 6, 3, 1}
	array2 := []int{6, 7, 8, 10}
	common := []int{12, 13, 14}
	h1 := CreateLinkList(array1, false)
	h2 := CreateLinkList(array2, false)
	hcommon := CreateLinkList(common, false)
	p1, p2 := h1, h2
	for p1.Next != nil {
		p1 = p1.Next
	}
	for p2.Next != nil {
		p2 = p2.Next
	}
	p1.Next = hcommon
	p2.Next = hcommon
	return h1, h2
}

// 反转链表

func ReverseV1(head *ListNode) *ListNode {
	dummy := new(ListNode)
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = dummy.Next
		dummy.Next = p
		p = tmp
	}
	return dummy.Next
}

func GetNode(head *ListNode, k int) *ListNode {
	p := head
	i := 1
	for i <= k-1 && p != nil {
		p = p.Next
		i++
	}
	return p
}

// 反转链表 递归

func ReverseV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *ListNode

// 反转链表中前N个元素
func ReverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := ReverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func Reversev3(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur, next = head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func ReverseBetween(a, b *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur, next = a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 每k个节点一反转
func ReverKGroup(head *ListNode, k int) *ListNode {
	var start, end *ListNode
	start, end = head, head
	for i := 0; i < k; i++ {
		// 向前k步
		if end == nil {
			//不足k步
			return head
		}
		end = end.Next
	}
	//反转前k个
	new_head := ReverseBetween(start, end)
	//前面已经将前k个节点反转,现在递归的处理:以第k+1个节点为头结点的后续链表
	start.Next = ReverKGroup(end, k)
	return new_head
}
