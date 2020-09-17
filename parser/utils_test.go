package parser

import (
	"fmt"
	"testing"
)

var testSyml = map[string]int{
	"(asd()fsdf)": 10,
	"((())":       0,
	"())))":       1,
}

func TestMatch(t *testing.T) {
	fmt.Println(matchingNextSymbol('[', "]]]]"))
}


func TestMachAll(t *testing.T) {
	for s, i := range testSyml {
		if index := matchingNextSymbol('(', s); index != i {
			t.Fatalf("%s index result %d expect %d", s, index, i)
		}
	}
}


