package compiler

import (
	"fmt"
	"strconv"

	"github.com/vedadiyan/tesseract/parser"
)

type (
	Arg        any
	Reference  string
	TypeCode   int
	Attributes map[string][]Arg
	Field      struct {
		TypeCode   TypeCode
		Attributes Attributes
	}
	Type struct {
		Attributes  Attributes
		Fields      map[string]Field
		PackageName string
	}
	KeyValue struct {
		Key   string
		Value any
	}
)

func NewType(packageName string, typeStatement parser.ITypeStatementContext) (*Type, error) {
	attributes, err := NewAttributes(typeStatement.AllAttribute())
	if err != nil {
		return nil, err
	}

	return &Type{
		Attributes: attributes,
	}, nil
}

func NewAttributes(attributeStatements []parser.IAttributeContext) (Attributes, error) {
	attributes := make(Attributes)
	for _, attribute := range attributeStatements {
		call := attribute.Call()
		name := call.IDENT().GetText()
		attributes[name] = make([]Arg, 0)
		for _, arg := range call.AllArg() {
			if assignment := arg.Assignment(); assignment != nil {
				continue
			}
			if call := arg.Call(); call != nil {
				continue
			}
			if str := arg.EscapedString(); str != nil {
				attributes[name] = append(attributes[name], str.GetText())
				continue
			}
			if number := arg.NUMBER(); number != nil {
				n, err := strconv.ParseFloat(number.GetText(), 64)
				if err != nil {
					return nil, err
				}
				attributes[name] = append(attributes[name], n)
				continue
			}
			if ident := arg.IDENT(); ident != nil {
				attributes[name] = append(attributes[name], Reference(ident.GetText()))
				continue
			}
			if data := arg.DATA(); data != nil {
				return nil, fmt.Errorf("$data is not valid in attributes")
			}
		}
	}
	return attributes, nil
}

func NewField(field parser.IFieldContext) (string, *Field, error) {
	attributes, err := NewAttributes(field.AllAttribute())
	if err != nil {
		return "", nil, err
	}
	typeCode := GetFieldType(field.FieldType())

	return field.IDENT().GetText(), &Field{
		Attributes: attributes,
		TypeCode:   typeCode,
	}, nil
}

func GetFieldType(fieldType parser.IFieldTypeContext) TypeCode {
	if type_ := fieldType.Type_(); type_ != nil {
		return GetTypeCode(type_)
	}
	if list := fieldType.List(); list != nil {
		typeCode := GetTypeCode(list.Type_())
		return 100 + typeCode
	}
	return 0
}

func GetTypeCode(type_ parser.ITypeContext) TypeCode {
	if intType := type_.INT(); intType != nil {
		return 1
	}
	if longType := type_.LONG(); longType != nil {
		return 2
	}
	if shortType := type_.SHORT(); shortType != nil {
		return 3
	}
	if byteType := type_.BYTE(); byteType != nil {
		return 4
	}
	if boolType := type_.BOOL(); boolType != nil {
		return 5
	}
	if stringType := type_.STRING(); stringType != nil {
		return 6
	}
	if dateTypeType := type_.DATETIME(); dateTypeType != nil {
		return 7
	}
	if unknownType := type_.UNKNOWN(); unknownType != nil {
		return 8
	}
	if referenceType := type_.IDENT(); referenceType != nil {
		return 9
	}
	return 0
}
