package mapper

func Insert(value string) error {
	var err error
	hash := Hash(value)
	bucketIndex := *hash & bucketMask
	bucket := &HashMap[bucketIndex]
	if len(*bucket) == 0 {
		bucket, err = Resize(bucket)
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
				link.Next = &MapEntry{
					Value: value,
					Count: 1,
					Next: nil,
				}
				break
			} else {
				link = link.Next
			}
		}
	}
	(*bucket)[index] = hashMapLocation

	HashMap[bucketIndex] = *bucket

	return nil
}