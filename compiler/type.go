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
		return 1
	}
	if longType := dataType.LONG(); longType != nil {
		return 2
	}
	if shortType := dataType.SHORT(); shortType != nil {
		return 3
	}
	if byteType := dataType.BYTE(); byteType != nil {
		return 4
	}
	if boolType := dataType.BOOL(); boolType != nil {
		return 5
	}
	if stringType := dataType.STRING(); stringType != nil {
		return 6
	}
	if dateTypeType := dataType.DATETIME(); dateTypeType != nil {
		return 7
	}
	if unknownType := dataType.UNKNOWN(); unknownType != nil {
		return 8
	}
	if referenceType := dataType.IDENT(); referenceType != nil {
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
