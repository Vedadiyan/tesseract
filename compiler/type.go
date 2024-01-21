package compiler

import (
	"fmt"
	"strconv"

	"github.com/vedadiyan/tesseract/parser"
)

type (
	Arg  any
	Call struct {
		Name string
		Args []Arg
	}
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
		args, err := GetArgs(call.AllArg())
		if err != nil {
			return nil, err
		}
		attributes[name] = args
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

func GetArg(arg parser.IArgContext) (Arg, error) {
	if call := arg.Call(); call != nil {
		args, err := GetArgs(call.AllArg())
		if err != nil {
			return nil, err
		}
		return &Call{
			Name: call.GetText(),
			Args: args,
		}, nil
	}
	if str := arg.EscapedString(); str != nil {
		return str.GetText(), nil
	}
	if number := arg.NUMBER(); number != nil {
		return strconv.ParseFloat(number.GetText(), 64)
	}
	if ident := arg.IDENT(); ident != nil {
		return Reference(ident.GetText()), nil
	}
	if data := arg.DATA(); data != nil {
		return nil, fmt.Errorf("$data is not valid in attributes")
	}
	return nil, fmt.Errorf("invalid argument")
}

func GetArgs(args []parser.IArgContext) ([]Arg, error) {
	arguments := make([]Arg, 0)
	for _, arg := range args {
		argument, err := GetArg(arg)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, argument)
	}
	return arguments, nil
}
