package graph

// initSingleSource initializes the shortest-path
// estimates for all vertices in graph G, and sets
// the estimate for vertex labeled s to 0
// The estimate is an upper bound on the weight of a shortest
// path from source s to v - which would be set to the largest
// int value depending on the target architecture
// after initialization - v.Distance = max for all v ∈ V - {src}
// This is achieved in an order 0(V)-time, where V is the number of Vertices
func (G *Graph) initSingleSource(src Label) {
	for _, v := range G.V {
		v.Distance = int(^uint(1) >> 1)
	}
	G.V[src].Distance = 0
}

// relaxing an edge (u, v) is done based on a condition
// that tells if we can improve the shortest path to v
// found so far by going through u and, if the condition
// passes, update v.Distance and v.Predecessor
// This is achieved in an order 0(1)-time
func (G *Graph) relax(e Edge) {
	w := G.E[e]
	improved := e.V.Distance > e.U.Distance+w
	if improved {
		e.V.Distance = e.U.Distance + w
		e.V.Predecessor = e.U
	}
}

// vertices satisfy heap.Heaper interface for binary priority queue
type vertices []*Vertex

func (v *vertices) Pop() interface{} {
	t := *v
	u := t[0]
	*v = t[1:]
	return u
}

func (v *vertices) Push(x interface{}) {
	*v = append(*v, x.(*Vertex))
}

func (v vertices) Len() int {
	return len(v)
}

func (v vertices) Less(i, j int) bool {
	return v[i].Distance < v[j].Distance
}

func (v *vertices) Swap(i, j int) {
	t := *v
	t[i], t[j] = t[j], t[i]
}

func (v vertices) Get(i int) interface{} {
	return v[i]
}

func (v *vertices) Set(i int, x interface{}) {
	//
}
