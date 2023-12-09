package shortestpath

import "zpy/algorithm/shortest_path/graph"

type DijistraPairs struct {
	all []Dijistra
}

func (d *DijistraPairs) Init(g *graph.DirGraph) {
	d.all = make([]Dijistra, g.V())
	for v := 0; v < g.V(); v++ {
		d.all[v].ShortestPath(g, graph.VertexID(v))
	}
}

func (d *DijistraPairs) Path(s, t graph.VertexID) <-chan graph.Edge {
	return d.all[s].PathTo(t)
}

func (d *DijistraPairs) Dist(s, t graph.VertexID) float64 {
	return d.all[s].DistTo(t)
}
