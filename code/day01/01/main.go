package main // 一个go文件需要在一个包
import "fmt"

// // 指针的使用
// func testPtr(num *int) {
// 	*num = 20
// }

// 运行go build main.go, 生成main.exe 文件
// 运行go build -o newName.exe main.go, 生成newName.exe 文件
// 运行go run main.go, 直接运行main.go 文件
// 运行gofmt -w main.go 直接格式main.go 文件
func main() {
	fmt.Println("hello, go")

}
