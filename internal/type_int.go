package internal

import (
	"fmt"
)

type typeInt struct {
	Type
	ref *reflection
	def string
}

func (t typeInt) PrintString() string {
	return fmt.Sprintf("return %s.Itoa(%s(e))", t.ref.GetImport("strconv").Name(), t.def)
}