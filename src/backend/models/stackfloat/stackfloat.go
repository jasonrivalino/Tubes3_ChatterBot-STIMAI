package stackfloat

type StackFloat struct {
	items []float64
}

func (s *StackFloat) Push(item float64) {
	s.items = append(s.items, item)
}

func (s *StackFloat) Pop() float64 {
	if len(s.items) == 0 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *StackFloat) Peek() float64 {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *StackFloat) IsEmpty() bool {
	return len(s.items) == 0
}
