package formats

type Format interface {
	ParseFile(string)string
}