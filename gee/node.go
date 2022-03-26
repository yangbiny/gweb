package gee

type node struct {
	pattern  string  // 当前的完整路径
	part     string  // 当前的部分路径
	children []*node // 子路径信息
	isWild   bool    // 是否严格匹配
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return n
		}
	}
	return nil
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}
