package graph

import (
	"fmt"
)

type VertexID int

func (v VertexID) Int() int {
	return int(v)
}

type Edge struct {
	From   VertexID
	To     VertexID
	Weight float64
}

type DirGraph struct {
	edges map[VertexID]map[VertexID]float64
	e     int // 边总数
}

func NewDirectedGraph() *DirGraph {
	return &DirGraph{
		edges: make(map[VertexID]map[VertexID]float64),
		e:     0,
	}
}

func (g *DirGraph) E() int {
	return g.e
}

func (g *DirGraph) V() int {
	return len(g.edges)
}

func (g *DirGraph) CheckVertex(vertex VertexID) bool {
	_, ok := g.edges[vertex]
	return ok
}

func (g *DirGraph) CheckEdge(from, to VertexID) bool {
	_, ok := g.edges[from][to]
	return ok
}

func (g *DirGraph) AddVertex(vertex VertexID) error {
	_, ok := g.edges[vertex]
	if ok {
		return fmt.Errorf("vertex already exist")
	}

	g.edges[vertex] = make(map[VertexID]float64)
	return nil
}

func (g *DirGraph) RemoveVertex(vertex VertexID) error {
	_, ok := g.edges[vertex]
	if !ok {
		return fmt.Errorf("unknown vertex")
	}

	delete(g.edges, vertex)

	for _, connextedVertices := range g.edges {
		delete(connextedVertices, vertex)
	}

	return nil
}

func (g *DirGraph) AddEdge(from, to VertexID, weight float64) error {
	if from == to {
		return fmt.Errorf("cannot add self loop")
	}

	if !g.CheckVertex(from) || !g.CheckVertex(to) {
		return fmt.Errorf("vertices don`t exist")
	}

	_, ok := g.edges[from][to]
	if ok {
		return fmt.Errorf("edge already defined")
	}

	if _, ok := g.edges[from]; !ok {
		g.edges[from] = make(map[VertexID]float64)
	}
	g.edges[from][to] = weight
	g.e++

	return nil
}

func (g *DirGraph) RemoveEdge(from, to VertexID) error {
	_, ok := g.edges[from][to]
	if !ok {
		return fmt.Errorf("edge doesn`t exist")
	}

	delete(g.edges[from], to)
	g.e--

	return nil
}

func (g *DirGraph) EdgesIter() <-chan *Edge {
	ch := make(chan *Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, w := range connectedVertices {
				ch <- &Edge{from, to, w}
			}
		}

		close(ch)
	}()

	return ch
}

func (g *DirGraph) VerticesIter() <-chan VertexID {
	ch := make(chan VertexID)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}

		close(ch)
	}()
	return ch
}

func (g *DirGraph) Edges(v VertexID) <-chan *Edge {
	ch := make(chan *Edge)
	go func() {
		if _, ok := g.edges[v]; !ok {
			ch <- nil
			return
		}

		for vertex, w := range g.edges[v] {
			ch <- &Edge{
				From:   v,
				To:     vertex,
				Weight: w,
			}
		}
		close(ch)
	}()

	return ch
}
