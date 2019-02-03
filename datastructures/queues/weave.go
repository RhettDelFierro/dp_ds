package ds

func weave(sourceOne, sourceTwo *Queue) Queue {
	q := Queue{}

	for len(sourceOne.Peek()) > 0 || len(sourceTwo.Peek()) > 0 {
		if len(sourceOne.Peek()) > 0 {
			q.Add([]string{sourceOne.Remove()})
		}
		if len(sourceTwo.Peek()) > 0 {
			q.Add([]string{sourceTwo.Remove()})
		}
	}

	return q
}
