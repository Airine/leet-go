package unionfind

import (
	"sort"

	"github.com/airine/leet-go/utils/structures"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	ufs := structures.NewUFSet(n)

	for _, p := range pairs {
		ufs.Union(p[0],p[1])
	}

	groups := map[int][]byte{}
	for i := 0; i < n; i++ {
		p := ufs.Find(i)
		groups[p] = append(groups[p], s[i])
	}

	for _, bytes := range groups {
		sort.Slice(bytes, func(i, j int) bool {return bytes[i] < bytes[j]})
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		p := ufs.Find(i)
		res[i] = groups[p][0]
		groups[p] = groups[p][1:]
	}

	return string(res)
}