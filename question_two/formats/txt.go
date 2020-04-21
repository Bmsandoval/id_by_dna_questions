package formats

import (
	"bufio"
	"fmt"
	"github.com/bmsandoval/question_two/reformatter"
	"log"
	"os"
	"regexp"
	"strings"
)

func init() {
	reformatter.SelfRegister(".txt", TextFormat{})
}

type TextFormat struct {}

func (tf TextFormat) ParseFile(filePath string) ([]reformatter.Common, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var commonArray []reformatter.Common

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xError := false
		text := scanner.Text()
		pairRegex, err := regexp.Compile("[a-zA-Z]+|[0-9]+")
		if err != nil {
			return nil, err
		}

		results := pairRegex.FindAllString(text, -1)

		output := ""
		for i,j := 0,1; j < len(results); i,j=i+2,j+2 {
			output += fmt.Sprintf("%s:%s ", results[i], results[j])
			testX := strings.ToLower(string(results[i][0]))
			if testX == "x" {
				xError = true
			}

			commonArray = append(commonArray, reformatter.Common{
				Part:       results[i],
				Number:     results[j],
				PartLength: len(results[i]),
			})
		}

		if xError {
			output = strings.TrimSpace(output)
			output += ", error found an X"
		}
		fmt.Println(output)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commonArray, nil
}
