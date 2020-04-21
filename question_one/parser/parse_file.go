package parser

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func ParseFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var groups []string
	var group *bytes.Buffer

	chunkCount := 0
	reader := bufio.NewReader(file)
	for true {
		lineBytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil{
			return nil, err
		}
		if lineBytes == nil {
			return nil, nil
		}
		var line string
		line = string(lineBytes)

		text := strings.TrimSpace(line)
		if text[0] == '>' {
			chunkCount++
			if group != nil {
				groups = append(groups, group.String())
			}
			group = &bytes.Buffer{}
		} else {
			group.WriteString(text)
		}
	}
	if group != nil {
		groups = append(groups, group.String())
	}

	return groups, nil
}
