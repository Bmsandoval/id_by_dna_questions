package main

import (
	"fmt"
	"github.com/bmsandoval/question_one/mapper"
	"github.com/bmsandoval/question_one/parser"
	"io"
	"log"
)

func main() {
	fileContents, err := parser.ParseFile("SRR1748776.fa")

	if err != nil && err != io.EOF{
		fmt.Println(err.Error())
		return
	}
	if fileContents == nil || len(fileContents) == 0 {
		return
	}

	log.Println(len(fileContents))

	count := 0
	for _, content := range fileContents {
		kmers := parser.ParseGroup(content, 25)
		for _, kmer := range kmers {
			if err := mapper.Insert(kmer); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if count % 50 == 0  {
			log.Println(count)
		}
		count++
	}

	mapper.Stats()
	fmt.Println("completed without errors")
}