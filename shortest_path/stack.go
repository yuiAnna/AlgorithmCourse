package shortestpath

import "zpy/algorithm/shortest_path/graph"

type node struct {
	data *graph.Edge
	next *node
}

type Stack struct {
	head *node
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Push(v *graph.Edge) {
	new := &node{
		data: v,
		next: s.head,
	}
	s.head = new
}

func (s *Stack) Pop() *graph.Edge {
	if s.head == nil {
		return nil
	}

	result := s.head.data
	s.head = s.head.next
	return result
}
