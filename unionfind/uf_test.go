package unionfind

import "testing"

func TestBase(t *testing.T) {
	l := [][]int{
		{0, 3},
		{1, 3},
		{2, 4},
		{1, 5},
	}

	uf := getUF(6)
	for _, v := range l {
		uf.union(v[0], v[1])
	}

	t.Log(uf.connected(0, 5))
	t.Log(uf.connected(2, 5))
	t.Log(uf.connected(3, 5))
}

// leetcode 990
func equationsPossible(equations []string) bool {
	uf := getUF(26)

	for _, v := range equations {
		if v[1] == '=' {
			uf.union(int(v[0]-'a'), int(v[3]-'a'))
		}
	}

	for _, v := range equations {
		if v[1] == '!' {
			if uf.connected(int(v[0]-'a'), int(v[3]-'a')) {
				return false
			}
		}
	}

	return true
}
