package queue

type Queue struct {
	queue []interface{}
	len   int
}

func New() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
		len:   0,
	}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) isEmpty() bool {
	return q.len == 0
}

func (q *Queue) Pop() (result interface{}) {
	result, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(v interface{}) {
	q.queue = append(q.queue, v)
	q.len++
}

func (q *Queue) Peek() interface{} {
	return q.queue[0]
}
