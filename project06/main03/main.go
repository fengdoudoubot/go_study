package main

import (
	"fmt"
	"strconv"
	_"runtime"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int, 10)

	lock sync.Mutex
)

type Cat struct {
	Name string
	Age int
}

func test()  {

	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test1(n int)  {
	res := 1
	for i := 1; i <= n; i++ {
		res += i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func writeData(intChan chan int)  {
	for i := 1; i <=50; i++ {
		intChan <- i
		fmt.Printf("wr:%v\n",i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool)  {
	for{
		v,ok := <-intChan
		// time.Sleep(time.Second)
		
		if !ok{
			break
		}
		fmt.Printf("read:%v\n",v)
	}
	exitChan<-true
	close(exitChan)
}

func putNum(intChan chan int)  {
	fmt.Println("put!")
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	fmt.Println("poo!")
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool 
	for{
		num ,ok := <- intChan
		if !ok{
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
		}
	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	exitChan<- true
	
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
func test3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //error
}


func main()  {

//协程
	// go test()//开启协程

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(" main() hello,golang" + strconv.Itoa(i))
	// 	time.Sleep(time.Second)
	// }

	// cpuNum := runtime.NumCPU()
	// fmt.Println("cpuNum=", cpuNum)
	// runtime.GOMAXPROCS(cpuNum-1)

	// for i := 1; i <= 20; i++ {
	// 	go test1(i)
	// }

	// time.Sleep(time.Second)

//互斥锁
	// lock.Lock()
	// for i, v := range myMap {
	// 	fmt.Printf("map[%d]=%d\n", i, v)
	// }
	// lock.Unlock()

//管道
	//管道是引用类型,队列先进先出
	// var intChan chan int
	// intChan = make(chan int, 3)

	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// intChan <- 1
	// intChan <- 10
	// num := 21
	// intChan <- num

	// //管道不能自动增长
	// fmt.Printf("len= %v cap=%v \n", len(intChan), cap(intChan)) 

	// var num2 int
	// num2 = <- intChan
	// fmt.Println("num2=", num2) 

	// allChan := make(chan interface{}, 3)

	// allChan<- 10
	// allChan<- "tom jack"
	// cat := Cat{"小花猫", 4}
	// allChan<- cat

	// <-allChan
	// <-allChan

	// newCat := <-allChan //从管道中取出的Cat是什么?

	// fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	
	// //fmt.Printf("newCat.Name=%v", newCat.Name)//错误!
	
	// a := newCat.(Cat) 
	// fmt.Printf("newCat.Name=%v", a.Name)

//管道遍历和关闭
	// intChan := make(chan int,3)
	// intChan<- 100
	// intChan<- 101
	// close(intChan)
	// intChan<- 11//error
	// n1 := <-intChan
	
	// intChan2 := make(chan int, 100)
	// for i := 0; i < 10; i++ {
	// 	intChan2<- i * 2  
	// }

	// close(intChan2)//必须要先关闭再遍历
	// for v := range intChan2 {
	// 	fmt.Println("v=", v)
	// }

//协程与管道
	// intChan := make(chan int ,50)
	// exitChan := make(chan bool,1)
	// go writeData(intChan)
	// go readData(intChan,exitChan)
	// // time.Sleep(time.Second*10)
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }



	// intChan := make(chan int , 1000)
	// primeChan := make(chan int, 2000)
	// exitChan := make(chan bool, 4)

	// go putNum(intChan)
	// for i := 0; i < 4; i++ {
	// 	go primeNum(intChan, primeChan, exitChan)
	// }

	// go func(){
	// 	for i := 0; i < 4; i++ {
	// 		<-exitChan
	// 	}
	// 	close(primeChan)
	// }()
	
	// for {
	// 	res, ok := <-primeChan
	// 	if !ok{
	// 		break
	// 	}
	// 	fmt.Printf("素数=%d\n", res)
	// }

	// fmt.Println("main线程退出")

//管道可以声明为只读或者只写（默认双向）
	//只写
		// var chan2 chan<- int
		// chan2 = make(chan int, 10)
		// chan2<- 20
	//只读
		// var chan3 <-chan int
		// num := <-chan3

//select解决管道阻塞
	// intChan := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	intChan<- i
	// }
	// stringChan := make(chan string, 5)
	// for i := 0; i < 5; i++ {
	// 	stringChan <- "hello" + fmt.Sprintf("%d", i)
	// }

	// for{
	// 	select{
	// 		case v := <-intChan : 
	// 			fmt.Printf("从intChan读取的数据%d\n", v)
	// 			time.Sleep(time.Second)
	// 		case v := <-stringChan :
	// 			fmt.Printf("从stringChan读取的数据%s\n", v)
	// 			time.Sleep(time.Second)
	// 		default :
	// 			fmt.Printf("都取不到了\n")
	// 			time.Sleep(time.Second)
	// 			return 
	// 			//break label
	// 	}
	// }

//recover捕获panic
	go sayHello()
	go test3()


	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}
//协程、互斥锁、管道