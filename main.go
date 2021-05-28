package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	aSecret := os.Getenv("A_SECRET")
	fmt.Println(aSecret)
	fmt.Println(reverseString(aSecret))
}

func reverseString(s string) string {
	r := []rune(s)
	sb := strings.Builder{}
	for i := len(r) - 1; i >= 0; i-- {
		sb.WriteRune(r[i])
	}
	return sb.String()
}

作者：TheWinds
链接：https://juejin.cn/post/6844903957391736840
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
