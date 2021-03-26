package main

import (
	"fmt"
	"os"
)

var paths []string = []string{
	"/server/project/dmq/DMQ/server/tx/header1/current",
	"/server/project/dmq/DMQ/server/tx/header1/finish_tx",
	"/server/project/dmq/DMQ/server/tx/header1/unfinish_tx",

	"/server/project/dmq/DMQ/server/tx/header2/current",
	"/server/project/dmq/DMQ/server/tx/header2/finish_tx",
	"/server/project/dmq/DMQ/server/tx/header2/unfinish_tx",

	"/server/project/dmq/DMQ/server/tx/header3/current",
	"/server/project/dmq/DMQ/server/tx/header3/finish_tx",
	"/server/project/dmq/DMQ/server/tx/header3/unfinish_tx",

	"/server/project/dmq/DMQ/server/data/header1/phone/data",
	"/server/project/dmq/DMQ/server/data/header1/phone/data",
	"/server/project/dmq/DMQ/server/data/header1/phone/data",

	"/server/project/dmq/DMQ/server/data/header2/phone/data",
	"/server/project/dmq/DMQ/server/data/header2/phone/data",
	"/server/project/dmq/DMQ/server/data/header2/phone/data",

	"/server/project/dmq/DMQ/server/data/header3/phone/data",
	"/server/project/dmq/DMQ/server/data/header3/phone/data",
	"/server/project/dmq/DMQ/server/data/header3/phone/data",

	"/server/project/dmq/DMQ/cachecenter/data",
}

func a() {
	for _, v := range paths {
		file, err := os.OpenFile(v, os.O_TRUNC|os.O_RDWR, 0755)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
	os.Remove("/server/project/dmq/DMQ/cachecenter/data")
}
func main() {
	a()
}
