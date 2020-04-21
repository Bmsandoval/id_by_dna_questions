package mapper

func Hash(s string) *uint64 {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	hash := uint64(h)
	return &hash
}