package mapper

import (
	"fmt"
)

func Stats() {
	distinct := 0
	total := 0
	var best *MapEntry
	var nilEntry *MapEntry
	// for each bucket
	for _, bucket := range HashMap {
		// walk each entry
		for _, entry := range bucket {
			if entry == nilEntry {
				continue
			}
			distinct++
			if best == nil || entry.Count > best.Count {
				best = entry
			}
			total += entry.Count

			// and visit all children
			item := entry
			for item.Next != nil {
				item = item.Next
				distinct++
				if best == nil || item.Count > best.Count {
					best =item
				}
				total += item.Count
			}
		}
	}

	if best != nil {
		fmt.Printf("Count distinct k-mers: %d\n", distinct)
		fmt.Printf("Total count k-mers: %d\n", total)
		fmt.Printf("Most occuring k-mer: %s\n", best.Value)
	}
}