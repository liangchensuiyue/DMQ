package main

import (
	"fmt"
	"os"
)

var paths []string = []string{
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header1\\current",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header1\\finish_tx",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header1\\unfinish_tx",

	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header2\\current",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header2\\finish_tx",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header2\\unfinish_tx",

	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header3\\current",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header3\\finish_tx",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\tx\\header3\\unfinish_tx",

	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header1\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header1\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header1\\phone\\data",

	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header2\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header2\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header2\\phone\\data",

	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header3\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header3\\phone\\data",
	"D:\\gaodongsheng\\goproject\\src\\go_code\\DMQ\\DMQ\\server\\data\\header3\\phone\\data",
}

func main() {
	for _, v := range paths {
		file, err := os.OpenFile(v, os.O_TRUNC|os.O_RDWR, 0755)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
}
