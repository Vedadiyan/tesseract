package compiler

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"text/template"

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
		TypeName   string
		Attributes Attributes
		Index      int
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

var (
	//go:embed type.go.tmpl
	_typeTemplate string
)

func CompilerTypes(packageName string, typeStatements []parser.ITypeStatementContext) (string, error) {
	types, err := GetTypes(packageName, typeStatements)
	if err != nil {
		return "", err
	}
	template, err := template.New("types").Parse(_typeTemplate)
	if err != nil {
		return "", err
	}
	output := bytes.NewBufferString("")
	err = template.Execute(output, struct {
		PackageName string
		Types       map[string]*Type
	}{PackageName: packageName, Types: types})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func GetTypes(packageName string, typeStatements []parser.ITypeStatementContext) (map[string]*Type, error) {
	types := make(map[string]*Type)
	for _, typeStatement := range typeStatements {
		for _, typeDefinition := range typeStatement.AllTypeDefinition() {
			_type, err := GetType(packageName, typeDefinition)
			if err != nil {
				return nil, err
			}
			types[typeDefinition.IDENT().GetText()] = _type
		}
	}
	return types, nil
}

func GetType(packageName string, typeDefinition parser.ITypeDefinitionContext) (*Type, error) {
	attributes, err := GetAttributes(typeDefinition.AllAttribute())
	if err != nil {
		return nil, err
	}
	fields := make(map[string]*Field)
	for index, field := range typeDefinition.AllField() {
		name, value, err := GetField(index, field)
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

func GetField(index int, field parser.IFieldContext) (string, *Field, error) {
	attributes, err := GetAttributes(field.AllAttribute())
	if err != nil {
		return "", nil, err
	}
	typeCode, typename := GetFieldType(field.FieldType())

	return field.IDENT().GetText(), &Field{
		Attributes: attributes,
		TypeCode:   typeCode,
		TypeName:   typename,
		Index:      index,
	}, nil
}

func GetFieldType(fieldType parser.IFieldTypeContext) (TypeCode, string) {
	if dataType := fieldType.DataType(); dataType != nil {
		return GetTypeCode(dataType)
	}
	if list := fieldType.List(); list != nil {
		typeCode, typeName := GetTypeCode(list.DataType())
		return 100 + typeCode, fmt.Sprintf("repated %s", typeName)
	}
	return 0, ""
}

func GetTypeCode(dataType parser.IDataTypeContext) (TypeCode, string) {
	if intType := dataType.INT(); intType != nil {
		return INT, "int32"
	}
	if longType := dataType.LONG(); longType != nil {
		return LONG, "int64"
	}
	if shortType := dataType.SHORT(); shortType != nil {
		return SHORT, "int32"
	}
	if byteType := dataType.BYTE(); byteType != nil {
		return BYTE, "int32"
	}
	if boolType := dataType.BOOL(); boolType != nil {
		return BOOL, "bool"
	}
	if stringType := dataType.STRING(); stringType != nil {
		return STRING, "string"
	}
	if dateTypeType := dataType.DATETIME(); dateTypeType != nil {
		return DATETIME, "string"
	}
	if unknownType := dataType.UNKNOWN(); unknownType != nil {
		return UNKNOWN, "Struct"
	}
	if referenceType := dataType.IDENT(); referenceType != nil {
		return TYPE, referenceType.GetText()
	}
	return NONE, ""
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
