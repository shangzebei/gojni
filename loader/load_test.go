package loader

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCode(t *testing.T) {
	fmt.Println(int('a'))
	var kk []string
	fmt.Println(reflect.TypeOf(kk))
}
