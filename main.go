package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/vedadiyan/tesseract/compiler"
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
        @JSON("field_name")
        @Validation("required")
        X []int
        @Name("field_name2")
        Y bool
        Z string
    }    
    MyType2 {
        @Name("field_name")
        A int 
        B []int64
        C []string
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
	x, y, err := compiler.CompilerTypes("test", p.AllTypeStatement())
	_ = x
	_ = y
	_ = err
	os.WriteFile("test.proto", []byte(x), os.ModePerm)
	os.WriteFile("test.json", []byte(y), os.ModePerm)
	//	_ = name
}
