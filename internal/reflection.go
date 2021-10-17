package internal

import "strings"

// reflection enum representation
type reflection struct {
	PackageName string
	TypeName string
	// BaseTypeName (string, int, float, etc...)
	BaseTypeName string
	// Values of enum
	Values []reflectionEnumValue
	// Opts
	Opts options
	// Imports that will be added to generated code
	Imports []importPkg
}
// GetImport return existing import or add path to import if not present
func (r *reflection) GetImport(path string) importPkg {
	for _, v := range r.Imports {
		if v.Path == path {
			return v
		}
	}
	r.Imports = append(r.Imports, importPkg{Path: path})
	return r.Imports[len(r.Imports)-1]
}

type reflectionEnumValue struct {
	Name  string
	Value interface{}
}

// options for go:generate tool
type options struct {
	// Priority Generate methods to compare enums
	Priority bool `arg:"-p,--priority"`
}


type importPkg struct {
	Alias string
	Path  string
}

func (i importPkg) Name() string {
	if i.Alias != "" {
		return i.Alias
	}
	return i.SplitPath()[0]
}

func (i importPkg) SplitPath() []string {
	return strings.Split(i.Path, "/")
}
