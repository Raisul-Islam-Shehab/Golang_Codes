package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

func scanDirectory(dirName string) {
	files, err := ioutil.ReadDir(dirName)

	if err != nil {
		panic(err)
	}
	fmt.Println(dirName)

	for _, file := range files {
		filePath := filepath.Join(dirName, file.Name())

		if file.IsDir() {
			scanDirectory(filePath)

		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	dirName := "/home/dsi/GOLANG/Get_the_Files"
	defer reportPanic()
	scanDirectory(dirName)
}
