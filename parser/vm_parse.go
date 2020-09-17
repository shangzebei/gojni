package parser

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"github.com/mohae/deepcopy"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Expr interface {
	exprNode()
}

type Class struct {
	Name string
}

type Define struct {
	Name  string
	Value Expr
}

type Assignment struct {
	From Expr
	To   Expr
}

type MethodSig struct {
	RetTyp   string
	ParamTyp []string
	Sig      string
}

type MethodMeta struct {
	Name string
	Sig  MethodSig
}

type Call struct {
	ClassTyp int
	Owner    Expr
	Method   MethodMeta
	Args     []Expr
}

type Arg struct {
	Typ  int
	ArgN int64
}

func (vm *Vm) checkSyntax(s string) bool {
	if !strings.HasSuffix(s, ";") {
		return false
	}
	return true
}

func (vm *Vm) Parse(s string) []Expr {
	if vm.check && !vm.checkSyntax(s) {
		return nil
	}
	s = strings.ReplaceAll(s, "\n", "")
	var tempChars []byte
	var spaceStrings []string
	var nodes []Expr
	var nh nodeHeap
	for i := 0; i < len(s); i += 1 {
		c := s[i]
		switch {
		case c == '[':
			index := strings.Index(s[i:], "]")
			vm.currentSig = s[i+1 : i+index]
			i = i + index
		case c == ' ':
			c := string(tempChars)
			if c != " " && c != "new" {
				spaceStrings = append(spaceStrings, string(tempChars))
				//fmt.Println(string(tempChars))
				tempChars = tempChars[0:0]
			}
		case c == '=':
			d := new(Define)
			if len(spaceStrings) == 2 {
				d.Name = spaceStrings[1]
				//Class spaceStrings[0]
			}
			if strings.Contains(string(tempChars), "@") {
				d.Name = string(tempChars[1:])
			}
			//Assignment
			tempChars = tempChars[0:0]
			assig := new(Assignment)
			assig.To = d
			nh.Push(assig)
			//tempChars = tempChars[0:0]
		case c == '(': //invoke sub Method
			ex := string(tempChars)
			cal := new(Call)
			next := matchingNextSymbol('(', s[i:])
			if next > 1 {
				cal.Args = vm.getArgs(s[i+1 : i+next])
			}
			if strings.HasPrefix(ex, "new") { //instance
				cal.Owner = Class{Name: strings.ReplaceAll(ex, "new", "")}
				cal.ClassTyp = NEWOBJECT
				cal.Method = MethodMeta{
					Name: "init",
				}
			} else { //static
				i := strings.LastIndexByte(ex, '.')
				cal.Owner = Class{ex[0:i]}
				cal.ClassTyp = STATICOJB
				cal.Method = MethodMeta{
					Name: ex[i+1:],
					Sig:  GetSig(vm.getCurrent()),
				}
			}
			topNode := nh.Top()
			if topNode != nil {
				if v, b := topNode.(*Assignment); b {
					v.From = cal
				}
				if e, b := topNode.(*Call); b {
					cal.Owner = e
					*e = *(deepcopy.Copy(cal).(*Call))
				}
			}
			nh.Push(cal)
			topNode = cal //@
			tempChars = tempChars[0:0]
		case c == ';':
			//nodes = append(nodes, topNode)
			tempChars = tempChars[0:0]
			spaceStrings = spaceStrings[0:0]
			if root := nh.Root(); root != nil {
				nodes = append(nodes, root.(Expr))
			}
			nh.Clear()
		default:
			//fmt.Println(string(byte(c)))
			tempChars = append(tempChars, c)
		}
	}
	if nh.Len() != 0 {
		return []Expr{nh.Root().(Expr)}
	}
	return nodes

}
func (vm *Vm) getArgs(args string) []Expr {
	var v []Expr
	sps := strings.Split(args, ",")
	for _, sp := range sps {
		switch {
		case regexp.MustCompile("[A-Za-z]").MatchString(sp):
			v = append(v, vm.Parse(sp)...)
		case strings.Contains(sp, "$"):
			i, e := strconv.ParseInt(sp[1:], 10, 32)
			if e != nil {
				panic(e)
			}
			v = append(v, Arg{ArgN: i, Typ: 0})
		case strings.Contains(sp, "@"):
			i, e := strconv.ParseInt(sp[1:], 10, 32)
			if e != nil {
				panic(e)
			}
			v = append(v, Arg{ArgN: i, Typ: 1})
		default:
			v = append(v, vm.Parse(sp)...)
		}
	}
	return v
}

func (vm *Vm) getCurrent() string {
	//if vm.currentSig == "" {
	//	panic("get sig error vm.currentSig == null")
	//}
	rel := vm.currentSig
	vm.currentSig = ""
	return rel
}

func (vm *Vm) run(exp Expr) {
	switch exp.(type) {
	case *Assignment:
		as := exp.(*Assignment)
		vm.run(as.From)
	case *Call:
		vm.call(exp.(*Call))
	case *Define:

	default:
		log.Fatal("not support")
	}
}

func (vm *Vm) call(c *Call) {
	var jcls jni.Jclass
	switch c.Owner.(type) {
	case *Class:
		cls := c.Owner.(*Class)
		jcls = vm.env.FindClass(cls.Name)
	case *Call:
		vm.call(c.Owner.(*Call))
	}
	jmethod := vm.env.GetMethodID(jcls, c.Method.Name, c.Method.Sig.String())
	if c.ClassTyp == STATICOJB { //static call
		sig := "Object"
		if s, b := sv[c.Method.Sig.RetTyp]; b {
			sig = s
		}
		method := reflect.ValueOf(vm.env).MethodByName(fmt.Sprintf("CallStatic%sMethodA", sig))
		ret := method.Call([]reflect.Value{reflect.ValueOf(jcls), reflect.ValueOf(jmethod)})
		fmt.Println(ret)
	} else if c.ClassTyp == NEWOBJECT {

	}

}
