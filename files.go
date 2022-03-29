package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) { // error return value no longer needed
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// instead of returning the value of the error value, pass it to "panic"
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			// no more need to store or check error return value
			scanDirectory(filePath)

		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	// no more need to store or check error return value
	scanDirectory(".")
}
