package priorityqueue

import (
	"fmt"
	"zpy/algorithm/shortest_path/queue"
)

type Item struct {
	Value    interface{}
	Priority float64
}

func (i Item) Less(than *Item) bool {
	return i.Priority < than.Priority
}

func (i Item) Greater(than *Item) bool {
	return i.Priority > than.Priority
}

type PQ struct {
	data []*Item
	n    int
}

// func New() *PQ {
// 	return &PQ{
// 		data: make([]Item, 1),
// 		n:    0,
// 	}
// }

func NewWithSize(size int) *PQ {
	return &PQ{
		data: make([]*Item, size+1),
		n:    0,
	}
}

func (p *PQ) IsEmpty() bool {
	return p.n == 0
}

func (p *PQ) Len() int {
	return p.n
}

func (p *PQ) Contains(v *Item) bool {
	for _, e := range p.data {
		if e == nil {
			continue
		}
		if e.Value == v.Value {
			return true
		}
	}

	return false
}

func (p *PQ) ChangePriority(v *Item) error {
	var storage = queue.New()

	popped := p.Pop()

	for v.Value != popped.Value {
		if p.n == 0 {
			return fmt.Errorf("priority queue doesn`t contain item: %v", *v)
		}

		storage.Push(popped)
		popped = p.Pop()
	}

	popped.Priority = v.Priority
	p.Insert(popped)

	for storage.Len() > 0 {
		p.Insert(storage.Pop().(*Item))
	}

	return nil
}

func (p *PQ) Insert(i *Item) {
	p.n++
	// if p.n == len(p.data) {
	// 	p.data = append(p.data, i)
	// } else {
	p.data[p.n] = i
	// }
	p.swim(p.n)
}

func (p *PQ) Pop() *Item {
	min := p.data[1]
	p.data[1], p.data[p.n] = p.data[p.n], p.data[1]
	p.n--
	p.sink(1)

	// if p.n < len(p.data)/4 {
	// 	p.data = p.data[:len(p.data)/2]
	// }

	return min
}

func (p *PQ) swim(i int) {
	for i > 1 && p.data[i/2].Greater(p.data[i]) {
		j := i / 2
		p.data[i], p.data[j] = p.data[j], p.data[i]
		i = j
	}
}

func (p *PQ) sink(i int) {
	for 2*i <= p.n {
		j := 2 * i
		if j < p.n && p.data[j].Greater(p.data[j+1]) {
			j++
		}
		if p.data[i].Less(p.data[j]) {
			break
		}

		p.data[i], p.data[j] = p.data[j], p.data[i]
		i = j
	}
}
