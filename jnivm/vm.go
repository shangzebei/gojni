package jnivm

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/loader"
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
	objs       ojbHeap
	methods    map[string]reflect.Method
}

func RunExpr(exp jparser.Expr) value {
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	return v.run(exp)
}

func RunSource(s string) value {
	cmp := jparser.Compiler{}
	v := VM{env: jni.AutoGetCurrentThreadEnv()}
	return v.runSegment(cmp.Parse(s))
}

func RunBitCode(byes []byte) {
	panic("not impl")
}

var methods map[string]reflect.Method

func GetJNIMethods() map[string]reflect.Method {
	return methods
}

func GetMethodWithName(name string) *reflect.Method {
	if v, b := methods[strings.ToLower(name)]; b {
		return &v
	} else {
		return nil
	}
}

func init() {
	var vm jni.Env
	methods = make(map[string]reflect.Method)
	f := reflect.TypeOf(vm)
	for i := 0; i < f.NumMethod(); i++ {
		m := f.Method(i)
		methods[strings.ToLower(m.Name)] = m
	}
}

func (vm *VM) runSegment(exp []jparser.Expr) value {
	var v value
	for _, expr := range exp {
		v = vm.run(expr)
	}
	return v
}

func (vm *VM) doAssignment(ass *jparser.Assignment) value {
	panic("not impl")
	return value{}
}

func (vm *VM) doDefine(ass *jparser.Define) value {
	panic("not impl")
	return value{}
}

func (vm *VM) run(exp jparser.Expr) value {
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
	return value{}
}

func (vm *VM) doCall(c *jparser.Call) value {
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
		jni.CheckNull(jMethod, fmt.Sprintf("not find static method %s [ %s ]", c.Method.Name, c.Method.Sig))
		ret := callJni(fmt.Sprintf("CallStatic%sMethodA", sig), vm.env, jCls, jMethod)
		return value{sig: c.Method.Sig, v: *ret}
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
		ret := callJni(fmt.Sprintf("Call%sMethodA", sig), vm.env, jObj, jMethod)
		return value{sig: c.Method.Sig, v: *ret}
	case jparser.OBJECTINVOKE:
		jni.CheckNull(jObj, "obj is null")
		jMethod := vm.env.GetMethodID(jCls, c.Method.Name, c.Method.Sig.Sig)
		jni.CheckNull(jMethod, fmt.Sprintf("not find OBJECTINVOKE method %s", c.Method.Name))
		ret := callJni(fmt.Sprintf("Call%sMethodA", sig), vm.env, jObj, jMethod)
		return value{sig: c.Method.Sig, v: *ret}
	default:
		panic(fmt.Sprintf("not support ClassTyp %d ", c.ClassTyp))
	}

}

func callJni(format string, args ...interface{}) *reflect.Value {
	m := GetMethodWithName(format)
	if m == nil {
		panic("method def is error")
	}
	var params []reflect.Value
	index := m.Type.NumIn() - 1
	for i, v := range args {
		of := reflect.ValueOf(v)
		if i >= index {
			params = append(params, reflect.ValueOf(uint64(loader.JabValueToUint(of))))
		} else {
			params = append(params, of)
		}
	}
	//fmt.Println(params)
	ret := m.Func.Call(params)
	if len(ret) != 0 {
		return &ret[0]
	}
	return nil
}
