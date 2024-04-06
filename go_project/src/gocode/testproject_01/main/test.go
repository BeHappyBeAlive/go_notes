package main  //声明文件所在的包，每个go文件必须有归属的包
import "fmt"  //引入程序中需要用到的包，为了使用这个包下的函数 比如Printlin函数
func main(){  //程序入口主函数
	fmt.Println("Hello Golang!") //打印控制台
}