package unionfind

type UF struct {
	parent []int
	sz     []int
}

// 初始化union find
func getUF(n int) *UF {
	parent := make([]int, n)
	sz := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i // 初始化父节点指向自己，即根节点指向自己
		sz[i] = i     // 树高初始化为1
	}

	return &UF{parent, sz}
}

// 关联
func (u *UF) union(i int, j int) {
	root_i := u.find(i)
	root_j := u.find(j)

	if root_i == root_j {
		return
	}

	// 小树挂到大树上,秩优化
	if u.sz[root_i] > u.sz[root_j] {
		u.parent[root_j] = root_i
	} else {
		u.parent[root_i] = root_j

		if u.sz[root_i] == u.sz[root_j] {
			u.sz[root_j] += 1
		}
	}
}

// 查找根节点
func (u *UF) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i]) // 递归,路径压缩
	}

	return u.parent[i]
}

// 判断是否连通
func (u *UF) connected(i int, j int) bool {
	root_i := u.find(i)
	root_j := u.find(j)

	return root_i == root_j
}
