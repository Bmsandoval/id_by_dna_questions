package parser

import (
	"fmt"
	"regexp"
)

func ParseGroup(group string, kmerLen int) ([]string, error) {
	pattern := fmt.Sprintf("([ATCG]{%d})", kmerLen)
	pairRegex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	results := pairRegex.FindAllString(group, -1)

	return results, nil
}