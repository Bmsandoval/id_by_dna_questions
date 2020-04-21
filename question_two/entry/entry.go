package entry

import (
	"fmt"
	"github.com/bmsandoval/question_two/reformatter"
	"os"
	"path/filepath"
)

func Entry() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a path to a file as a command line parameter")
		return
	}
	fileName := os.Args[1]

	ext := filepath.Ext(fileName)

	var err error

	rf := reformatter.GetReformatter(ext)
	if rf == nil {
		fmt.Println("not configured for this file type")
		return
	}

	commonArray, err := (*rf).ParseFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//log.Println(commonArray)
	_ = commonArray
}