package internal

import (
	"fmt"
)

type typeFloat struct {
	Type
	ref *reflection
	def string
}

func (t typeFloat) PrintString() string {
	return fmt.Sprintf("return %s.Sprntf(\"%%v\", %s(e))", t.ref.GetImport("fmt").Name(), t.def)
}