package internal

import (
	"github.com/alexflint/go-arg"
	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Parser struct {
	ast *ast.File
}

func New(fileName string) (*Parser, error) {
	set := token.NewFileSet()
	f, err := parser.ParseFile(set, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, errors.Wrapf(err, "enum parser: Can't parse file with enum def '%s'", fileName)
	}
	return &Parser{ast: f}, nil
}

func (p Parser) Scan() (*reflection, error) {
	ref := reflection{}
	if err := p.parsePackage(&ref); err != nil {
		return nil, err
	}
	if err := p.parseEnum(&ref); err != nil {
		return nil, err
	}
	if err := p.parseEnumValues(&ref); err != nil {
		return nil, err
	}
	return &ref, nil
}

func (p Parser) parsePackage(ref *reflection) error {
	ast.Inspect(p.ast, func(node ast.Node) bool {
		p, ok := node.(*ast.File)
		if ok {
			ref.PackageName = p.Name.Name
			return false
		}
		return true
	})
	if ref.PackageName == "" {
		return errors.New("enum parser: Can't find package name")
	}
	return nil
}

func (p Parser) parseEnum(ref *reflection) error {
	ast.Inspect(p.ast, func(node ast.Node) bool {
		p, ok := node.(*ast.GenDecl)
		if ok && p.Tok == token.TYPE {
			generationCmd := strings.Replace(p.Doc.List[0].Text, "//go:generate ", "", 1)
			if !strings.Contains(generationCmd, "goenum") {
				return true
			}
			argParser, err := arg.NewParser(arg.Config{}, &ref.Opts)
			if err != nil {
				panic(err)
			}
			args, err := shellquote.Split(strings.Replace(generationCmd, "goenum", "", 1))
			if err != nil {
				panic(err)
			}
			err = argParser.Parse(args)
			if err != nil {
				panic(err)
			}
			s := p.Specs[0].(*ast.TypeSpec)
			ref.TypeName = s.Name.Name
			ref.BaseTypeName = s.Type.(*ast.Ident).Name
			return false
		}
		return true
	})
	if ref.TypeName == "" {
		return errors.New("enum parser: Can't find type")
	}

	if ref.BaseTypeName == "" {
		return errors.New("enum parser: Can't find base type name")
	}
	return nil
}

func (p Parser) parseEnumValues(ref *reflection) error {
	ast.Inspect(p.ast, func(node ast.Node) bool {
		p, ok := node.(*ast.File)
		if ok {
			for i, v := range p.Scope.Objects {
				decl, ok := v.Decl.(*ast.ValueSpec)
				if !ok {
					continue
				}
				if decl.Type.(*ast.Ident).Name != ref.TypeName {
					continue
				}
				ref.Values = append(ref.Values, reflectionEnumValue{Name: i, Value: decl.Values[0].(*ast.BasicLit).Value})
			}
			return false
		}
		return true
	})
	return nil
}
