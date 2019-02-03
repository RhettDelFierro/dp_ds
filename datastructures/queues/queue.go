package ds

// see https://blog.golang.org/slices
type Queue []string

func (q *Queue) Add(record []string) {
	*q = append(record, *q...)
}

// see => Pointers to slices: Method receivers
func (q *Queue) Remove() string {
	s := *q
	last := s[len(s)-1]
	*q = s[:len(s)-1]
	return last
}

func (q *Queue) Peek() string {
	s := *q
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}
