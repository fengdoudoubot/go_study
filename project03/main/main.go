package main
import(
	"fmt"
	"strings"
	"sort"
	_"math/rand"
	_"time"
	"errors"
	_"strconv"
)

// var te = test()
// func test() int{
// 	fmt.Println("test")
// 	return 1
// }

//执行顺序：被引入的包中的全局变量定义->init->全局变量定义->init->main
//完成初始化工作
// func init(){
// 	fmt.Println("init")
// }

//全局匿名函数
var(
	Fun1 = func(n1 int ,n2 int )int {
		return n1 + n2
	}
)

//闭包
func AddUpper() func(int) int {//返回函数
	var n int = 10
	return func(x int) int{//这个函数和n形成整体构成闭包
		n = n + x
		return n
	}
}
func makeSuffix(suffix string) func(string) string{
	return func (name string ) string{
		if !strings.HasSuffix(name ,suffix){
			return name + suffix
		}
		return name
	}
}

//defer,及时释放函数创建的资源，如文件资源、数据库连接
func sum(n1 int ,n2 int )int {
	defer fmt.Println("ok1 ",n1)//压入栈(独立的栈)中，先入后出,函数执行完毕后执行
	defer fmt.Println("ok2 ",n2)
	//值也将拷贝入栈
	n1++
	n2++
	fmt.Println("ok4",n1)
	fmt.Println("ok5",n2)
	res := n1 + n2
	defer fmt.Println("ok3 ",res)
	return res
}

//错误处理
func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover()  // recover()是内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//自定义错误
func readConf(name string) (erro error){
	if name == "xx.ini" {
		return nil
	} else{
		return errors.New("nono")
	}
}

//var Age int = 20//right
//Name := "aaa"//error 赋值语句不能出现在函数体外

func fbn(n int)([]uint64){
	fbnSlice := make([]uint64,n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	}
	return fbnSlice
}

func main(){
//	fmt.Println("main")

//匿名函数
	//1、
	// res1 := func(n1 int ,n2 int )int {
	// 	return n1 + n2
	// }(10,20)
	// fmt.Println("res1:",res1)
	// //2、
	// a := func(n1 int ,n2 int )int {
	// 	return n1 + n2
	// }
	// res2 := a(10,10)
	// fmt.Println("res2:",res2)
	// //全局匿名函数
	// res3 := Fun1(1,1)
	// fmt.Println("res3:",res3)

//闭包
	// add := AddUpper()
	// fmt.Println(add(1))
	// fmt.Println(add(2))
	// fmt.Println(add(3))
	
	// f := makeSuffix(".jpg")
	// fmt.Println(f("aaa.jpg"))
	// fmt.Println(f("bbb"))
	
//defer
	// res := sum(1,2)
	// fmt.Println("res ",res)

//字符串常用函数
	// str := "hello北京"
	// //长度
	// fmt.Println(len(str))//中文占3个字节
	// //遍历
	// str2 := []rune(str)
	// for i := 0;i < len(str2);i++{
	// 	fmt.Printf("%c  ",str2[i])
	// }
	// //字符串转整数
	// n,err := strconv.Atoi("hi")
	// if err != nil{
	// 	fmt.Println("e!")
	// }else{
	// 	fmt.Println(n)
	// }
	// //字符串转整数
	// str = strconv.Itoa(1234)
	// fmt.Println(str)
	// //字符串转[]byte
	// var bytes = []byte("hi,youo")
	// //byte[]转string
	// str = string([]byte{97,98,99})

//错误处理
	// test()	
	// fmt.Println("ok")

//自定义错误
	// err := readConf("nnn")
	// if(err != nil){
	// 	panic(err)
	// }

//数组在go中是值类型,默认值传递
	// var arr [3]int = [3]int{1,2,3}
	// var arr2 = [3]int{0,2,4}
	// var arr3 = [...]int{9,1,2,3}
	// var arr4 = [...]int{1:30,0:20,2:100}
	// fmt.Println(arr4)
	// for i,v := range arr4{
	// 	fmt.Println(i,"xx",v)
	// }
	
	// rand.Seed(time.Now().UnixNano())
	// var arr [5]int
	// for i :=0;i < len(arr);i++{
	// 	arr[i] = rand.Intn(100)
	// }
	// fmt.Println(arr)
	// temp := 0
	// for i := 0;i < len(arr) / 2;i++{
	// 	temp = arr[len(arr) - 1 - i]
	// 	arr[len(arr) - 1 - i] = arr[i]
	// 	arr[i] = temp
	// }
	// fmt.Println(arr)

//切片
	//1)
	// var arr [5]int = [5]int{1,2,3,4,5}
	// //切片是数组的引用，引用类型
	// slice := arr[1:3]//从下标1到下标3不包含3
	// fmt.Println("slice 的元素是 =", slice) 
	// fmt.Println("slice 的元素个数 =", len(slice)) 
	// fmt.Println("slice 的容量 =", cap(slice)) // 切片的容量是可以动态变化
	// //切片从底层来说其实是一个结构体

	// //2
	// var slice2 []int = make([]int,5 ,10)

	// //3
	// var slice3 []string = []string{"aa","bb","c"}
	
	// //切片可以继续切片
	// slice4 := slice[1:2]
	
	// //append动态追加
	// slice5 :=[]int{100,200,300}
	// slice6 := append(slice5,400,600)//slice5没有变化
	// slice4 = append(slice,slice5...)
	
	//copy
	// var slice4 []int = []int{1, 2, 3, 4, 5}
	// var slice5 = make([]int,3,10)
	// copy(slice5, slice4)
	// fmt.Println("slice4=", slice4)
	// fmt.Println("slice5=", slice5)//只拷贝5的len长度的值

	//string（本质是一个Byte数组）也可以切片处理,但string是不可变的
	//修改string
	// str := "hello"
	// arr := []byte(str)
	// arr[0] = 'z'
	// str = string(arr)
	//上面不能处理中文，byte是按字节处理,将byte换成rune，rune是按照字符处理

	// fnbSlice := fbn(20)
	// fmt.Println("fnbSlice=", fnbSlice)

//二维数组
	//1
	// var arr [4][6]int
	//2
	// var arr3 [2][3]int = [2][3]int{{1,2,3},{2,3,4}}

	// for i, v := range arr3 {
	// 	for j, v2 := range v {
	// 		fmt.Printf("arr3[%v][%v]=%v \t",i, j, v2)
	// 	}
	// 	// fmt.Println()
	// 	// fmt.Println(i,  v)
	// }

//map
	//map是引用类型。map会自动扩容

	//slice、map、function 不能用来做key因为不能用==来判断
	//map声明之后不会分配空间，使用前需要先make
	// var a map[string]string
	// a = make(map[string]string,10)
	// a["no1"] = "jia"//不make直接赋值会报错
	// //key重复会覆盖//key是无序的//value可以重复
	// cities := make(map[string]string)
	// cities["no1"] = "北京"
	// cities["no2"] = "天津"
	// cities["no3"] = "上海"

	// fmt.Println( len(cities))

	// heroes := map[string]string{
	// 	"hero1" : "宋江",
	// 	"hero2" : "卢俊义",
	// 	"hero3" : "吴用",//最后也要，
	// }
	// heroes["hero4"] = "林冲"

	// studentMap := make(map[string]map[string]string)
	
	// studentMap["stu01"] =  make(map[string]string,3)
	// studentMap["stu01"]["name"] = "tom"
	// studentMap["stu01"]["sex"] = "男"
	// studentMap["stu01"]["address"] = "北京长安街~"

	// studentMap["stu02"] =  make(map[string]string) //这句话不能少!!
	// studentMap["stu02"]["name"] = "mary"
	// studentMap["stu02"]["sex"] = "女"
	// studentMap["stu02"]["address"] = "上海黄浦江~"

	// fmt.Println( len(studentMap))
	// fmt.Println( len(studentMap["stu01"]))
	// fmt.Println(studentMap)
	// fmt.Println(studentMap["stu02"])

	// delete(cities,"no1")
	//删除指定的key不存在，不报错！也不操作
	//删除全部除了遍历还可以可以map = make(map[string]string)

	//看某个key是否存在
	// val,ok := cities["no1"]
	// if ok {
	// 	fmt.Printf("有no1 key 值为%v\n", val)
	// } else {
	// 	fmt.Printf("没有no1 key\n")
	// }

	//map遍历for-range
	// for k,v := range cities{
	// 	fmt.Printf("k=%v v=%v\n", k, v)
	// }
	// for _,v1 := range studentMap{
	// 	for k2, v2 := range v1{
	// 		fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
	// 	}
	// 	fmt.Println()
	// }



//map切片
	// var monsters []map[string]string
	// monsters = make([]map[string]string,2)
	// if monsters[0] == nil{
	// 	monsters[0] = make(map[string]string)
	// 	monsters[0]["name"] = "牛魔王"
	// 	monsters[0]["age"] = "500"
	// }

	// if monsters[1] == nil {
	// 	monsters[1] = make(map[string]string, 2)
	// 	monsters[1]["name"] = "玉兔精"
	// 	monsters[1]["age"] = "400"
	// }

	// newMonster := map[string]string{
	// 	"name" : "新的妖怪~火云邪神",
	// 	"age" : "200",
	// }
	// monsters = append(monsters, newMonster)

//map排序
	// map1 := make(map[int]int, 10)
	// map1[10] = 100
	// map1[1] = 13
	// map1[4] = 56
	// map1[8] = 90

	// var keys []int
	// for k, _ := range map1 {
	// 	keys = append(keys, k)
	// }

	// sort.Ints(keys)
	// fmt.Println(keys)

	// for _, k := range keys{
	// 	fmt.Printf("map1[%v]=%v \n", k, map1[k])
	// }

//	
}

//init函数、匿名函数、闭包、defer
//字符串常用函数、错误处理、自定义错误、数组
//切片（slice）、二维数组、map
