package entry

import (
	"fmt"
	"github.com/bmsandoval/question_two/reformatter"
	"os"
	"path/filepath"
)

// Entry is required so the formats can self-register
func Entry() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a path to a file as a command line parameter")
		return
	}
	fileName := os.Args[1]
	ext := filepath.Ext(fileName)

	// reformatters adhere to an interface, get your ext specific one
	rf := reformatter.GetReformatter(ext)
	if rf == nil {
		fmt.Println("not configured for this file type")
		return
	}

	// now just parse the file given the parser/reformatter you got
	commonArray, err := (*rf).ParseFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//log.Println(commonArray)
	_ = commonArray
}