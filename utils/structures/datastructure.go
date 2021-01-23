package structures

// Start of UnionFind

// UnionFindSet ...
type UnionFindSet struct {
	Parent []int
	Count  int
}

func NewUFSet(n int) UnionFindSet {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return UnionFindSet{Parent: parent, Count: n}
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
	ufs.Parent[pb] = pa
	ufs.Count--
}

// End of UnionFindSet
