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
		Fields      map[string]*Field
		PackageName string
	}
	KeyValue struct {
		Key   string
		Value any
	}
)

const (
	NONE TypeCode = iota
	INT
	LONG
	SHORT
	BYTE
	BOOL
	DATETIME
	STRING
	UNKNOWN
	TYPE
	LIST_OF_NONE TypeCode = 100 + iota
	LIST_OF_INT
	LIST_OF_LONG
	LIST_OF_SHORT
	LIST_OF_BYTE
	LIST_OF_BOOL
	LIST_OF_DATETIME
	LIST_OF_STRING
	LIST_OF_UNKNOWN
	LIST_OF_TYPE
)

func GetType(packageName string, typeDefinition parser.ITypeDefinitionContext) (*Type, error) {
	attributes, err := GetAttributes(typeDefinition.AllAttribute())
	if err != nil {
		return nil, err
	}
	fields := make(map[string]*Field)
	for _, field := range typeDefinition.AllField() {
		name, value, err := GetField(field)
		if err != nil {
			return nil, err
		}
		fields[name] = value
	}
	return &Type{
		Attributes: attributes,
		Fields:     fields,
	}, nil
}

func GetAttributes(attributeStatements []parser.IAttributeContext) (Attributes, error) {
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

func GetField(field parser.IFieldContext) (string, *Field, error) {
	attributes, err := GetAttributes(field.AllAttribute())
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
	if dataType := fieldType.DataType(); dataType != nil {
		return GetTypeCode(dataType)
	}
	if list := fieldType.List(); list != nil {
		typeCode := GetTypeCode(list.DataType())
		return 100 + typeCode
	}
	return 0
}

func GetTypeCode(dataType parser.IDataTypeContext) TypeCode {
	if intType := dataType.INT(); intType != nil {
		return INT
	}
	if longType := dataType.LONG(); longType != nil {
		return LONG
	}
	if shortType := dataType.SHORT(); shortType != nil {
		return SHORT
	}
	if byteType := dataType.BYTE(); byteType != nil {
		return BYTE
	}
	if boolType := dataType.BOOL(); boolType != nil {
		return BOOL
	}
	if stringType := dataType.STRING(); stringType != nil {
		return STRING
	}
	if dateTypeType := dataType.DATETIME(); dateTypeType != nil {
		return DATETIME
	}
	if unknownType := dataType.UNKNOWN(); unknownType != nil {
		return UNKNOWN
	}
	if referenceType := dataType.IDENT(); referenceType != nil {
		return TYPE
	}
	return NONE
}

func GetArg(arg parser.IArgContext) (Arg, error) {
	if call := arg.Call(); call != nil {
		args, err := GetArgs(call.AllArg())
		if err != nil {
			return nil, err
		}
		return &Call{
			Name: call.IDENT().GetText(),
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
