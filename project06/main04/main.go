package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{})  {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal", rVal)
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func reflect01(b interface{})  {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)

}

type Monster struct {
	Name  string `json:"name"`
	Age   int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
	
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd !=  reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	
	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII码）
	val.Method(1).Call(nil) //获取到第二个方法

	
	//调用结构体的第1个方法Method(0)
	var params []reflect.Value  //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/

}

func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

func main()  {
	// var num int = 100	
	// reflectTest(num)
	
	// var num int = 100	
	// reflect01(&num)
	// fmt.Println("num=", num)

	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}
//反射