package stack

type Stack struct {
	data []int
	top  int
}

func MakeStack() *Stack {
	s := new(Stack)
	s.data = make([]int, 0)
	s.top = -1
	return s
}

func (s *Stack) Push(value int) {
	s.top++
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, bool) {
	if s.top == -1 {
		// 空栈
		return 0, false
	}
	value := s.data[s.top]
	s.top--
	return value, true

}

func (s *Stack) Len() int {
	return s.top + 1
}

func (s *Stack) Top() (int, bool) {
	if s.top == -1 {
		// 空
		return 0, false
	}
	return s.data[s.top], true
}
