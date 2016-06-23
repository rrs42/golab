package main

import "os"
import "fmt"

func main() {
	const (
		path = "/sys"
	)

	fd, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}

	fileinfo, err := fd.Readdir(0)

	for _, fi := range fileinfo {
		fmt.Println(fi.Name())
	}

}
