package jparser

import (
	"fmt"
	"testing"
)

var testSets = []string{
	"java.lang.Thread t = java.lang.Thread.currentThread[java.lang.Thread()]();@1=t.getName[void()]();",
	"@1=new java.util.Date[void()]().toString[void()]();",
	"java.lang.Thread t = java.lang.Thread.currentThread[java.lang.Thread()]();@1=t.getName[java.lang.String(void)]();",
	"@1=t.getName[void()]();@2=t.getSex[void(int)]();",
	"@1=com.szb.Miner.show[void(int)]($1,$2);", //param
	"@2=$1.getName[java.lang.String()]();@3=com.szb.Miner.show[void(java.lang.String)](@2);",
	"@3=$1.show[void(java.lang.String)]($2.getName[java.lang.String()]());",
	"$2.getName[java.lang.String()]();",
	"new com.szb.Hello[void()]().show[void()]()",
	"com.szb.Hello.send[void()]()",
	"com.szb.Hello.send[void()](@2)",
}

func TestParse(t *testing.T) {
	v := Compiler{check: false}
	fmt.Println(Print(v.Parse(testSets[10])))
}

func TestRun(t *testing.T) {
	v := Compiler{}
	for i, set := range testSets {
		fmt.Printf("[%d] %s\n", i, set)
		fmt.Println(Print(v.Parse(set)))
	}
}
