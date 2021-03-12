package jparser

import (
	"fmt"

	"gitee.com/aifuturewell/gojni/jni"
)

type Compiler struct {
	env   jni.Env
	check bool
}

const (
	STATICOJB int = iota
	NEWOBJECT
	OBJECTINVOKE
)

func (Define) exprNode() {}
func (c *Define) String() string {
	return fmt.Sprintf("(def {%s}  {%s})", c.Name, c.Value)
}

func (Assignment) exprNode() {}
func (a *Assignment) String() string {
	return fmt.Sprintf("(From %s ==> %s )", a.From, a.To)
}

func (Call) exprNode() {}
func (c *Call) String() string {
	return fmt.Sprintf("(CALL Class: %s type %d Method: %s Args %s)", c.Owner, c.ClassTyp, c.Method, c.Args)
}

func (Class) exprNode() {}
func (c *Class) String() string {
	return fmt.Sprintf("(Owner: %s )", c.Name)
}

func (a Arg) exprNode() {}
func (a Arg) String() string {
	return fmt.Sprintf("%s%d", []string{"$", "@"}[a.Typ], a.ArgN)
}
