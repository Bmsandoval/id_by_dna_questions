package mapper

const bucketCount = 64
const bucketMask = 63 // first 5 bits

const initialSize = 10000 // 1000 -> 10000 cuts time in half

type MapEntry struct {
	Value string
	Count int
	Next *MapEntry
}

var HashMap [bucketCount][]*MapEntry


