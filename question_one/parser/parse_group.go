package parser

import (
	"fmt"
	"regexp"
)

func ParseGroup(group string, kmerLen int) []string {
	var kmers []string

	for i := 0; i < len(group)-kmerLen; i++ {
		subGroup := group[i:i+kmerLen]
		pattern := fmt.Sprintf("([ATCG]{%d})", kmerLen)
		match, _ := regexp.MatchString(pattern, subGroup)
		if match {
			kmers = append(kmers, subGroup)
		}
	}

	return kmers
}