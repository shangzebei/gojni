package utils

import "log"

//#####################################################################################
type SybHeap []uint8

var syb = map[byte]byte{
	40:  41,
	91:  93,
	123: 125,
}

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

func MatchingNextSymbol(s byte, str string) int {
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
