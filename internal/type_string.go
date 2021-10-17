package internal

type TypeString struct {
	Type
	ref *reflection
}

func (t TypeString) PrintString() string {
	return "return string(e)"
}
