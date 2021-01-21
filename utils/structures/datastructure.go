package structures

// Start of UnionFind

// UnionFindSet ...
type UnionFindSet struct {
	Parent []int
	Rank   []int
	Count  int
}

func NewUFSet(n int) UnionFindSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return UnionFindSet{Parent: parent, Rank: rank, Count: n}
}

func (ufs *UnionFindSet) Find(x int) int {
	if ufs.Parent[x] != x {
		ufs.Parent[x] = ufs.Find(ufs.Parent[x])
	}
	return ufs.Parent[x]
}

func (ufs *UnionFindSet) Union(a, b int) {
	pa, pb := ufs.Find(a), ufs.Find(b)
	if pa == pb {
		return
	}
	if ufs.Rank[pa] < ufs.Rank[pb] {
		pa, pb = pb, pa
	}
	ufs.Parent[pb] = pa
	ufs.Count--
	ufs.Rank[pa] += ufs.Rank[pb]
}

// End of UnionFindSet
