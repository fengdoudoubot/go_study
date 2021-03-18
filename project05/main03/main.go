package main

import(
	"fmt"
	"os"//file
	"bufio"
	"io"
	_"io/ioutil"
)

func CopyFile(dstFileName string, srcFileName string)(written int64, err error){
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return 
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main()  {
//读文件
	// file , err := os.Open("e:/test.txt")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	
	// fmt.Println(file)
	
	// err = file.Close()
	// if err != nil{
		// 	fmt.Println(err)
		// }
		
	// file , err := os.Open("e:/test.txt")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	
	// defer file.Close()

	// reader := bufio.NewReader(file)
	// for{
	// 	str ,err := reader.ReadString('\n')
	// 	if err == io.EOF{//io.EOF表示文件的末尾
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
	
	//一次性读取到位（适用于文件不大的情况）
	// file := "e:/test.txt"
	// content , err := ioutil.ReadFile(file)//open和close被封装在ReadFile文件内部
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// //fmt.Printf("%v",content)
	// fmt.Printf("%v",string(content))

//写文件
	//创建文件，写入内容
		// filePath := "e:/abc.txt"
		// file , err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE,0666)
		// if err != nil{
			// 	fmt.Print(err)
		// 	return 
		// }
		
		// defer file.Close()
		
		// str := "hello,ll\r\n"
		// writer := bufio.NewWriter(file)//writer是带缓存的
		// for i := 0; i < 5; i++ {
			// 	writer.WriteString(str)//内容是先写入到缓存的
		// }
		
		// writer.Flush()
	
	//覆盖文件内容
		// filePath := "e:/abc.txt"
		// file , err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC,0666)
		// if err != nil{
		// 		fmt.Print(err)
		// 	return 
		// }
		
		// defer file.Close()
		
		// str := "hello,www\r\n"
		// writer := bufio.NewWriter(file)
		// for i := 0; i < 5; i++ {
		// 		writer.WriteString(str)
		// }
		
		// writer.Flush()
		
	//追加内容
		// filePath := "e:/abc.txt"
		// file , err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND,0666)
		// if err != nil{
			// 		fmt.Print(err)
		// 	return 
		// }
		
		// defer file.Close()
		
		// str := "accc\r\n"
		// writer := bufio.NewWriter(file)
		// for i := 0; i < 2; i++ {
			// 		writer.WriteString(str)
		// }
		
		// writer.Flush()
		
	//读写
		// filePath := "e:/abc.txt"
		// file , err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND,0666)
		// if err != nil{
		// 		fmt.Print("err:",err)
		// 	return 
		// }
		
		// defer file.Close()
		
		// reader := bufio.NewReader(file)
		// for{
		// 	str,err := reader.ReadString('\n')
		// 	if err == io.EOF {
		// 		break 
		// 	}
		// 	fmt.Print(str)
		
		// }

		// fmt.Println()
		// str := ".1\r\n"
		// writer := bufio.NewWriter(file)
		// for i := 0; i < 5; i++ {
		// 		writer.WriteString(str)
		// }
		
		// writer.Flush()
	
	//将一个文件的内容导入到另一个文件
		// filePath1 := "e:/abc.txt"
		// filePath2 := "e:/bcd.txt"

		// data ,err := ioutil.ReadFile(filePath1)
		// if err != nil {
		// 	fmt.Printf("err=%v\n", err)
		// 	return
		// }
		
		// err = ioutil.WriteFile(filePath2,data,0666)
		// if err != nil {
		// 	fmt.Printf("write file error=%v\n", err)
		// }



//拷贝文件
	// srcFile := "e:/ttt.PNG"
	// dstFile := "e:/aaa.PNG"
	// _, err := CopyFile(dstFile, srcFile)
	// if err == nil {
	// 	fmt.Printf("拷贝完成\n")
	// } else {
	// 	fmt.Printf("拷贝错误 err=%v\n", err)
	// }

//

}
//文件