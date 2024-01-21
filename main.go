package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/vedadiyan/tesseract/parser"
)

const program = `
package Test;
import "somefile.h";
use Something;
use Something;

const (
    test = ""
)

type (
    MyType {
        @Name(Test("ok"),"field_name")
        @Test("field_name")
        field []int, 
        @Name("field_name2")
        field2 otherType,
        field2 otherType
    }    
    MyType2 {
        @Name("field_name")
        field int, 
        field2 otherType,
        field2 otherType
	}
)

@Prefix("internal.test")
service (
    @Subject("Message")
    @Balanced()
    @Cached(TTL(3600))
    SendMessage(@Mapper("filename") MyType) @Mapper("filename") MyType;
)

@Package("github.com/vedadiyan/test")
backend (
    SendMessage => HttpPost(URL(""), Body($data), Header("", ""), Authorization(Env("")))
)

gateway (
    @Route("")
    @Services(SendMessage)
    get GetMessage();
)
`

type customErrorListener struct {
	*antlr.DefaultErrorListener
}

func (l *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("Parser Error at line %d, column %d: %s\n", line, column, msg))
}

func main() {
	input := antlr.NewInputStream(program)
	lexer := parser.NewtesseractLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	test := parser.NewtesseractParser(stream)
	test.AddErrorListener(&customErrorListener{})
	//	name := test.PackageStatement().IDENT().GetText()
	p := test.Program()
	var c any
	c = p.AllTypeStatement()[0].AllField()[0].Attribute(0).Call().AllArg()[0].EscapedString()
	c = p.AllTypeStatement()[0].AllField()[0].Attribute(0).Call().AllArg()[0].Call()
	x := p.AllImportStatement()[0].GetText()
	_ = c
	_ = test
	_ = x
	//	_ = name
}
