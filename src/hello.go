// Hello, Go 语言!

// 表示该 go 代码所属的包
package main

// 导入 fmt 包，注意不得包含没有使用的包
import "fmt"

// main 函数不能带参数和返回值，传入的参数保存在 os.Args 中
// 如果需要支持命令行参数，可以用 flag 包
//
// 函数形式:
//
//   func 函数名(参数列表) (返回值列表) {
//            函数体 
//   }
//
func main() {
	// 不强制代码加分号
	fmt.Println("Hello, Go 语言!")
}

