package main

import (
	"fmt"
	"time"
)



// func B() []func() {
// 	b := make([]func(), 3, 3)
// 	for i := 0; i < 3; i++ {
// 		b[i] = func() {
// 			fmt.Println(i)
// 		}
// 	}
// 	return b
// 	}//输出3，3，3
	
// func B() []func() {
// 	b := make([]func(), 3, 3)
// 	j := 0
// 	for i := 0; i < 3; i++ {
// 		j = i
// 		b[i] = func() {
// 			fmt.Println(j)
// 		}
// 	}
// 	return b
// }//输出0，1，2
	
// func main() {
// 	c := B()
// 	c[0]()
// 	c[1]()
// 	c[2]()
// }





// func B() []func() {
// 	b := make([]func(), 3, 3)
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(i)
// 		b[i] = func() {
// 			fmt.Println(i)
// 		}
// 	}
// 	return b
// 	}//输出3，3，3
	
func B() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		j := i
		fmt.Printf("%v,i,%v\n",&i,i)
		fmt.Printf("%v,j,%v\n",&j,j)
		b[i] = func() {//此处相当于j和函数成了闭包？
			fmt.Printf("%v,i,%v\n",&i,i)
			fmt.Printf("%v,j,%v\n",&j,j)
		}
	}
	return b
}//输出0，1，2
	
func main() {
	c := B()
	fmt.Println("sssss")
	c[0]()
	fmt.Println("0000")
	time.Sleep(time.Second * 3)
	c[1]()
	fmt.Println("1111")
	time.Sleep(time.Second * 3)
	c[2]()
	fmt.Println("222")
	time.Sleep(time.Second * 3)

	// for i:=0;i<5;i++{
    //     go func(x int){
    //         fmt.Println(x)
    //     }(i)

	// }
	// for i:=0; i<5; i++ {
	// 	go func(){
	// 		fmt.Println(i) //i变量值也是引用.创建5个线程执行函数， for循环执行过程中可能执行完的时候，线程刚好处于i的某个值。 
	// 	}()
		
	// }
	time.Sleep(time.Second * 1)
}




