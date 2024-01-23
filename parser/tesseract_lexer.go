// Code generated from ./grammar/tesseract.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type tesseractLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TesseractLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tesseractlexerLexerInit() {
	staticData := &TesseractLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'[]'", "','", "", "'import'", "'use'", "'package'", "'type'", "'const'",
		"'service'", "'gateway'", "'backend'", "'int'", "'int64'", "'int16'",
		"'byte'", "'bool'", "'string'", "'time'", "'struct'", "'get'", "'post'",
		"'put'", "'patch'", "'delete'", "'$data'", "'='", "':'", "'\"'", "'>'",
		"'<'", "'('", "')'", "'{'", "'}'", "';'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "EscapedString", "IMPORT", "USE", "PACKAGE", "TYPE", "CONST",
		"SERVICE", "GATEWAY", "BACKEND", "INT", "LONG", "SHORT", "BYTE", "BOOL",
		"STRING", "DATETIME", "UNKNOWN", "GET", "POST", "PUT", "PATCH", "DELETE",
		"DATA", "EQ", "COLON", "DOUBLE_QUOTE", "GT", "LT", "LP", "RP", "LCB",
		"RCB", "SEMI", "AT", "IDENT", "NUMBER", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "EscapedString", "IMPORT", "USE", "PACKAGE", "TYPE",
		"CONST", "SERVICE", "GATEWAY", "BACKEND", "INT", "LONG", "SHORT", "BYTE",
		"BOOL", "STRING", "DATETIME", "UNKNOWN", "GET", "POST", "PUT", "PATCH",
		"DELETE", "DATA", "EQ", "COLON", "DOUBLE_QUOTE", "GT", "LT", "LP", "RP",
		"LCB", "RCB", "SEMI", "AT", "IDENT", "NUMBER", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 292, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 91, 8, 2, 10, 2, 12, 2, 94, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 36, 4, 36, 254, 8, 36, 11, 36, 12, 36, 255,
		5, 36, 258, 8, 36, 10, 36, 12, 36, 261, 9, 36, 1, 37, 4, 37, 264, 8, 37,
		11, 37, 12, 37, 265, 1, 37, 1, 37, 4, 37, 270, 8, 37, 11, 37, 12, 37, 271,
		5, 37, 274, 8, 37, 10, 37, 12, 37, 277, 9, 37, 1, 38, 4, 38, 280, 8, 38,
		11, 38, 12, 38, 281, 1, 38, 1, 38, 1, 39, 4, 39, 287, 8, 39, 11, 39, 12,
		39, 288, 1, 39, 1, 39, 0, 0, 40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 1, 0, 6, 3, 0, 10, 10,
		13, 13, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32,
		32, 301, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 84, 1, 0, 0,
		0, 5, 86, 1, 0, 0, 0, 7, 97, 1, 0, 0, 0, 9, 104, 1, 0, 0, 0, 11, 108, 1,
		0, 0, 0, 13, 116, 1, 0, 0, 0, 15, 121, 1, 0, 0, 0, 17, 127, 1, 0, 0, 0,
		19, 135, 1, 0, 0, 0, 21, 143, 1, 0, 0, 0, 23, 151, 1, 0, 0, 0, 25, 155,
		1, 0, 0, 0, 27, 161, 1, 0, 0, 0, 29, 167, 1, 0, 0, 0, 31, 172, 1, 0, 0,
		0, 33, 177, 1, 0, 0, 0, 35, 184, 1, 0, 0, 0, 37, 189, 1, 0, 0, 0, 39, 196,
		1, 0, 0, 0, 41, 200, 1, 0, 0, 0, 43, 205, 1, 0, 0, 0, 45, 209, 1, 0, 0,
		0, 47, 215, 1, 0, 0, 0, 49, 222, 1, 0, 0, 0, 51, 228, 1, 0, 0, 0, 53, 230,
		1, 0, 0, 0, 55, 232, 1, 0, 0, 0, 57, 234, 1, 0, 0, 0, 59, 236, 1, 0, 0,
		0, 61, 238, 1, 0, 0, 0, 63, 240, 1, 0, 0, 0, 65, 242, 1, 0, 0, 0, 67, 244,
		1, 0, 0, 0, 69, 246, 1, 0, 0, 0, 71, 248, 1, 0, 0, 0, 73, 250, 1, 0, 0,
		0, 75, 263, 1, 0, 0, 0, 77, 279, 1, 0, 0, 0, 79, 286, 1, 0, 0, 0, 81, 82,
		5, 91, 0, 0, 82, 83, 5, 93, 0, 0, 83, 2, 1, 0, 0, 0, 84, 85, 5, 44, 0,
		0, 85, 4, 1, 0, 0, 0, 86, 92, 5, 34, 0, 0, 87, 88, 5, 34, 0, 0, 88, 91,
		5, 34, 0, 0, 89, 91, 8, 0, 0, 0, 90, 87, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0,
		91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1,
		0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 34, 0, 0, 96, 6, 1, 0, 0, 0, 97,
		98, 5, 105, 0, 0, 98, 99, 5, 109, 0, 0, 99, 100, 5, 112, 0, 0, 100, 101,
		5, 111, 0, 0, 101, 102, 5, 114, 0, 0, 102, 103, 5, 116, 0, 0, 103, 8, 1,
		0, 0, 0, 104, 105, 5, 117, 0, 0, 105, 106, 5, 115, 0, 0, 106, 107, 5, 101,
		0, 0, 107, 10, 1, 0, 0, 0, 108, 109, 5, 112, 0, 0, 109, 110, 5, 97, 0,
		0, 110, 111, 5, 99, 0, 0, 111, 112, 5, 107, 0, 0, 112, 113, 5, 97, 0, 0,
		113, 114, 5, 103, 0, 0, 114, 115, 5, 101, 0, 0, 115, 12, 1, 0, 0, 0, 116,
		117, 5, 116, 0, 0, 117, 118, 5, 121, 0, 0, 118, 119, 5, 112, 0, 0, 119,
		120, 5, 101, 0, 0, 120, 14, 1, 0, 0, 0, 121, 122, 5, 99, 0, 0, 122, 123,
		5, 111, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126,
		5, 116, 0, 0, 126, 16, 1, 0, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5,
		101, 0, 0, 129, 130, 5, 114, 0, 0, 130, 131, 5, 118, 0, 0, 131, 132, 5,
		105, 0, 0, 132, 133, 5, 99, 0, 0, 133, 134, 5, 101, 0, 0, 134, 18, 1, 0,
		0, 0, 135, 136, 5, 103, 0, 0, 136, 137, 5, 97, 0, 0, 137, 138, 5, 116,
		0, 0, 138, 139, 5, 101, 0, 0, 139, 140, 5, 119, 0, 0, 140, 141, 5, 97,
		0, 0, 141, 142, 5, 121, 0, 0, 142, 20, 1, 0, 0, 0, 143, 144, 5, 98, 0,
		0, 144, 145, 5, 97, 0, 0, 145, 146, 5, 99, 0, 0, 146, 147, 5, 107, 0, 0,
		147, 148, 5, 101, 0, 0, 148, 149, 5, 110, 0, 0, 149, 150, 5, 100, 0, 0,
		150, 22, 1, 0, 0, 0, 151, 152, 5, 105, 0, 0, 152, 153, 5, 110, 0, 0, 153,
		154, 5, 116, 0, 0, 154, 24, 1, 0, 0, 0, 155, 156, 5, 105, 0, 0, 156, 157,
		5, 110, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 54, 0, 0, 159, 160,
		5, 52, 0, 0, 160, 26, 1, 0, 0, 0, 161, 162, 5, 105, 0, 0, 162, 163, 5,
		110, 0, 0, 163, 164, 5, 116, 0, 0, 164, 165, 5, 49, 0, 0, 165, 166, 5,
		54, 0, 0, 166, 28, 1, 0, 0, 0, 167, 168, 5, 98, 0, 0, 168, 169, 5, 121,
		0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 101, 0, 0, 171, 30, 1, 0, 0,
		0, 172, 173, 5, 98, 0, 0, 173, 174, 5, 111, 0, 0, 174, 175, 5, 111, 0,
		0, 175, 176, 5, 108, 0, 0, 176, 32, 1, 0, 0, 0, 177, 178, 5, 115, 0, 0,
		178, 179, 5, 116, 0, 0, 179, 180, 5, 114, 0, 0, 180, 181, 5, 105, 0, 0,
		181, 182, 5, 110, 0, 0, 182, 183, 5, 103, 0, 0, 183, 34, 1, 0, 0, 0, 184,
		185, 5, 116, 0, 0, 185, 186, 5, 105, 0, 0, 186, 187, 5, 109, 0, 0, 187,
		188, 5, 101, 0, 0, 188, 36, 1, 0, 0, 0, 189, 190, 5, 115, 0, 0, 190, 191,
		5, 116, 0, 0, 191, 192, 5, 114, 0, 0, 192, 193, 5, 117, 0, 0, 193, 194,
		5, 99, 0, 0, 194, 195, 5, 116, 0, 0, 195, 38, 1, 0, 0, 0, 196, 197, 5,
		103, 0, 0, 197, 198, 5, 101, 0, 0, 198, 199, 5, 116, 0, 0, 199, 40, 1,
		0, 0, 0, 200, 201, 5, 112, 0, 0, 201, 202, 5, 111, 0, 0, 202, 203, 5, 115,
		0, 0, 203, 204, 5, 116, 0, 0, 204, 42, 1, 0, 0, 0, 205, 206, 5, 112, 0,
		0, 206, 207, 5, 117, 0, 0, 207, 208, 5, 116, 0, 0, 208, 44, 1, 0, 0, 0,
		209, 210, 5, 112, 0, 0, 210, 211, 5, 97, 0, 0, 211, 212, 5, 116, 0, 0,
		212, 213, 5, 99, 0, 0, 213, 214, 5, 104, 0, 0, 214, 46, 1, 0, 0, 0, 215,
		216, 5, 100, 0, 0, 216, 217, 5, 101, 0, 0, 217, 218, 5, 108, 0, 0, 218,
		219, 5, 101, 0, 0, 219, 220, 5, 116, 0, 0, 220, 221, 5, 101, 0, 0, 221,
		48, 1, 0, 0, 0, 222, 223, 5, 36, 0, 0, 223, 224, 5, 100, 0, 0, 224, 225,
		5, 97, 0, 0, 225, 226, 5, 116, 0, 0, 226, 227, 5, 97, 0, 0, 227, 50, 1,
		0, 0, 0, 228, 229, 5, 61, 0, 0, 229, 52, 1, 0, 0, 0, 230, 231, 5, 58, 0,
		0, 231, 54, 1, 0, 0, 0, 232, 233, 5, 34, 0, 0, 233, 56, 1, 0, 0, 0, 234,
		235, 5, 62, 0, 0, 235, 58, 1, 0, 0, 0, 236, 237, 5, 60, 0, 0, 237, 60,
		1, 0, 0, 0, 238, 239, 5, 40, 0, 0, 239, 62, 1, 0, 0, 0, 240, 241, 5, 41,
		0, 0, 241, 64, 1, 0, 0, 0, 242, 243, 5, 123, 0, 0, 243, 66, 1, 0, 0, 0,
		244, 245, 5, 125, 0, 0, 245, 68, 1, 0, 0, 0, 246, 247, 5, 59, 0, 0, 247,
		70, 1, 0, 0, 0, 248, 249, 5, 64, 0, 0, 249, 72, 1, 0, 0, 0, 250, 259, 7,
		1, 0, 0, 251, 258, 7, 2, 0, 0, 252, 254, 5, 46, 0, 0, 253, 252, 1, 0, 0,
		0, 254, 255, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		258, 1, 0, 0, 0, 257, 251, 1, 0, 0, 0, 257, 253, 1, 0, 0, 0, 258, 261,
		1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 74, 1, 0,
		0, 0, 261, 259, 1, 0, 0, 0, 262, 264, 7, 3, 0, 0, 263, 262, 1, 0, 0, 0,
		264, 265, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266,
		275, 1, 0, 0, 0, 267, 269, 5, 46, 0, 0, 268, 270, 7, 3, 0, 0, 269, 268,
		1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0,
		0, 0, 272, 274, 1, 0, 0, 0, 273, 267, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0,
		275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 76, 1, 0, 0, 0, 277, 275,
		1, 0, 0, 0, 278, 280, 7, 4, 0, 0, 279, 278, 1, 0, 0, 0, 280, 281, 1, 0,
		0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0,
		283, 284, 6, 38, 0, 0, 284, 78, 1, 0, 0, 0, 285, 287, 7, 5, 0, 0, 286,
		285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 6, 39, 0, 0, 291, 80, 1, 0,
		0, 0, 11, 0, 90, 92, 255, 257, 259, 265, 271, 275, 281, 288, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// tesseractLexerInit initializes any static state used to implement tesseractLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewtesseractLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TesseractLexerInit() {
	staticData := &TesseractLexerLexerStaticData
	staticData.once.Do(tesseractlexerLexerInit)
}

// NewtesseractLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewtesseractLexer(input antlr.CharStream) *tesseractLexer {
	TesseractLexerInit()
	l := new(tesseractLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TesseractLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "tesseract.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tesseractLexer tokens.
const (
	tesseractLexerT__0          = 1
	tesseractLexerT__1          = 2
	tesseractLexerEscapedString = 3
	tesseractLexerIMPORT        = 4
	tesseractLexerUSE           = 5
	tesseractLexerPACKAGE       = 6
	tesseractLexerTYPE          = 7
	tesseractLexerCONST         = 8
	tesseractLexerSERVICE       = 9
	tesseractLexerGATEWAY       = 10
	tesseractLexerBACKEND       = 11
	tesseractLexerINT           = 12
	tesseractLexerLONG          = 13
	tesseractLexerSHORT         = 14
	tesseractLexerBYTE          = 15
	tesseractLexerBOOL          = 16
	tesseractLexerSTRING        = 17
	tesseractLexerDATETIME      = 18
	tesseractLexerUNKNOWN       = 19
	tesseractLexerGET           = 20
	tesseractLexerPOST          = 21
	tesseractLexerPUT           = 22
	tesseractLexerPATCH         = 23
	tesseractLexerDELETE        = 24
	tesseractLexerDATA          = 25
	tesseractLexerEQ            = 26
	tesseractLexerCOLON         = 27
	tesseractLexerDOUBLE_QUOTE  = 28
	tesseractLexerGT            = 29
	tesseractLexerLT            = 30
	tesseractLexerLP            = 31
	tesseractLexerRP            = 32
	tesseractLexerLCB           = 33
	tesseractLexerRCB           = 34
	tesseractLexerSEMI          = 35
	tesseractLexerAT            = 36
	tesseractLexerIDENT         = 37
	tesseractLexerNUMBER        = 38
	tesseractLexerNEWLINE       = 39
	tesseractLexerWS            = 40
)
