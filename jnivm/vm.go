package jnivm

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/jparser"
	"gitee.com/aifuturewell/gojni/utils"
)

type VM struct {
	env        jni.Env
	tempResult interface{}
	objs       ojbHelp
}

func RunExpr(exp jparser.Expr) {
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	v.run(exp)
}

func RunSource(s string) {
	cmp := jparser.Compiler{}
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	v.runSegment(cmp.Parse(s))
}

func RunBitCode(byes []byte) {

}

func (vm *VM) runSegment(exp []jparser.Expr) {
	for _, expr := range exp {
		vm.run(expr)
	}
}

func (vm *VM) run(exp jparser.Expr) {
	switch exp.(type) {
	case *jparser.Assignment:
		as := exp.(*jparser.Assignment)
		vm.run(as.From)
	case *jparser.Call:
		vm.call(exp.(*jparser.Call))
	case *jparser.Define:

	default:
		log.Fatal("not support")
	}
}

func (vm *VM) call(c *jparser.Call) {

	//fmt.Println(c.String())

	var jCls jni.Jclass
	var jObj jni.Jobject

	switch c.Owner.(type) {
	case *jparser.Class:
		cls := c.Owner.(*jparser.Class)
		jCls = vm.env.FindClass(strings.ReplaceAll(cls.Name, ".", "/"))
		jni.CheckNull(jCls, cls.Name)
	case *jparser.Call:
		vm.call(c.Owner.(*jparser.Call))
	default:
		panic("not find c.Owner type " + reflect.ValueOf(c.Owner).Type().String())
	}

	sig := "Object"
	if s, b := utils.SigDecodeMap[c.Method.Sig.RetTyp]; b {
		sig = s
	}

	if o := vm.objs.Pop(); o != nil {
		fmt.Println("set v")
		oo := o.(ojb)
		jCls = oo.jCls
		jObj = oo.jObj
	}

	jni.CheckNull(jCls, "jCls is null")

	switch c.ClassTyp {
	case jparser.STATICOJB:
		jMethod := vm.env.GetStaticMethodID(jCls, c.Method.Name, c.Method.Sig.Sig)
		jni.CheckNull(jMethod, fmt.Sprintf("not find static method %s", c.Method.Name))
		CallMethod := reflect.ValueOf(vm.env).MethodByName(fmt.Sprintf("CallStatic%sMethodA", sig))
		ret := CallMethod.Call([]reflect.Value{reflect.ValueOf(jCls), reflect.ValueOf(jMethod)})
		log.Println(fmt.Sprintf("static obj ret %d", len(ret)))
	case jparser.NEWOBJECT:
		jMethod := vm.env.GetMethodID(jCls, c.Method.Name, c.Method.Sig.Sig)
		fmt.Println(c.Method.Name, c.Method.Sig.Sig)
		jni.CheckNull(jMethod, fmt.Sprintf("not find method %s", c.Method.Name))
		if c.Method.Name == "<init>" {
			fmt.Println("Push")
			jObj = vm.env.NewObjectA(jCls, jMethod)
			vm.objs.Push(ojb{
				jCls: jCls,
				jObj: jObj,
			})
		}
		jni.CheckNull(jObj, "obj is null")
		CallMethod := reflect.ValueOf(vm.env).MethodByName(fmt.Sprintf("Call%sMethodA", sig))
		ret := CallMethod.Call([]reflect.Value{reflect.ValueOf(jObj), reflect.ValueOf(jMethod)})
		log.Println(fmt.Sprintf("new obj ret %d", len(ret)))
	case jparser.OBJECTINVOKE:
		jni.CheckNull(jObj, "obj is null")
		jMethod := vm.env.GetMethodID(jCls, c.Method.Name, c.Method.Sig.Sig)
		jni.CheckNull(jMethod, fmt.Sprintf("not find OBJECTINVOKE method %s", c.Method.Name))
		CallMethod := reflect.ValueOf(vm.env).MethodByName(fmt.Sprintf("Call%sMethodA", sig))
		ret := CallMethod.Call([]reflect.Value{reflect.ValueOf(jObj), reflect.ValueOf(jMethod)})
		log.Println(fmt.Sprintf("new obj ret %d", len(ret)))
	default:
		panic(fmt.Sprintf("not support ClassTyp %d ", c.ClassTyp))
	}

}
