package main //只要包含 main函数就是一个main包 就是主包，跟文件名没任何关系
import (
	"fmt"
	"time"
)

// main函数
func main() {
	fmt.Println("Hello world!!!")

	time.Sleep(1 * time.Second)
	fmt.Println("Hello world!!!")
}
