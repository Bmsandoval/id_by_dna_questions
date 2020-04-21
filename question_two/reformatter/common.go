package reformatter

type Common struct {
	Part string
	Number string
	PartLength int
}

type CommonFormat []Common

func (cf CommonFormat) Insert(common Common) {
	cf = append(cf, common)
}
