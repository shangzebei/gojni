package jparser

import (
	"fmt"
	"log"
)

var syb = map[byte]byte{
	40:  41,
	91:  93,
	123: 125,
}

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

//#####################################################################################
type SybHeap []uint8

func (h *SybHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *SybHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *SybHeap) Len() int {
	return len(*h)
}

func (h *SybHeap) Pop() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *SybHeap) Push(v interface{}) {
	*h = append(*h, v.(uint8))
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

func matchingNextSymbol(s byte, str string) int {
	if s != str[0] {
		log.Printf("%s not start with %s", str, string(s))
		return 0
	}
	var pool SybHeap
	if f, b := syb[s]; b {
		for i := 0; i < len(str); i++ {
			c := str[i]
			if c == s {
				pool.Push(s)
			}
			if c == f {
				pool.Pop()
			}
			if pool.Len() == 0 {
				return i
			}
		}
	}
	return 0
}

func Print(nodes []Expr) string {
	var s string
	for _, node := range nodes {
		s += fmt.Sprintf("\t%s\n", node)
	}
	return s
}
