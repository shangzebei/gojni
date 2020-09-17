package parser

import (
	"fmt"
	"testing"
)

var testSets = []string{
	"java.lang.Thread t = java.lang.Thread.currentThread();@1=t.getName[void()]();",
	"@1=new java.util.Date().toString[void()]();",
	"java.lang.Thread t = java.lang.Thread.currentThread();@1=t.getName[java.lang.String(void)]();",
	"@1=t.getName[void()]();@2=t.getSex[void(int)]();",
	"@1=com.szb.Miner.show[void(int)]($1,$2);", //param
	"@2=$1.getName[java.lang.String()]();@3=com.szb.Miner.show[void(java.lang.String)](@2);",
	"@3=$1.show($2.getName[java.lang.String()]());",
	"$2.getName[java.lang.String()]();",
}

func TestParse(t *testing.T) {
	v := Vm{}
	fmt.Println(v.Parse(testSets[6]))
}

func TestRun(t *testing.T) {
	v := Vm{}
	for i, set := range testSets {
		fmt.Printf("[%d] %s\n", i, set)
		fmt.Println(Print(v.Parse(set)))
	}
}
