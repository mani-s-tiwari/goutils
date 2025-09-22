package dsu

type DSU struct {
	parent []int
	rank   []int
}

func New(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent: parent, rank: rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rx, ry := d.Find(x), d.Find(y)
	if rx == ry {
		return false
	}
	if d.rank[rx] < d.rank[ry] {
		d.parent[rx] = ry
	} else if d.rank[rx] > d.rank[ry] {
		d.parent[ry] = rx
	} else {
		d.parent[ry] = rx
		d.rank[rx]++
	}
	return true
}
