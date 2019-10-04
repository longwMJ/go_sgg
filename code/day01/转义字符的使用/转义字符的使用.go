package main // 一个go文件需要在一个包
import "fmt"

// 运行 go run 转义字符的使用.go
func main() {
	// \t 排版
	fmt.Println("hello, \t go") // hello,   go

	// \n 回车
	fmt.Println("hello, \n go")
	// hello,
	// go

	// \"
	fmt.Println("hello, \" I Love You \" go ") // hello, " I Love You " go

	// \r 回车 当前行的 go字符串直接覆盖hello, 的前两个字符(go只有两个字符)
	fmt.Println("hello, \r go") // golo,

	// 练习
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")

	// 姓名    年龄    籍贯    住址
	// john    12      河北    北京

}
