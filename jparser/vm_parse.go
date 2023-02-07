package jparser

import (
	"fmt"
	"github.com/shangzebei/gojni/jni"
	"github.com/shangzebei/gojni/native"
	"github.com/shangzebei/gojni/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/mohae/deepcopy"
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

type MethodMeta struct {
	Name string
	Sig  *native.MethodSig
}

type Call struct {
	ClassTyp int
	Owner    Expr
	Method   *MethodMeta
	Args     []Expr
}

type Arg struct {
	Typ  int
	ArgN int64
}

func (vm *Compiler) checkSyntax(s string) bool {
	if !strings.HasSuffix(s, ";") {
		return false
	}
	return true
}

func (vm *Compiler) Parse(s string) []Expr {
	if vm.check && !vm.checkSyntax(s) {
		return nil
	}
	s = strings.ReplaceAll(s, "\n", "")
	var tempChars []byte
	var spaceStrings []string
	var nodes []Expr
	var nh NodeHeap
	var currentSig StrHeap
	for i := 0; i < len(s); i += 1 {
		c := s[i]
		switch {
		case c == '[':
			index := strings.Index(s[i:], "]")
			currentSig.Push(s[i+1 : i+index])
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
			next := utils.MatchingNextSymbol('(', s[i:])
			if next > 1 {
				cal.Args = vm.getArgs(s[i+1 : i+next])
				i += next
			}
			if strings.HasPrefix(ex, "new") { //instance
				cal.Owner = &Class{Name: strings.ReplaceAll(ex, "new", "")}
				cal.ClassTyp = NEWOBJECT
				cal.Method = &MethodMeta{
					Name: "<init>",
				}
			} else { //static
				i := strings.LastIndexByte(ex, '.')
				cal.Owner = &Class{ex[0:i]}
				cal.ClassTyp = STATICOJB
				cal.Method = &MethodMeta{
					Name: ex[i+1:],
				}
			}
			if sig := currentSig.Pop(); sig != nil {
				cal.Method.Sig = native.EncodeToSig(sig.(string))
			} else {
				jni.ThrowException(fmt.Sprintf("method [%s] no sign in %s", cal.Method.Name, s))
			}
			topNode := nh.Top()
			if topNode != nil {
				if v, b := topNode.(*Assignment); b {
					v.From = cal
				}
				if e, b := topNode.(*Call); b {
					cal.Owner = e
					cal.ClassTyp = OBJECTINVOKE
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
func (vm *Compiler) getArgs(args string) []Expr {
	var v []Expr
	sps := strings.Split(args, ",")
	for _, sp := range sps {
		switch {
		case regexp.MustCompile("[A-Za-z]").MatchString(sp):
			v = append(v, vm.Parse(sp)...)
		case strings.Contains(sp, "$"):
			i, e := strconv.ParseInt(sp[1:], 10, 32)
			if e != nil {
				jni.ThrowException(e.Error())
			}
			v = append(v, Arg{ArgN: i, Typ: 0})
		case strings.Contains(sp, "@"):
			i, e := strconv.ParseInt(sp[1:], 10, 32)
			if e != nil {
				jni.ThrowException(e.Error())
			}
			v = append(v, Arg{ArgN: i, Typ: 1})
		default:
			v = append(v, vm.Parse(sp)...)
		}
	}
	return v
}
