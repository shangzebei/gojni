package vm

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jparser"
	"gitee.com/aifuturewell/gojni/native"
	"gitee.com/aifuturewell/gojni/utils"
	"log"
	"reflect"
	"strings"

	"gitee.com/aifuturewell/gojni/jni"
)

type VM struct {
	env        jni.Env
	tempResult interface{}
	objs       ojbHeap
	methods    map[string]reflect.Method
}

func RunExpr(exp jparser.Expr) native.Value {
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	return v.run(exp)
}

func RunSource(s string) native.Value {
	cmp := jparser.Compiler{}
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	return v.runSegment(cmp.Parse(s))
}

func RunBitCode(byes []byte) {
	panic("not impl")
}

func (vm *VM) runSegment(exp []jparser.Expr) native.Value {
	var v native.Value
	for _, expr := range exp {
		v = vm.run(expr)
	}
	return v
}

func (vm *VM) doAssignment(ass *jparser.Assignment) native.Value {
	panic("not impl")
	return native.Value{}
}

func (vm *VM) doDefine(ass *jparser.Define) native.Value {
	panic("not impl")
	return native.Value{}
}

func (vm *VM) run(exp jparser.Expr) native.Value {
	switch exp.(type) {
	case *jparser.Assignment:
		as := exp.(*jparser.Assignment)
		return vm.doAssignment(as)
	case *jparser.Call:
		return vm.doCall(exp.(*jparser.Call))
	case *jparser.Define:
		return vm.doDefine(exp.(*jparser.Define))
	default:
		log.Fatal("not support")
	}
	return native.Value{}
}

func (vm *VM) doCall(c *jparser.Call) native.Value {
	var jCls jni.Jclass
	var jObj jni.Jobject

	switch c.Owner.(type) {
	case *jparser.Class:
		cls := c.Owner.(*jparser.Class)
		jCls = vm.env.FindClass(strings.ReplaceAll(cls.Name, ".", "/"))
		jni.CheckNull(jCls, cls.Name)
	case *jparser.Call:
		vm.doCall(c.Owner.(*jparser.Call))
	default:
		panic("not find c.Owner type " + reflect.ValueOf(c.Owner).Type().String())
	}

	sig := c.Method.Sig.RetTyp.GetType()

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
		jni.CheckNull(jMethod, fmt.Sprintf("not find static method %s [ %s ]", c.Method.Name, c.Method.Sig))
		ret := utils.CallJni(utils.GetFormatCallFunc("CallStatic%sMethodA", sig), vm.env, jCls, jMethod)
		return native.NewValue(c.Method.Sig.RetTyp, ret)
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
		ret := utils.CallJni(utils.GetFormatCallFunc("Call%sMethodA", sig), vm.env, jObj, jMethod)
		return native.NewValue(c.Method.Sig.RetTyp, ret)
	case jparser.OBJECTINVOKE:
		jni.CheckNull(jObj, "obj is null")
		jMethod := vm.env.GetMethodID(jCls, c.Method.Name, c.Method.Sig.Sig)
		jni.CheckNull(jMethod, fmt.Sprintf("not find OBJECTINVOKE method %s", c.Method.Name))
		ret := utils.CallJni(utils.GetFormatCallFunc("Call%sMethodA", sig), vm.env, jObj, jMethod)
		return native.NewValue(c.Method.Sig.RetTyp, ret)
	default:
		panic(fmt.Sprintf("not support ClassTyp %d ", c.ClassTyp))
	}

}
