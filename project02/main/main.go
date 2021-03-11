package main
import(
	"fmt"
	_"go_code/project02/utils"
	//abc "go_code/project02/utils"//别名abc
	_"time"
	_"math/rand"
	_"math"
	_"go_code/project02/ex"
	_"unsafe"
	_"strconv"
)
//别名
type othFun func(int , int )int

func test() bool{
	fmt.Println("test!!")
	return true
}

//求和与差
func getSumAndSub(n1 int , n2 int ) (int ,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum,sub
}

//斐波那契数
func fbn(n int ) int{
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

//猴子吃桃子
func peach(n int) int{
	if n > 10 ||n < 1 {
		fmt.Println("error")
	}
	if n == 10{
		return 1
	} else {
		return (peach(n + 1) +1) * 2
	}
}

func test2(n1 *int){
	*n1 = *n1 + 10
	fmt.Println("test!!",*n1)
}

func getSum(n1 int ,n2 int) int{
	return n1 + n2
}
//也可以写成func getSum(n1,n2 int) int {}

//函数做形参
func myFun(funvar func(int ,int ) int ,num1 int ,num2 int) int {
	return funvar(num1,num2)
}

func myFun2(funvar2 othFun ,num1 int ,num2 int) int {
	return funvar2(num1,num2)
}

//支持对返回值命名
func getSumAndSub2(n1 int , n2 int ) (sum int ,sub int) {
	sum = n1 + n2//
	sub = n1 - n2
	return //不用再写名字了
}

//可变参数
func sum(n1 int,args... int) int{//args是切片,可变参数要放在形参列表的最后
	sum := n1
	for i := 0 ;i < len(args);i++{
		sum +=args[i]
	}
	return sum
}

//main函数-》package main
func main(){
	
//fmt.Println(ex.NumOther)//访问其他包的变量
	
//除法
	// fmt.Println(10/4)
	// var n1 float32 = 10/4
	// fmt.Println(n1)
	
	// var n2 float32 = 10.0/4
	// fmt.Println(n2)
	
//% 
	//a % b = a - a / b *b
	// fmt.Println(10%4)

//++  --
	//b := n1--//error
	//++  --只能独立使用！！！
	//--i//error只能写在后面

//逻辑运算符
	//短路与
	// if 10 < 9 && test(){
	// 	fmt.Println("noe!!")
	// }

//赋值运算
	// var a = 10
	// var b = 1
	// a = a + b
	// b = a-b
	// a = a-b
	//无中间量交换值

//接收键盘输入
	// var name string
	// var age byte
	// var sal float32
	// //1
	// fmt.Println("name:")
	// fmt.Scanln(&name)
	// fmt.Println("age:")
	// fmt.Scanln(&age)
	// fmt.Println("sal:")
	// fmt.Scanln(&sal)
	// fmt.Printf("名字是%v\n年龄是%v\n薪水是%v", name, age, sal)
	// //2
	// fmt.Println("空格隔开")
	// fmt.Scanf("%s %d %f",&name,&age,&sal)
	// fmt.Printf("名字是%v\n年龄是%v\n薪水是%v", name, age, sal)

//进制输出
	// var i int = 5
	// fmt.Printf("%b\n",i)
	// var j = 011
	// fmt.Println("j=",j)
	// var x = 0x11
	// fmt.Println("x=",x)


//if
	// var age int
	// fmt.Println("in:")
	// fmt.Scanln(&age)

	// if age>10 {
	// 	fmt.Println("ok")
	// }

	// if age := 10 ; age > 0{
	// 	fmt.Println("iii")
	// } 

	// if age := -1 ; age > 0 {
	// 	fmt.Println("iii",age)
	// } else {
	// 	fmt.Println("sss",age)
	// }

	// var n1 float64 = 12.0
	// var n2 float64 = 18.0
	// if n1 > 10.0 && n2 < 20.0 {
	// 	fmt.Println("ddd")
	// }

	// var year int = 2020
	// if (year % 4 == 0 && year % 100 != 0) || year % 400 ==0 {
	// 	fmt.Println("y")
	// }

	// var b bool = false
	// if b = false {//error
	// 	fmt.Println("dd")
	// }

	// var a float64 = 3.0
	// var b float64 = 100.0
	// var c float64 = 6.0
	
	// m := b * b - 4 * a *c
	// if m > 0 {
	// 	x1 := (-b + math.Sqrt(m)) / 2 * a
	// 	x2 := (-b - math.Sqrt(m)) / 2 * a
	// 	fmt.Println(x1,"\t",x2)
	// } else if m == 0 {
	// 	x1 := (-b + math.Sqrt(m)) / 2 * a
	// 	fmt.Println(x1)
	// } else {
	// 	fmt.Println("nnn")
	// }

//switch
	// var k byte
	// fmt.Println("in:")
	// fmt.Scanf("%c",&k)

	// switch k {
	// case 'a':
	// 	fmt.Println("1")
	// case 'b':
	// 	fmt.Println("2")
	// case 'c':
	// 	fmt.Println("3")
	// case 'd':
	// 	fmt.Println("4")
	// case 'f':
	// 	fmt.Println("5")
	// default://不是必须的
	// 	fmt.Println("000")
	// }

	//	var n1 int32 = 20
	//	var n2 int64 = 20
	// 	switch n1{
	// //		case n2 ://error类型不匹配
	// 		case 10,4,20:
	// 			fmt.Println("000")
	// 		default:
	// 			fmt.Println("11")
	// 	}

	// var a int = 20
	// switch{//可以不带表达式
	// 	case a == 10 :
	// 		fmt.Println("000")
	// 	case a == 20 :
	// 		fmt.Println("11")
	// }
	
	// switch a := 10; {//可以不带表达式
	// 	case a == 10 :
	// 		fmt.Println("000")
	// 	case a == 20 :
	// 		fmt.Println("11")
	// }

	// var num int =10
	// switch num {
	// 	case 10:
	// 		fmt.Println("ok1")
	// 		fallthrough//穿透，只能穿透一层
	// 	case 20:
	// 		fmt.Println("ok0")
	// 	case 30:
	// 		fmt.Println("ok2")
	// }

//for
	// //1
	// for i := 1 ;i <=10 ;i++{
	// 	fmt.Println(i)
	// }
	// //2
	// j := 1
	// for j <= 10 {
	// 	fmt.Println(j)
	// 	j++
	// }
	// //3
	// k := 1
	// for {
	// 	fmt.Println(k)
	// 	k++
	// 	if k == 5{
	// 		break
	// 	}
	// }

	//1
	// var str string = "hello,world!"
	// for i := 0 ;i < len(str) ;i++{
	// 	fmt.Printf("%c \n",str[i])
	// }
	//2
	// for index,val := range str{
	// 	fmt.Printf("index:%d,val:%c\n",index,val)
	// }
	//有中文情况不一样可以用切片或者for-range

//golang没有while/do……while
	//代替while
	// var i int = 1
	// for{
	// 	if i > 10{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
	//代替do……while
	// var i int = 1
	// for{
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 10{
	// 		break
	// 	}
	// }

//多重循环
	// var totalLevel int = 6
	// for i := 1;i <= totalLevel; i++ {
	// 	for j := 1;j <= totalLevel - i;j++{
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1;j <= 2 * i -1; j++{
	// 		if j == 1 || j == 2 * i - 1 || i == totalLevel{
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	} 
	// 	fmt.Println()
	// }
	
//break
	//随机数
	// var count int = 0
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100) + 1
	// 	fmt.Println("n=", n)
	// 	count++
	// 	if (n == 99) {
	// 		break //表示跳出for循环
	// 	}
	// }

	// fmt.Println("一共使用了 ", count)

	//break 标签
	// lable2: 
	// for i := 0; i < 4; i++ {
	// 	//lable1: 
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			break lable2 
	// 		}
	// 		fmt.Println("j=", j) 
	// 	}
	// }

//continue
	// lable2: 
	// for i := 0; i < 4; i++ {
	// 	//lable1: 
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue lable2
	// 		}
	// 		fmt.Println("j=", j) 
	// 	}
	// }

//goto(不主张)
	// var n int = 30
	// fmt.Println("ok1")
	// if n > 20 {
	// 	goto label1
	// }
	// fmt.Println("ok2")
	// fmt.Println("ok3")
	// fmt.Println("ok4")
	// label1:
	// fmt.Println("ok5")
	// fmt.Println("ok6")
	// fmt.Println("ok7")

//函数
	// res := utils.Cal(1.1, 2.1, '+')//C！大写//abc.Cal(1.1, 2.1, '+')
	// fmt.Println(res)

	// res1,res2 := getSumAndSub(3,1)
	// fmt.Printf("res1=%v res2=%v\n", res1, res2)
	// //只想要一个值
	// res3,_ :=getSumAndSub(2,1)
	// fmt.Println("res3=", res3)
	
//递归调用
	// fmt.Println(fbn(4))
	// fmt.Println("1:",peach(1))

//指针传递
	// num := 20
	// test2(&num)
	// fmt.Println("tmain!",num)

//函数也是一种数据类型,可赋值调用，也可做形参
	// a :=getSum
	// res := a(10,10)
	// fmt.Printf("a的类型%T，getSum类型%T,%v\n", a , getSum, res)
	// res2 := myFun(getSum,20,20)
	// fmt.Println("2:",res2)

//支持自定义数据类型
	// type myInt int//别名，但是二者还是不同类型
	// var num1 myInt
	// num1 = 40
	// var num2 int
	// //num2 = num1//error//go认为这是不同类型

	// res3 := myFun2(getSum,11,12)
	// fmt.Println("2:",res3)

//支持对返回值命名/
	// sum,sub := getSumAndSub2(3,2)
	// fmt.Println(sum,",,", sub)

//支持可变参数
	res4 := sum(10,0,10,12)
	fmt.Println("2:",res4)

}

//运算符、键盘输入、循环、分支控制(if-else、switch)
//函数、包