package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileReadStatus struct {
	isInputFileExist bool
	isHtmlTag bool
	isBodyTag bool
	inputFile FileData
	tempFile FileData
}

type FileData struct {
	absPathAndFileName string
	dir string
	fileNameandExt string
	ext string
	fileName string
	fileSize int64
}


var fileCount int

//var inputFile = "D:/go/work/src/github.com/htmlfixer/src/htmlfixer/006870_ReadingFiles.htm"

// var outputFile string = "D:/go/work/src/github.com/htmlfixer/src/htmlfixer/output.htm"

func main() {

	inFile, err := os.Open("./FileList.txt")

	defer inFile.Close()

	if err != nil {
		panic(err)
	}



	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		filePathNameLine := scanner.Text()

		filePathName := strings.TrimRight(filePathNameLine, "\n")




		fileCount++

		fmt.Println("File Count ", fileCount)

		ProcessHtmlFile(filePathName)

	}

}

func CreateFileStatusInfo(filePathName string) FileReadStatus {

	var status = new(FileReadStatus)

	status.isBodyTag = false
	status.isHtmlTag = false

	fPath, err := filepath.Abs(filePathName)

	if err != nil {
		panic(err)
	}

	status.inputFile.absPathAndFileName = fPath
	status.inputFile.dir = filepath.Dir(fPath)
	ext:= filepath.Ext(fPath)
	status.inputFile.ext = ext
	fNameExt := filepath.Base(fPath)
	status.inputFile.fileNameandExt = fNameExt
	status.inputFile.fileName = strings.TrimRight(fNameExt, ext)

	return status
}

func ProcessHtmlFile(filePathName string)  {

	fPathInfo := CreateFileStatusInfo(filePathName)

	inFile, err := os.Open(fPathInfo.inputFile.absPathAndFileName)

	defer inFile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		fmt.Println(scanner.Text())

	}

}
