package day02

import "os"

func GetParam() (list []string) {
	list = os.Args
	return
}
