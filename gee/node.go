package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 当前的完整路径
	part     string  // 当前的部分路径
	children []*node // 子路径信息
	isWild   bool    // 是否模糊匹配
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
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

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") || strings.HasPrefix(n.part, ":") {
		if n.pattern == "" {
			return n
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		search := child.search(parts, height+1)
		if search != nil {
			return search
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	var hasAppend = false
	var tempNode []*node
	for _, child := range n.children {
		if child.part == part {
			nodes = append(nodes, child)
			hasAppend = true
		}
		if child.isWild {
			tempNode = append(tempNode, child)
		}
	}
	if !hasAppend && len(tempNode) > 0 {
		for _, child := range tempNode {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
