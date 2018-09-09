package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTreeInternal(out, path, printFiles, "")
}

func dirTreeInternal(out io.Writer, path string, printFiles bool, padding string) error {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error during reading file...", err)
	}

	filesInfo, err := file.Readdir(0)

	if err != nil {
		fmt.Println("error during reading directory", err)
	}

	filteredFilesInfo := make([]os.FileInfo, 0)

	if printFiles {
		filteredFilesInfo = filesInfo
	} else {
		for _, fileInfo := range filesInfo {
			if fileInfo.IsDir() {
				filteredFilesInfo = append(filteredFilesInfo, fileInfo)
			}
		}
	}

	length := len(filteredFilesInfo)

	for idx, fileInfo := range filteredFilesInfo {
		isLast := idx == length-1
		var symbol string

		if isLast {
			symbol = "└"
		} else {
			symbol = "├"
		}

		if fileInfo.IsDir() {
			fmt.Fprintf(out, padding+symbol+"───"+fileInfo.Name()+"\n")
			var newPadding string
			if isLast {
				newPadding = padding + "\t"
			} else {
				newPadding = padding + "│\t"
			}
			dirTreeInternal(out, path+"/"+fileInfo.Name(), printFiles, newPadding)
		} else {
			var size string
			if fileInfo.Size() == 0 {
				size = "(empty)"
			} else {
				size = "(" + strconv.FormatInt(fileInfo.Size(), 10) + "b)"
			}
			fmt.Fprintf(out, padding+symbol+"───"+fileInfo.Name()+" "+size+"\n")
		}
	}

	return nil
}
