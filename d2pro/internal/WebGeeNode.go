package internal

import (
	"fmt"
	"strings"
)

type WebGeeNode struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []INode //子节点
	isWild   bool    //是否精确匹配
}

func (w *WebGeeNode) SetPattern(pattern string) {
	w.pattern = pattern
}

func (w *WebGeeNode) GetPart() string {
	return w.part
}

func (w *WebGeeNode) SetPart(part string) {
	w.part = part
}

func (w *WebGeeNode) GetChildren() []INode {
	return w.children
}

func (w *WebGeeNode) SetChildren(children []INode) {
	w.children = children
}

func (w *WebGeeNode) GetIsWild() bool {
	return w.isWild
}

func (w *WebGeeNode) SetIsWild(isWild bool) {
	w.isWild = isWild
}

func (w *WebGeeNode) GetPattern() string {
	return w.pattern
}
func NewWebGeeNode(part string, isWild bool) *WebGeeNode {
	return &WebGeeNode{part: part, isWild: isWild}
}

// 使用递归插入
// 通过height定位当前处理的路径段（parts[height]），若无匹配子节点则创建新节点（支持通配符:和*标记）
func (w *WebGeeNode) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		w.pattern = pattern
		return
	}
	fmt.Printf("节点插入: pattern=%s parts=%v\n", pattern, parts)
	if w == nil {
		fmt.Println("节点为空")
	}
	part := parts[height]
	child := w.MatchChild(part)
	if child == nil {
		child = NewWebGeeNode(part, part[0] == ':' || part[0] == '*')
		w.SetChildren(append(w.GetChildren(), child))

	}
	child.Insert(pattern, parts, height+1)
}

// 搜索路径片段
// height是为了控制递归的深度
func (w *WebGeeNode) Search(parts []string, height int) INode {
	if len(parts) == height || strings.HasPrefix(w.part, "*") {
		if w.pattern == "" {
			return nil
		}
		return w
	}
	part := parts[height]
	children := w.MatchChildren(part)
	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 返回第一个匹配成功的节点
func (w *WebGeeNode) MatchChild(part string) INode {
	for _, child := range w.children {
		if child.GetPart() == part || child.GetIsWild() {
			return child
		}
	}
	return nil
}

// 返回所有匹配成功的节点
func (w *WebGeeNode) MatchChildren(part string) []INode {
	nodes := make([]INode, 0)
	for _, child := range w.children {
		if child.GetPart() == part || child.GetIsWild() {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
