package mapper

func Resize(bucket *[]*MapEntry) (*[]*MapEntry, error) {
	oldBucket := *bucket
	if oldBucket == nil || len(oldBucket) == 0 {
		newBucket := make([]*MapEntry, initialSize)
		return &newBucket, nil
	}

	newBucket := make([]*MapEntry, len(oldBucket)*2)
	// golang should do garbage collection after nilifying
	bucket = &newBucket

	return bucket, nil
}