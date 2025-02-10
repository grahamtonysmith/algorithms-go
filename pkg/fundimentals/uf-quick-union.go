package fundimentals

type UFQuickUnion struct {
	id    []int // is now a parent link representation of a forest of trees
	count int
}

func (uf *UFQuickUnion) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}

	return p
}

func (uf *UFQuickUnion) Union(p int, q int) {
	// give p and q the same root

	i := uf.Find(p)
	j := uf.Find(q)

	if i == j {
		return
	}

	uf.id[i] = j
	uf.count -= 1
}

func (uf *UFQuickUnion) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UFQuickUnion) Count() int {
	return uf.count
}

func NewUFQuickUnion(n int) *UFQuickUnion {
	id := make([]int, n)

	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &UFQuickUnion{
		id:    id,
		count: n,
	}
}
