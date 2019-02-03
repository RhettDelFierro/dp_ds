package ds

import "fmt"

// see https://blog.golang.org/slices
type Queue []string

func (q *Queue) Add(record []string) {
	*q = append(record, *q...)
	fmt.Println(*q)
}

func (q *Queue) Remove() {
	s := *q
	*q = s[:len(s)-1]
}
