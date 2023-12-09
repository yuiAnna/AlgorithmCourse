package shortestpath

import (
	"fmt"
	"math"
	"os"
	"zpy/algorithm/shortest_path/graph"
	pq "zpy/algorithm/shortest_path/priorityqueue"
)

type Dijistra struct {
	edgeTo []*graph.Edge
	distTo []float64
	pq     *pq.PQ
}

func (d *Dijistra) ShortestPath(g *graph.DirGraph, s graph.VertexID) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}
	}()

	d.edgeTo = make([]*graph.Edge, g.V())
	d.distTo = make([]float64, g.V())
	d.pq = pq.NewWithSize(g.V())

	INF := math.Inf(1)
	for i, _ := range d.distTo {
		d.distTo[i] = INF
	}
	d.distTo[s] = 0.0

	d.pq.Insert(&pq.Item{Value: s, Priority: 0})
	for !d.pq.IsEmpty() {
		d.relex(g, d.pq.Pop().Value.(graph.VertexID))
	}
}

func (d *Dijistra) relex(g *graph.DirGraph, v graph.VertexID) {
	for e := range g.Edges(v) {
		w := e.To
		if d.distTo[w] > d.distTo[v]+e.Weight {
			d.distTo[w] = d.distTo[v] + e.Weight
			d.edgeTo[w] = e

			if d.pq.Contains(&pq.Item{Value: w}) {
				err := d.pq.ChangePriority(&pq.Item{
					Value:    w,
					Priority: d.distTo[w],
				})
				if err != nil {
					panic(err)
				}
			} else {
				d.pq.Insert(&pq.Item{
					Value:    w,
					Priority: d.distTo[w],
				})
			}
		}
	}
}

func (d *Dijistra) DistTo(v graph.VertexID) float64 {
	return d.distTo[v]
}

func (d *Dijistra) HasPathTo(v graph.VertexID) bool {
	return d.distTo[v] < math.Inf(1)
}

func (d *Dijistra) PathTo(v graph.VertexID) <-chan graph.Edge {
	if !d.HasPathTo(v) {
		return nil
	}

	ch := make(chan graph.Edge)
	go func() {
		stack := Stack{}
		for e := d.edgeTo[v]; e != nil; e = d.edgeTo[e.From] {
			stack.Push(e)
		}
		for !stack.IsEmpty() {
			ch <- *stack.Pop()
		}

		close(ch)
	}()

	return ch
}
