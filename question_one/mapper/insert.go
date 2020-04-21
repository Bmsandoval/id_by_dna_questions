package mapper

func Insert(value string) error {
	var err error

	// Get the hash
	hash := Hash(value)

	// use a mask to get the right bucket
	bucketIndex := *hash & bucketMask
	bucket := &HashMap[bucketIndex]
	if len(*bucket) == 0 {
		bucket, err = Size(bucket)
		if err != nil {
			return err
		}
	}
	index := *hash % uint64(len(*bucket))
	hashMapLocation := (*bucket)[index]
	var nilEntry *MapEntry
	var link *MapEntry
	if hashMapLocation == nil {
		hashMapLocation = &MapEntry{
			Value: value,
			Count: 1,
		}
	} else {
		link = hashMapLocation
		for true {
			if link.Value == value {
				// equivalent k-mer, increment count
				link.Count++
				break
			} else if link.Next == nilEntry {
				// if we are at the end of the list, add new one
				link.Next = &MapEntry{
					Value: value,
					Count: 1,
					Next: nil,
				}
				break
			} else {
				// if not at end of list, visit next child
				link = link.Next
			}
		}
	}
	(*bucket)[index] = hashMapLocation

	HashMap[bucketIndex] = *bucket

	return nil
}