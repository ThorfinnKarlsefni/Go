package stack

type Stack1 struct {
	data []interface{}
}

func (s *Stack1) Push(n []interface{}) {
	s.data = append(s.data,n)
}

func (s *Stack1) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

func (s *Stack1) Size() int {
	return len(s.data)
}