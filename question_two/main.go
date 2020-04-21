package main

import (
	"fmt"
	"github.com/bmsandoval/question_two/formats"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a path to a file as a command line parameter")
		return
	}
	fileName := os.Args[1]

	ext := filepath.Ext(fileName)

	var err error

	var commonArray []formats.Common
	if ext == ".txt" {
		tf := formats.TextFormat{}
		commonArray, err = tf.ParseFile(fileName)
	} else if ext == ".json" {
		jf := formats.JsonFormat{}
		commonArray, err = jf.ParseFile(fileName)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//log.Println(commonArray)
	_ = commonArray
}