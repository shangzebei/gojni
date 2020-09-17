package parser

import (
	"fmt"
	"log"
)

type sybHeap []uint8

var syb = map[byte]byte{
	40:  41,
	91:  93,
	123: 125,
}

type nodeHeap []Expr

func (h *nodeHeap) Top() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[h.Len()-1]
}

func (h *nodeHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *nodeHeap) Push(v interface{}) {
	*h = append(*h, v.(Expr))
}

func (h *nodeHeap) Len() int {
	return len(*h)
}
func (h *nodeHeap) Clear() {
	*h = (*h)[0:0]
}

func (h *nodeHeap) Root() (v interface{}) {
	return (*h)[0]
}

func (h *sybHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *sybHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *sybHeap) Len() int {
	return len(*h)
}

func (h *sybHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *sybHeap) Push(v interface{}) {
	*h = append(*h, v.(uint8))
}

func matchingNextSymbol(s byte, str string) int {
	if s != str[0] {
		log.Printf("%s not start with %s", str, string(s))
		return 0
	}
	var pool sybHeap
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
