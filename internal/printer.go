package internal

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

const tmpl = `
{{ if .Ref.Opts.Priority }}
var priority{{ .Ref.TypeName }} = []{{ .Ref.TypeName }}{
	{{ range $k, $v := .Ref.Values }}
	   {{ $v.Name }},
	{{ end }}}
{{ end }}

func (e {{ .Ref.TypeName }}) String() string {
	{{ .Type.PrintString }}
}

func (e {{ .Ref.TypeName }}) Values() []interface{} {
	return []interface{}{
		{{ range $k, $v := .Ref.Values }}
	   	{{ $v.Name }},
		{{ end }}
	}
}

func (e {{ .Ref.TypeName }}) In(values ...{{ .Ref.TypeName }}) bool {
	for _, v := range values {
		if v == e {
			return true
		}
	}
	return false
}

{{ if .Ref.Opts.Priority }}
func (e {{ .Ref.TypeName }}) Priority() int {
	for k, v := range priority{{ .Ref.TypeName }} {
		if v == e {
			return k
		}
	}
	return -1
}

func (e {{ .Ref.TypeName }}) Compare(v {{ .Ref.TypeName }}) int {
	if e.Priority() > v.Priority() {
		return 1
	}
	if e.Priority() < v.Priority() {
		return -1
	}
	return 0
}

func (e {{ .Ref.TypeName }}) LessThan(v {{ .Ref.TypeName }}) bool {
	return e.Priority() < v.Priority()
}

func (e {{ .Ref.TypeName }}) GreaterThan(v {{ .Ref.TypeName }}) bool {
	return e.Priority() > v.Priority()
}
{{ end }}
`

// templatePrinter generate enum source based on template
type templatePrinter struct {
	Ref  *reflection
	Type Type
}

func NewTemplatePrinter(ref *reflection, t Type) *templatePrinter {
	return &templatePrinter{Ref: ref, Type: t}
}

func (p templatePrinter) Print(w io.Writer) error {
	wTmp := strings.Builder{}
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	if err = t.Execute(&wTmp, p); err != nil {
		return err
	}
	_, err = w.Write([]byte(fmt.Sprintf("%s\n\n%s\n\n%s", p.printPackage(), p.printImports(), wTmp.String())))
	return err
}

func (p templatePrinter) printPackage() string {
	return fmt.Sprintf("package %s", p.Ref.PackageName)
}

func (p templatePrinter) printImports() string {
	imports := p.Ref.Imports
	if len(imports) == 0 {
		return ""
	}
	if len(imports) == 1 {
		return fmt.Sprintf("import \"%s\"", imports[0].Path)
	}
	b := strings.Builder{}
	b.WriteString("import (")
	for _, v := range imports {
		if v.Alias == "" {
			b.WriteString(fmt.Sprintf("\n\"%s\"", v.Path))
			continue
		}
		b.WriteString(fmt.Sprintf("\n%s \"%s\"", v.Alias, v.Path))
	}
	b.WriteString(")")
	return b.String()
}
