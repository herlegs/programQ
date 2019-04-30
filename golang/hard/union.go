package hard

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	unionMap := make(map[string]*Parent)
	for i, e := range equations {
		union(unionMap, e[0], e[1], values[i])
	}
	result := make([]float64, len(queries))
	for i, query := range queries {
		_, leftExist := unionMap[query[0]]
		_, rightExist := unionMap[query[1]]
		if !leftExist || !rightExist {
			result[i] = -1
			continue
		}
		aRoot, bRoot := find(unionMap, query[0]), find(unionMap, query[1])
		if aRoot.key == bRoot.key {
			result[i] = aRoot.ratio / bRoot.ratio
		} else {
			result[i] = -1
		}
	}
	return result
}

func union(m map[string]*Parent, a, b string, ratio float64) {
	aRoot, bRoot := find(m, a), find(m, b)
	if aRoot.key != bRoot.key {
		root := m[aRoot.key]
		root.key = bRoot.key
		root.ratio = bRoot.ratio * ratio / aRoot.ratio
	}
}

func find(m map[string]*Parent, a string) *Parent {
	if _, exist := m[a]; !exist {
		m[a] = &Parent{a, 1}
		return m[a]
	}
	if m[a].key != a {
		root := find(m, m[a].key)
		m[a].key = root.key
		m[a].ratio = m[a].ratio * root.ratio
	}
	return m[a]
}

type Parent struct {
	key   string
	ratio float64
}
