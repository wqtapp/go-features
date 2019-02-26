package params

import (
	"fmt"
	"testing"
)

func TestValueParam(t *testing.T){
	num := 1
	if !testFuncParam(num,num){
		t.Error("testFuncParam error")
	}
}

func testValueAssign(t *testing.T){
	num1 := 1
	num2 := num1
	//变量赋值，两个变量的值相同，但不是一个变量（变量地址不同）
	if &num1 == &num2{
		t.Error("value assign pointer equals")
	}
}

func testPointerAssign(t testing.T){
	num := 1
	num1 := &num
	num2 := &num
	//指针指向变量的值跟之前的变量值相同
	if *num1 != num || *num2 != num {
		t.Error("ptr assign not equal the value")
	}
	//指向同一个变量的指针值相同，都是变量的地址
	if num1 != num2{
		t.Error("ptr not equal")
	}
	//指向同一个变量的不通指针的地址不同
	if &num1 == &num2{
		t.Error("different ptr addr is the same")
	}

	*num1 = 2
	//使用指针更改所指向值的值
	if *num1 != num || *num2 != num{
		t.Error("ptr change value fail")
	}
}

func testFuncParam(param1,param2 int) bool{
	//值参数传递过来的是参数的副本，所以参数值一致
	if param1 != param2 {
		fmt.Println("error: same param not equals")
		return false
	}
	//值参数传递过来的是值得副本，所以不同的参数的变量地址不同
	if &param1 == &param2{
		fmt.Println("error: value params with same pointer")
		return false
	}
	return true
}

func testStructAssign(t testing.T){
	object := Struct{}
	subObj := SubStruct{object}
	ptrSubObj := PtrSubStruct{&object}
	//子类型中的匿名类型是副本
	if &subObj.Struct == &object || &subObj.num == &object.num{
		t.Error("obj the same address")
	}
	object.num = 2
	subObj.num = 1
	//再次证明，无相关性，是值得副本
	if object.num == subObj.num{
		t.Error("obj the same address")
	}

	//子类型中的匿名类型是子类型变量的副本，依然指向子类型本身
	if ptrSubObj.Struct != &object || &ptrSubObj.num != &object.num{
		t.Error("ptrObj the same address")
	}
	object.num = 3
	//再次证明，相关性，是值得指针的副本
	if object.num != ptrSubObj.num{
		t.Error("obj the same address")
	}
}

type Struct struct {
	num int
}

type SubStruct struct {
	Struct
}

type PtrSubStruct struct {
	*Struct
}