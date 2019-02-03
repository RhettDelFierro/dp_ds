package ds

// see https://blog.golang.org/slices
type Stack []string

func (q *Stack) Add(record []string) {
	*q = append(*q, record...)
}

// see => Pointers to slices: Method receivers
func (q *Stack) Remove() {
	s := *q
	*q = s[:len(s)-1]
}

func (q *Stack) Peek() string {
	s := *q
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}
