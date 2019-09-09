package day05

import (
	"fmt"
	"io"
	"os"
)

func Copy(src string, dst string) {
	_is, err := os.Open(src)
	if !IsExist(src) || nil != err {
		panic("未找到文件：" + src)
	}
	defer _is.Close()
	if IsExist(dst) || nil != err {
		panic("文件已存在：" + dst)
	}
	_os, err := os.Create(dst)
	defer _os.Close()
	var _tmp = make([]byte, 2048)
	for {
		c, err := _is.Read(_tmp)
		if err == io.EOF && c == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		_os.Write(_tmp)
	}
}

func IsExist(file string) bool {
	_, err := os.Stat(file)
	return nil == err || os.IsExist(err)
}
