package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileReadStatus struct {
	isBody bool
	fileName string
	fileSize int64
}

var fileStatus FileReadStatus

var inputFile = "D:/go/work/src/github.com/htmlfixer/src/htmlfixer/006870_ReadingFiles.htm"

// var outputFile string = "D:/go/work/src/github.com/htmlfixer/src/htmlfixer/output.htm"

func main() {

	fileStatus.isBody = false

	inFile, err := os.Open("./006870_ReadingFiles.htm")

	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
