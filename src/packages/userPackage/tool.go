package userPackage

import (
	"fmt"
)

func init() {
	fmt.Println("tool init")
}

// 可以被导出的函数
func FuncPublic() {
	fmt.Println("FuncPublic")
}
