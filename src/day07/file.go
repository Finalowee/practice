package day07

import (
	"fmt"
	"os"
)

var info os.FileInfo

func init() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("参数错误: cmd file")
		return
	}
	info, err := os.Stat(list[1])
	if nil != err {
		fmt.Println("文件不存在或没有访问权限:", list[1])
	}
	fmt.Println(info)
}
