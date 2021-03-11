package main
import (
	"fmt"
)

//如果结构体的字段类型是: 指针，slice，和map的零值都是 nil 还没有分配空间
//需要先make，才能使用.
type Person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int //指针 
	slice []int //切片
	map1 map[string]string //map
}

type Person2 struct{
	Name string
	Age int
}

type Per Person2//别名，但不是同一类型，只能强转

type A struct {
	Num int
}
type B struct {
	Num int
}

//struct tag
type Monster struct{
	Name string `json:"name"` 
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func (a A )test(){
	a.Num = 12//a是值传递
	fmt.Println(a.Num)
}
func(a *A) test2(){
	(*a).Num = 11
	fmt.Println(a.Num)
}

type integer int
func (i integer) print(){
	fmt.Println(i)
}

type Student struct {
	Name string
	Age int
}
//实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

type Stu struct {
	Name string
	Age int
}

func main(){

//结构体，值类型
	//1
	// var p1 Person
	// //2
	// p2 := Person2{}//也可以直接赋值
	// //3
	// var p3 *Person2= new(Person)
	// (*p3).Name = "smith" //也可以 p3.Name = "smith"
	// //4
	// var person *Person2 = &Person2{}
	// var person *Person2 = &Person2{"mary", 60}

	// fmt.Println(p1)

	// p1.slice = make([]int, 10)
	// p1.slice[0] = 100 //ok

	// p1.map1 = make(map[string]string)
	// p1.map1["key1"] = "tom" 
	// fmt.Println(p1)

	var stu5 *Stu = &Stu{"小王", 29} 
	stu6 := &Stu{"小王~", 39}

	var stu7 = &Stu{
		Name : "小李", 
		Age :49,
	}
//结构体之间的转换
	//完全相同的字段
	// var a A
	// var b B
	// a = A(b)

//struct tag
	// monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	// jsonStr, err := json.Marshal(monster)
	// if err != nil {
	// 	fmt.Println("json 处理错误 ", err)
	// }
	// fmt.Println("jsonStr", string(jsonStr))

//方法
	// var a A 
	// a.Num = 10
	// a.test()
	// fmt.Println(a.Num)
	// (&a).test2()//也可以a.test2()
	// fmt.Println(a.Num)
	// var i integer = 200
	// i.print()
	// //实现string方法
	// stu := Student{
	// 	Name : "oo",
	// 	Age : 20,
	// }
	// fmt.Println(&stu)//自动调用string方法

//	
}

//结构体、方法