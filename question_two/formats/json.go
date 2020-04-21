package formats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type JsonFormat struct {}

type JsonFile struct {
	Count int
	Parts []string
	Numbers []int
}

func (jf JsonFormat) ParseFile(filePath string) ([]Common, error) {
	var commonArray []Common
	file, _ := ioutil.ReadFile(filePath)

	data := JsonFile{}
	_ = json.Unmarshal(file, &data)

	output := ""
	xError := false
	for i := 0; i < len(data.Numbers); i++ {
		output += fmt.Sprintf("%s:%d ", data.Parts[i], data.Numbers[i])
		testX := strings.ToLower(data.Parts[i])
		if testX == "x" {
			xError = true
		}

		commonArray = append(commonArray, Common{
			Part:       data.Parts[i],
			Number:     string(data.Numbers[i]),
			PartLength: len(data.Parts[i]),
		})
	}

	if xError {
		output = strings.TrimSpace(output)
		output += ", error found an X"
	}
	fmt.Println(output)

	return commonArray, nil
}
