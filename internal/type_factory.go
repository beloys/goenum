package internal

import (
	"github.com/pkg/errors"
)

var ErrTypeFactoryUnknownType = errors.New("enum type factory: Unknown type")

type typeFactory struct {
}

func NewTypeFactory() *typeFactory {
	return &typeFactory{}
}

func (f typeFactory) Create(ref *reflection) (Type, error) {
	switch ref.BaseTypeName {
	case "string":
		return &TypeString{ref: ref}, nil
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return &typeInt{ref: ref, def: ref.BaseTypeName}, nil
	case "float", "float32", "float64":
		return &typeInt{ref: ref, def: ref.BaseTypeName}, nil
	}
	return nil, errors.Wrap(ErrTypeFactoryUnknownType, ref.TypeName)
}
