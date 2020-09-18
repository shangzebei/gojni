package parser

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
)

type Vm struct {
	env   jni.Env
	check bool
}

const (
	STATICOJB int = iota
	NEWOBJECT
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
	return fmt.Sprintf("(CALL Class: %s Method: %s Args %s)", c.Owner, c.Method, c.Args)
}

func (Class) exprNode() {}
func (c *Class) String() string {
	return fmt.Sprintf("(Owner: %s )", c.Name)
}

func (a MethodSig) String() string {
	return fmt.Sprintf("sig %s", a.Sig)
}

func (a Arg) exprNode() {}
func (a Arg) String() string {
	return fmt.Sprintf("%s%d", []string{"$", "@"}[a.Typ], a.ArgN)
}
