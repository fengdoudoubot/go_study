package main
import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main()  {
	conn ,err := net.Dial("tcp","172.31.156.235:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return 
	}
	// fmt.Println("conn succ=", conn)
	reader := bufio.NewReader(os.Stdin)
	line ,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err=", err)
	}
	n , err := conn.Write([]byte(line))
	fmt.Printf("seed=%v", n)

}