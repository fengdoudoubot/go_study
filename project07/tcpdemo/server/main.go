package main
import(
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()

	for{
		buf := make([]byte, 1024)
		fmt.Printf("等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n ,err := conn.Read(buf)//等待信息，阻塞
		if err != nil {
			fmt.Printf("客户端退出")
			return 
		}
		fmt.Print(string(buf[:n]))
	}
}

func main()  {
	fmt.Println("服务开始监听")
	listen ,err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return 
	}
	defer listen.Close()

	for{
		fmt.Println("等待客户端来链接....")
		conn ,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		}else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}