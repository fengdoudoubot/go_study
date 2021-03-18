package main
import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Student struct {

}

func TypeJudge(items... interface{}) {
	for index, x := range items {
		switch x.(type) {
			case bool :
				fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
			case float32 :
				fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
			case float64 :
				fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
			case string :
				fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
			case Student :
				fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
			case *Student :
				fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, x)
			default :
				fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, x)
		}
	}
}

func main()  {
//类型断言
	// var a interface{}
	// var point Point = Point{1,2}
	// a = point
	// var b Point
	// b = a.(Point)
	// fmt.Println(b)

	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2  
	// // y := x.(float64)//panic

	// if y, ok := x.(float64); ok {//检测转化是否成功，不报panic
	// 	fmt.Println("convert success")
	// 	fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	// } else {
	// 	fmt.Println("convert fail")
	// }
	// fmt.Println("继续执行...")

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)



}
///类型断言
