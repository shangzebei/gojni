package jparser

import (
	"fmt"
)

type NodeHeap []Expr

func (h *NodeHeap) Top() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[h.Len()-1]
}

func (h *NodeHeap) Pop() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *NodeHeap) Push(v interface{}) {
	*h = append(*h, v.(Expr))
}

func (h *NodeHeap) Len() int {
	return len(*h)
}
func (h *NodeHeap) Clear() {
	*h = (*h)[0:0]
}

func (h *NodeHeap) Root() (v interface{}) {
	return (*h)[0]
}

//####################################################################

type StrHeap []string

func (h *StrHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *StrHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *StrHeap) Len() int {
	return len(*h)
}

func (h *StrHeap) Pop() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *StrHeap) Push(v interface{}) {
	*h = append(*h, v.(string))
}

func Print(nodes []Expr) string {
	var s string
	for _, node := range nodes {
		s += fmt.Sprintf("\t%s\n", node)
	}
	return s
}
