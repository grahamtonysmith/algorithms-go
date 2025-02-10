package fundimentals

type UFQuickUnionWeighted struct {
	id    []int
	sz    []int
	count int
}

func (uf *UFQuickUnionWeighted) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}

	return p
}

func (uf *UFQuickUnionWeighted) Union(p int, q int) {
	// make the smaller root point to larger one

	i := uf.Find(p)
	j := uf.Find(q)

	if i == j {
		return
	}

	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}

	uf.count -= 1
}

func (uf *UFQuickUnionWeighted) Connected(p int, q int) bool {
	return uf.id[p] == uf.id[q]
}

func (uf *UFQuickUnionWeighted) Count() int {
	return uf.count
}

func NewUFQuickUnionWeighted(n int) *UFQuickUnionWeighted {
	id := make([]int, n)
	sz := make([]int, n)

	for i := 0; i < n; i++ {
		id[i] = i
	}

	for i := 0; i < n; i++ {
		sz[i] = 1
	}

	return &UFQuickUnionWeighted{
		id:    id,
		sz:    sz,
		count: n,
	}
}
