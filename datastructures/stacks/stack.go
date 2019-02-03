package ds

import "fmt"

// see https://blog.golang.org/slices
type Stack []string

func (q *Stack) Add(record []string) {
	*q = append(*q, record...)
	fmt.Println(*q)
}

// see => Pointers to slices: Method receivers
func (q *Stack) Remove() {
	s := *q
	*q = s[:len(s)-1]
}
