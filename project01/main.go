package main
// import "fmt"
// import "unsafe"
//一次引用多个包
import(
	"fmt"
	_"unsafe"
	_"strconv"
	//_表示暂时忽略一个包
)
//一次定义多个全局变量
	// var(
	// 	n1 = 200
	// 	n2 = "so"
	// 	n3 = 10.10
	// )
func main(){

/*	
//一次定义多个变量的3种
	var x1,x2,x3 int  //int占多大与系统关
	fmt.Println("x1=",x1, "\tx2=",x2,"\tx3=",x3)

	var z1, z2, z3 = 10, "qq", 2.2
	fmt.Println("\nz1=",z1,"\tz2=",z2,"\tz3=",z3)
	
	a1, a2, a3 := 110, "aaq", 2.22
	fmt.Println("\na1=",a1,"\ta2=",a2,"\ta3=",a3)

	fmt.Println("\nn1=",n1,"\tn2=",n2,"\tn3=",n3)

//数据范围
	var j int8 = 127//-128-127
	var k uint16 = 255//0-255
	
	var b byte = 255//0-255
	fmt.Println("\nk=",k,"\tj=",j,"\tb=",b)

	//查看变量的数据类型和字节大小
	fmt.Printf("\nk的类型 ：%T，字节数：%d",k,unsafe.Sizeof(k))
	
	var num1 = 1.1//默认float64
	fmt.Printf("\nnum的类型 ：%T，字节数：%d",num1,unsafe.Sizeof(num1))
	
//字符
	var num2 byte = 'a'
	var num3 byte = '0'
	fmt.Println("\nnum2=",num2,"  num3=",num3)//码值输出
	fmt.Printf("\nnum2=%c  num3=%c",num2,num3)//对应字符
	//var num4 byte = '北'//溢出
	var num4 int = '北'
	fmt.Printf("\nnum4=%c",num4)//对应字符
	//按码值运算
	var num5 = 10 + 'a'
	fmt.Println("\nnum5=",num5)

//字符串
	//golang中统一UTF-8，字符串一旦赋值就不能修改了
	var str = "hello"
	//str[0] = 'a'//error
	fmt.Println("\nstr=",str)
	//``反引号输出可防止攻击
	var str2 string = `
	mport(
		"fmt"
		"unsafe"
	)
	`
	fmt.Println("\nstr2=",str2)

	//+保留在上一行
	var str3 = "hh" + "ddd" +"sss"+
	"ssss"+"d"
	fmt.Println("\nstr3=",str3)

//基本数据类型的转换
	var n1 int32 = 100
	var n2 float32 = float32(n1)
	fmt.Printf("\nn1=%v n2=%v",n1,n2)
	//不会报错按溢出处理
	var n3 int64 = 99999
	var n4 int8 = int8(n3)
	fmt.Printf("\nn3=%v n4=%v",n3,n4)

//基本数据类型转换string
	
	//1、Sprintf函数
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var char1 byte = 'h'
	var str string //空的str

	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n", str, str)//%q双引号括起来

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", char1)
	fmt.Printf("str type %T str=%q\n", str, str)

	//2、
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n", str, str)
	
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)
	
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)



//string->其他
	var str string = "true"
	var b bool

	b , _ = strconv.ParseBool(str)//_忽略另一个返回值
	fmt.Printf("b type %T  b=%v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T  n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

*/

//指针
	var i int = 10
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("i=%v\n",*ptr)
	
	*ptr = 101
	fmt.Printf("i=%v\n",*ptr)
	fmt.Printf("i=%v\n",i)
	
}
 

//变量、基本数据类型、指针