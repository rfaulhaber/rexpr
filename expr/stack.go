package expr

type stack struct {
	arr []string
}

func (s *stack) Push(str string) {
	s.arr = append(s.arr, str)
}

func (s *stack) Pop() string {
	if len(s.arr) > 0 {
		lastIndex := len(s.arr) - 1

		ret := s.arr[lastIndex]
		newArr := s.arr[:lastIndex]

		s.arr = newArr

		return ret
	}

	return ""
}

func (s *stack) Empty() bool {
	return len(s.arr) == 0
}

func (s *stack) Size() int {
	return len(s.arr)
}
