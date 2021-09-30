package queue

type Queue struct {
	data []interface{}
	size int
}

func (q *Queue) Size() int {
	return (*q).size
}

func (q *Queue) IsEmpty() bool {
	return (*q).size == 0
}

func (q *Queue) Enqueue(el interface{}) {
	(*q).data = append((*q).data, el)
	(*q).size++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}
	result := (*q).data[0]
	(*q).data = (*q).data[1:]
	(*q).size--
	return result
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty.")
	}
	return (*q).data[0]
}
