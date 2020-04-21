package main

import (
	"fmt"
	"github.com/bmsandoval/question_one/mapper"
	"github.com/bmsandoval/question_one/parser"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a path to a file as a command line parameter")
		return
	}
	fileName := os.Args[1]

	fileContents, err := parser.ParseFile(fileName)

	if err != nil && err != io.EOF{
		fmt.Println(err.Error())
		return
	}
	if fileContents == nil || len(fileContents) == 0 {
		return
	}

	for _, content := range fileContents {
		kmers, err := parser.ParseGroup(content, 25)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, kmer := range kmers {
			if err := mapper.Insert(kmer); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	mapper.Stats()
}