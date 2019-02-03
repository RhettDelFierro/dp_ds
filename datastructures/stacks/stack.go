package ds

// see https://blog.golang.org/slices
type Stack []string

func (q *Stack) Push(record []string) {
	*q = append(*q, record...)
}

// see => Pointers to slices: Method receivers
func (q *Stack) Pop() string {
	s := *q
	last := s[len(s)-1]
	*q = s[:len(s)-1]
	return last
}

func (q *Stack) Peek() string {
	s := *q
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}
