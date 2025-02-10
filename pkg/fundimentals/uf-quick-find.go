package fundimentals

type UFQuickFind struct {
	id    []int
	count int
}

func (uf *UFQuickFind) Find(p int) int {
	return uf.id[p]
}

func (uf *UFQuickFind) Union(p int, q int) {
	// put p and q in the same component

	pId := uf.Find(p)
	qId := uf.Find(q)

	if pId == qId {
		return
	}

	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pId {
			uf.id[i] = qId
		}
	}

	uf.count -= 1
}

func (uf *UFQuickFind) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UFQuickFind) Count() int {
	return uf.count
}

func NewUFQuickFind(n int) *UFQuickFind {
	id := make([]int, n)

	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &UFQuickFind{
		id:    id,
		count: n,
	}
}
