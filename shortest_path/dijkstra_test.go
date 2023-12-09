package shortestpath

import (
	"fmt"
	"os"
	"testing"
	"zpy/algorithm/shortest_path/graph"
)

func TestDijkstra(t *testing.T) {
	h := graph.NewDirectedGraph()

	for i := 0; i < 5; i++ {
		h.AddVertex(graph.VertexID(i))
	}

	h.AddEdge(graph.VertexID(0), graph.VertexID(1), 10)
	h.AddEdge(graph.VertexID(1), graph.VertexID(2), 20)
	h.AddEdge(graph.VertexID(2), graph.VertexID(3), 40)
	h.AddEdge(graph.VertexID(0), graph.VertexID(2), 50)
	h.AddEdge(graph.VertexID(0), graph.VertexID(3), 80)
	h.AddEdge(graph.VertexID(0), graph.VertexID(4), 10)
	h.AddEdge(graph.VertexID(4), graph.VertexID(3), 10)

	var dijistra Dijistra
	dijistra.ShortestPath(h, graph.VertexID(0))
	for i := 1; i < 5; i++ {
		for j := range dijistra.PathTo(graph.VertexID(i)) {
			fmt.Fprintf(os.Stdout, "%v->%v: %.2f, ", j.From, j.To, j.Weight)
		}
		fmt.Fprint(os.Stdout, "\n")
	}
}
