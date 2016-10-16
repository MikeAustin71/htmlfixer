package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileReadStatus is a data structure used to track file
// processing status.
type FileReadStatus struct {
	isInputFileExist	bool
	isHTMLTag        	bool
	isStyleTag			bool
	isBodyTag        	bool
	inputFile        	FileData
	tempFile         	FileData
}

// FileData is used to track file characteristics
type FileData struct {
	absPathAndFileName 	string
	dir                	string
	fileNameAndExt     	string
	ext                	string
	fileName        	string
	volume				string
	fileSize           	int64
}

func main() {

	inFile, err := os.Open("./FileList.txt")

	defer inFile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	files := make([]string, 0)

	for scanner.Scan() {

		filePathNameLine := scanner.Text()
		filePathNameLine = strings.Trim(filePathNameLine," ")

		if filePathNameLine == "" {
			continue
		}

		files = append(files, filePathNameLine)
	}

	for _, s := range files {

		fInfo := NewFileStatusInfo(s)

		processHTMLFile(&fInfo)
	}

	fmt.Println("Program Completed!")

}

// NewFileStatusInfo creates and initializes a FileReadStatus structure.
func NewFileStatusInfo(filePathName string) FileReadStatus {

	fPath, err := filepath.Abs(filePathName)

	if err != nil {
		fmt.Println("NewFileStatusInfo(): Error filePathName =", filePathName )
		panic(err)
	}

	var status = FileReadStatus {
		isBodyTag: false,
		isHTMLTag: false,
		isStyleTag: false,
	}

	status.inputFile.absPathAndFileName = fPath
	// no trailing slash on this dir string
	status.inputFile.dir = filepath.Dir(fPath)
	ext := filepath.Ext(fPath)
	status.inputFile.ext = ext
	fNameExt := filepath.Base(fPath)
	status.inputFile.fileNameAndExt = fNameExt
	status.inputFile.fileName = strings.TrimRight(fNameExt, ext)
	status.inputFile.volume = filepath.VolumeName(fPath)

	status.tempFile.dir = status.inputFile.dir
	status.tempFile.ext = ".tmp"
	status.tempFile.fileName = status.inputFile.fileName
	status.tempFile.volume = status.inputFile.volume
	status.tempFile.fileNameAndExt =
		status.inputFile.fileName + status.tempFile.ext
	status.tempFile.absPathAndFileName =
		status.inputFile.dir + string(os.PathSeparator) + status.tempFile.fileNameAndExt

	return status
}

func processHTMLFile(fInfo *FileReadStatus) {


	inFile, err := os.Open(fInfo.inputFile.absPathAndFileName)

	defer inFile.Close()

	if err != nil {
		fmt.Println("File Open Error File = ", fInfo.inputFile.absPathAndFileName)
		panic(err)
	}

	outFile, err := os.Create(fInfo.tempFile.absPathAndFileName)

	defer outFile.Close()

	if err != nil {
		fmt.Println("Error Creating Output file: ", fInfo.tempFile.absPathAndFileName)
		panic(err)
	}

	writer := bufio.NewWriter(outFile)

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		str := scanner.Text()

		if strings.Contains(str, "<style>") {
			fInfo.isStyleTag = true
		}

		if strings.Contains(str, "</style>"){
			fInfo.isStyleTag = false
		}

		if strings.Contains(str, "<html>"){
			fInfo.isHTMLTag = true
		}

		if strings.Contains(str, "</html>") {
			fInfo.isHTMLTag = false
		}

		if strings.Contains(str, "<body") {
			fInfo.isBodyTag = true
		}

		if strings.Contains(str, "</body>") {
			fInfo.isBodyTag = false
		}

		if fInfo.isStyleTag && strings.Contains(str, "mso-style-"){
			continue
		}

		if strings.Contains(str, "windowtext"){
			str = strings.Replace(str,"windowtext", "black", -1)
		}

		if fInfo.isBodyTag && strings.Contains(str, "name=\"_"){
			str = strings.Replace(str,"name=\"_", "id=\"a_", -1)
		}

		if fInfo.isBodyTag && strings.Contains(str, "#_"){
			str = strings.Replace(str,"#_", "#a_", -1)
		}

		fmt.Fprintln(writer, str)
	}

	writer.Flush()
}
