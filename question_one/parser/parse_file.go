package parser

import (
	"bufio"
	"bytes"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text[0] == '>' {
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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
