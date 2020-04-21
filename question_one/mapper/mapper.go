package mapper

const bucketCount = 64
const bucketMask = 63 // first 5 bits

const initialSize = 100000

type MapEntry struct {
	Value string
	Count int
	Next *MapEntry
}

var HashMap [bucketCount][]*MapEntry


