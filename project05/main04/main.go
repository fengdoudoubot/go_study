package main

import (
	"fmt"
	_"os"
	"flag"
)

type Monster struct {
	Name string `json:"monster_name"` //反射机制
	Age int `json:"monster_age"`
	Birthday string //....
	Sal float64
	Skill string
}

func testStruct() {
	//演示
	monster := Monster{
		Name :"牛魔王",
		Age : 500 ,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster) //..
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))

}

func main(){
//命令行参数
	// fmt.Println("args:",len(os.Args))
	// for i,v := range os.Args{
	// 	fmt.Printf("args[%v]=%v\n", i, v)
	// }

	// var user string
	// var pwd string
	// var host string
	// var port int

	// flag.StringVar(&user, "u", "", "默认为空")
	// flag.StringVar(&pwd, "pwd", "", "默认为空")
	// flag.StringVar(&host, "h", "localhost", "默认为localhost")
	// flag.IntVar(&port, "port", 3306, "默认为3306")
	
	// flag.Parse()

	// fmt.Printf("user=%v pwd=%v host=%v port=%v", 
	// 	user, pwd, host, port)



}
//命令行参数