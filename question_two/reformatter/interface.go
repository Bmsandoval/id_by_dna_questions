package reformatter

type Format interface {
	ParseFile(string)([]Common, error)
}

var Reformatters = map[string]Format{}

func SelfRegister(suffix string, parser Format) {
	Reformatters[suffix]=parser
}

func GetReformatter(suffix string) *Format {
	if reformatter, ok := Reformatters[suffix]; ok {
		return &reformatter
	}
	return nil
}