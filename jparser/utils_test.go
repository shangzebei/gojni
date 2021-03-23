package jparser

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/utils"
	"testing"
)

var testSyml = map[string]int{
	"(asd()fsdf)": 10,
	"((())":       0,
	"())))":       1,
}

func TestMatch(t *testing.T) {
	fmt.Println(utils.MatchingNextSymbol('[', "]]]]"))
}

func TestMachAll(t *testing.T) {
	for s, i := range testSyml {
		if index := utils.MatchingNextSymbol('(', s); index != i {
			t.Fatalf("%s index result %d expect %d", s, index, i)
		}
	}
}
