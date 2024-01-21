// Code generated from ./grammar/tesseract.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // tesseract

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type tesseractParser struct {
	*antlr.BaseParser
}

var TesseractParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tesseractParserInit() {
	staticData := &TesseractParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'[]'", "", "'import'", "'use'", "'package'", "'type'", "'const'",
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
		"program", "importStatement", "useStatement", "packageStatement", "typeStatement",
		"constStatement", "serviceStatement", "gatewayStatement", "backendStatement",
		"type", "list", "fieldType", "field", "assignment", "arg", "call", "attribute",
		"request", "response", "rpc", "api",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 272, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 51, 8, 0, 10, 0, 12,
		0, 54, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 5, 4, 69, 8, 4, 10, 4, 12, 4, 72, 9, 4, 1, 4, 1, 4, 1,
		4, 5, 4, 77, 8, 4, 10, 4, 12, 4, 80, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 87, 8, 4, 10, 4, 12, 4, 90, 9, 4, 1, 4, 1, 4, 5, 4, 94, 8, 4, 10,
		4, 12, 4, 97, 9, 4, 1, 4, 1, 4, 5, 4, 101, 8, 4, 10, 4, 12, 4, 104, 9,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 111, 8, 5, 10, 5, 12, 5, 114, 9,
		5, 1, 5, 1, 5, 1, 6, 5, 6, 119, 8, 6, 10, 6, 12, 6, 122, 9, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 127, 8, 6, 10, 6, 12, 6, 130, 9, 6, 1, 6, 1, 6, 1, 7, 5,
		7, 135, 8, 7, 10, 7, 12, 7, 138, 9, 7, 1, 7, 1, 7, 1, 7, 5, 7, 143, 8,
		7, 10, 7, 12, 7, 146, 9, 7, 1, 7, 1, 7, 1, 8, 5, 8, 151, 8, 8, 10, 8, 12,
		8, 154, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 172, 8, 11, 1, 12, 5,
		12, 175, 8, 12, 10, 12, 12, 12, 178, 9, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 190, 8, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 197, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15,
		202, 8, 15, 10, 15, 12, 15, 205, 9, 15, 1, 15, 1, 15, 5, 15, 209, 8, 15,
		10, 15, 12, 15, 212, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 5,
		17, 220, 8, 17, 10, 17, 12, 17, 223, 9, 17, 1, 17, 1, 17, 1, 18, 5, 18,
		228, 8, 18, 10, 18, 12, 18, 231, 9, 18, 1, 18, 1, 18, 1, 19, 5, 19, 236,
		8, 19, 10, 19, 12, 19, 239, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 5, 20, 249, 8, 20, 10, 20, 12, 20, 252, 9, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 258, 8, 20, 10, 20, 12, 20, 261, 9, 20, 1,
		20, 1, 20, 5, 20, 265, 8, 20, 10, 20, 12, 20, 268, 9, 20, 1, 20, 1, 20,
		1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 0, 3, 2, 0, 3, 3, 38, 38, 2, 0, 12, 19, 37, 37,
		1, 0, 20, 24, 286, 0, 42, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 59, 1, 0, 0,
		0, 6, 63, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 105, 1, 0, 0, 0, 12, 120,
		1, 0, 0, 0, 14, 136, 1, 0, 0, 0, 16, 152, 1, 0, 0, 0, 18, 164, 1, 0, 0,
		0, 20, 166, 1, 0, 0, 0, 22, 171, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 182,
		1, 0, 0, 0, 28, 196, 1, 0, 0, 0, 30, 198, 1, 0, 0, 0, 32, 215, 1, 0, 0,
		0, 34, 221, 1, 0, 0, 0, 36, 229, 1, 0, 0, 0, 38, 237, 1, 0, 0, 0, 40, 250,
		1, 0, 0, 0, 42, 52, 3, 6, 3, 0, 43, 51, 3, 2, 1, 0, 44, 51, 3, 4, 2, 0,
		45, 51, 3, 8, 4, 0, 46, 51, 3, 10, 5, 0, 47, 51, 3, 12, 6, 0, 48, 51, 3,
		16, 8, 0, 49, 51, 3, 14, 7, 0, 50, 43, 1, 0, 0, 0, 50, 44, 1, 0, 0, 0,
		50, 45, 1, 0, 0, 0, 50, 46, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 50, 48, 1,
		0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 1, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56, 5, 4, 0,
		0, 56, 57, 5, 3, 0, 0, 57, 58, 5, 35, 0, 0, 58, 3, 1, 0, 0, 0, 59, 60,
		5, 5, 0, 0, 60, 61, 5, 37, 0, 0, 61, 62, 5, 35, 0, 0, 62, 5, 1, 0, 0, 0,
		63, 64, 5, 6, 0, 0, 64, 65, 5, 37, 0, 0, 65, 66, 5, 35, 0, 0, 66, 7, 1,
		0, 0, 0, 67, 69, 3, 32, 16, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1,
		0, 0, 0, 73, 74, 5, 7, 0, 0, 74, 95, 5, 31, 0, 0, 75, 77, 3, 32, 16, 0,
		76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1,
		0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 37, 0, 0, 82,
		83, 5, 33, 0, 0, 83, 88, 3, 24, 12, 0, 84, 85, 5, 1, 0, 0, 85, 87, 3, 24,
		12, 0, 86, 84, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 34,
		0, 0, 92, 94, 1, 0, 0, 0, 93, 78, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93,
		1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0,
		98, 102, 5, 32, 0, 0, 99, 101, 5, 35, 0, 0, 100, 99, 1, 0, 0, 0, 101, 104,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 9, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 8, 0, 0, 106, 112, 5, 31, 0, 0, 107,
		108, 5, 37, 0, 0, 108, 109, 5, 26, 0, 0, 109, 111, 7, 0, 0, 0, 110, 107,
		1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 32, 0, 0,
		116, 11, 1, 0, 0, 0, 117, 119, 3, 32, 16, 0, 118, 117, 1, 0, 0, 0, 119,
		122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 5, 9, 0, 0, 124, 128, 5, 31,
		0, 0, 125, 127, 3, 38, 19, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0,
		0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130,
		128, 1, 0, 0, 0, 131, 132, 5, 32, 0, 0, 132, 13, 1, 0, 0, 0, 133, 135,
		3, 32, 16, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1,
		0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0,
		0, 139, 140, 5, 10, 0, 0, 140, 144, 5, 31, 0, 0, 141, 143, 3, 40, 20, 0,
		142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 148,
		5, 32, 0, 0, 148, 15, 1, 0, 0, 0, 149, 151, 3, 32, 16, 0, 150, 149, 1,
		0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0,
		0, 153, 155, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5, 11, 0, 0, 156,
		157, 5, 31, 0, 0, 157, 158, 5, 37, 0, 0, 158, 159, 5, 26, 0, 0, 159, 160,
		5, 29, 0, 0, 160, 161, 3, 30, 15, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5,
		32, 0, 0, 163, 17, 1, 0, 0, 0, 164, 165, 7, 1, 0, 0, 165, 19, 1, 0, 0,
		0, 166, 167, 5, 2, 0, 0, 167, 168, 3, 18, 9, 0, 168, 21, 1, 0, 0, 0, 169,
		172, 3, 18, 9, 0, 170, 172, 3, 20, 10, 0, 171, 169, 1, 0, 0, 0, 171, 170,
		1, 0, 0, 0, 172, 23, 1, 0, 0, 0, 173, 175, 3, 32, 16, 0, 174, 173, 1, 0,
		0, 0, 175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0,
		177, 179, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 180, 5, 37, 0, 0, 180,
		181, 3, 22, 11, 0, 181, 25, 1, 0, 0, 0, 182, 183, 5, 37, 0, 0, 183, 189,
		5, 26, 0, 0, 184, 190, 5, 3, 0, 0, 185, 190, 5, 38, 0, 0, 186, 190, 1,
		0, 0, 0, 187, 190, 5, 37, 0, 0, 188, 190, 3, 30, 15, 0, 189, 184, 1, 0,
		0, 0, 189, 185, 1, 0, 0, 0, 189, 186, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0,
		189, 188, 1, 0, 0, 0, 190, 27, 1, 0, 0, 0, 191, 197, 5, 38, 0, 0, 192,
		197, 5, 3, 0, 0, 193, 197, 5, 37, 0, 0, 194, 197, 5, 25, 0, 0, 195, 197,
		3, 30, 15, 0, 196, 191, 1, 0, 0, 0, 196, 192, 1, 0, 0, 0, 196, 193, 1,
		0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 195, 1, 0, 0, 0, 197, 29, 1, 0, 0,
		0, 198, 199, 5, 37, 0, 0, 199, 203, 5, 31, 0, 0, 200, 202, 3, 28, 14, 0,
		201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 210, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207,
		5, 1, 0, 0, 207, 209, 3, 28, 14, 0, 208, 206, 1, 0, 0, 0, 209, 212, 1,
		0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0,
		0, 212, 210, 1, 0, 0, 0, 213, 214, 5, 32, 0, 0, 214, 31, 1, 0, 0, 0, 215,
		216, 5, 36, 0, 0, 216, 217, 3, 30, 15, 0, 217, 33, 1, 0, 0, 0, 218, 220,
		3, 32, 16, 0, 219, 218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1,
		0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0,
		0, 224, 225, 5, 37, 0, 0, 225, 35, 1, 0, 0, 0, 226, 228, 3, 32, 16, 0,
		227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229,
		230, 1, 0, 0, 0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233,
		5, 37, 0, 0, 233, 37, 1, 0, 0, 0, 234, 236, 3, 32, 16, 0, 235, 234, 1,
		0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0,
		0, 238, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 241, 5, 37, 0, 0, 241,
		242, 5, 31, 0, 0, 242, 243, 3, 34, 17, 0, 243, 244, 5, 32, 0, 0, 244, 245,
		3, 36, 18, 0, 245, 246, 5, 35, 0, 0, 246, 39, 1, 0, 0, 0, 247, 249, 3,
		32, 16, 0, 248, 247, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0,
		0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0,
		253, 254, 7, 2, 0, 0, 254, 255, 5, 37, 0, 0, 255, 259, 5, 31, 0, 0, 256,
		258, 3, 34, 17, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0,
		0, 0, 262, 266, 5, 32, 0, 0, 263, 265, 3, 36, 18, 0, 264, 263, 1, 0, 0,
		0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270, 5, 35, 0, 0, 270, 41,
		1, 0, 0, 0, 25, 50, 52, 70, 78, 88, 95, 102, 112, 120, 128, 136, 144, 152,
		171, 176, 189, 196, 203, 210, 221, 229, 237, 250, 259, 266,
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

// tesseractParserInit initializes any static state used to implement tesseractParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtesseractParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TesseractParserInit() {
	staticData := &TesseractParserStaticData
	staticData.once.Do(tesseractParserInit)
}

// NewtesseractParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtesseractParser(input antlr.TokenStream) *tesseractParser {
	TesseractParserInit()
	this := new(tesseractParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TesseractParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "tesseract.g4"

	return this
}

// tesseractParser tokens.
const (
	tesseractParserEOF           = antlr.TokenEOF
	tesseractParserT__0          = 1
	tesseractParserT__1          = 2
	tesseractParserEscapedString = 3
	tesseractParserIMPORT        = 4
	tesseractParserUSE           = 5
	tesseractParserPACKAGE       = 6
	tesseractParserTYPE          = 7
	tesseractParserCONST         = 8
	tesseractParserSERVICE       = 9
	tesseractParserGATEWAY       = 10
	tesseractParserBACKEND       = 11
	tesseractParserINT           = 12
	tesseractParserLONG          = 13
	tesseractParserSHORT         = 14
	tesseractParserBYTE          = 15
	tesseractParserBOOL          = 16
	tesseractParserSTRING        = 17
	tesseractParserDATETIME      = 18
	tesseractParserUNKNOWN       = 19
	tesseractParserGET           = 20
	tesseractParserPOST          = 21
	tesseractParserPUT           = 22
	tesseractParserPATCH         = 23
	tesseractParserDELETE        = 24
	tesseractParserDATA          = 25
	tesseractParserEQ            = 26
	tesseractParserCOLON         = 27
	tesseractParserDOUBLE_QUOTE  = 28
	tesseractParserGT            = 29
	tesseractParserLT            = 30
	tesseractParserLP            = 31
	tesseractParserRP            = 32
	tesseractParserLCB           = 33
	tesseractParserRCB           = 34
	tesseractParserSEMI          = 35
	tesseractParserAT            = 36
	tesseractParserIDENT         = 37
	tesseractParserNUMBER        = 38
	tesseractParserNEWLINE       = 39
	tesseractParserWS            = 40
)

// tesseractParser rules.
const (
	tesseractParserRULE_program          = 0
	tesseractParserRULE_importStatement  = 1
	tesseractParserRULE_useStatement     = 2
	tesseractParserRULE_packageStatement = 3
	tesseractParserRULE_typeStatement    = 4
	tesseractParserRULE_constStatement   = 5
	tesseractParserRULE_serviceStatement = 6
	tesseractParserRULE_gatewayStatement = 7
	tesseractParserRULE_backendStatement = 8
	tesseractParserRULE_type             = 9
	tesseractParserRULE_list             = 10
	tesseractParserRULE_fieldType        = 11
	tesseractParserRULE_field            = 12
	tesseractParserRULE_assignment       = 13
	tesseractParserRULE_arg              = 14
	tesseractParserRULE_call             = 15
	tesseractParserRULE_attribute        = 16
	tesseractParserRULE_request          = 17
	tesseractParserRULE_response         = 18
	tesseractParserRULE_rpc              = 19
	tesseractParserRULE_api              = 20
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackageStatement() IPackageStatementContext
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllUseStatement() []IUseStatementContext
	UseStatement(i int) IUseStatementContext
	AllTypeStatement() []ITypeStatementContext
	TypeStatement(i int) ITypeStatementContext
	AllConstStatement() []IConstStatementContext
	ConstStatement(i int) IConstStatementContext
	AllServiceStatement() []IServiceStatementContext
	ServiceStatement(i int) IServiceStatementContext
	AllBackendStatement() []IBackendStatementContext
	BackendStatement(i int) IBackendStatementContext
	AllGatewayStatement() []IGatewayStatementContext
	GatewayStatement(i int) IGatewayStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PackageStatement() IPackageStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProgramContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProgramContext) AllUseStatement() []IUseStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUseStatementContext); ok {
			len++
		}
	}

	tst := make([]IUseStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUseStatementContext); ok {
			tst[i] = t.(IUseStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) UseStatement(i int) IUseStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStatementContext)
}

func (s *ProgramContext) AllTypeStatement() []ITypeStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeStatementContext); ok {
			len++
		}
	}

	tst := make([]ITypeStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeStatementContext); ok {
			tst[i] = t.(ITypeStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TypeStatement(i int) ITypeStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeStatementContext)
}

func (s *ProgramContext) AllConstStatement() []IConstStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstStatementContext); ok {
			len++
		}
	}

	tst := make([]IConstStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstStatementContext); ok {
			tst[i] = t.(IConstStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ConstStatement(i int) IConstStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstStatementContext)
}

func (s *ProgramContext) AllServiceStatement() []IServiceStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceStatementContext); ok {
			len++
		}
	}

	tst := make([]IServiceStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceStatementContext); ok {
			tst[i] = t.(IServiceStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ServiceStatement(i int) IServiceStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceStatementContext)
}

func (s *ProgramContext) AllBackendStatement() []IBackendStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBackendStatementContext); ok {
			len++
		}
	}

	tst := make([]IBackendStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBackendStatementContext); ok {
			tst[i] = t.(IBackendStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) BackendStatement(i int) IBackendStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBackendStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBackendStatementContext)
}

func (s *ProgramContext) AllGatewayStatement() []IGatewayStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGatewayStatementContext); ok {
			len++
		}
	}

	tst := make([]IGatewayStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGatewayStatementContext); ok {
			tst[i] = t.(IGatewayStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) GatewayStatement(i int) IGatewayStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGatewayStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGatewayStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *tesseractParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tesseractParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.PackageStatement()
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719480752) != 0 {
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(43)
				p.ImportStatement()
			}

		case 2:
			{
				p.SetState(44)
				p.UseStatement()
			}

		case 3:
			{
				p.SetState(45)
				p.TypeStatement()
			}

		case 4:
			{
				p.SetState(46)
				p.ConstStatement()
			}

		case 5:
			{
				p.SetState(47)
				p.ServiceStatement()
			}

		case 6:
			{
				p.SetState(48)
				p.BackendStatement()
			}

		case 7:
			{
				p.SetState(49)
				p.GatewayStatement()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	EscapedString() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIMPORT, 0)
}

func (s *ImportStatementContext) EscapedString() antlr.TerminalNode {
	return s.GetToken(tesseractParserEscapedString, 0)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *tesseractParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tesseractParserRULE_importStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(tesseractParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(tesseractParserEscapedString)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(tesseractParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUseStatementContext is an interface to support dynamic dispatch.
type IUseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsUseStatementContext differentiates from other interfaces.
	IsUseStatementContext()
}

type UseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseStatementContext() *UseStatementContext {
	var p = new(UseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_useStatement
	return p
}

func InitEmptyUseStatementContext(p *UseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_useStatement
}

func (*UseStatementContext) IsUseStatementContext() {}

func NewUseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStatementContext {
	var p = new(UseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_useStatement

	return p
}

func (s *UseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStatementContext) USE() antlr.TerminalNode {
	return s.GetToken(tesseractParserUSE, 0)
}

func (s *UseStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *UseStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, 0)
}

func (s *UseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterUseStatement(s)
	}
}

func (s *UseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitUseStatement(s)
	}
}

func (p *tesseractParser) UseStatement() (localctx IUseStatementContext) {
	localctx = NewUseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tesseractParserRULE_useStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(tesseractParserUSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(tesseractParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(tesseractParserPACKAGE, 0)
}

func (s *PackageStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *tesseractParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tesseractParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(tesseractParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(tesseractParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeStatementContext is an interface to support dynamic dispatch.
type ITypeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllLCB() []antlr.TerminalNode
	LCB(i int) antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllRCB() []antlr.TerminalNode
	RCB(i int) antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsTypeStatementContext differentiates from other interfaces.
	IsTypeStatementContext()
}

type TypeStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeStatementContext() *TypeStatementContext {
	var p = new(TypeStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_typeStatement
	return p
}

func InitEmptyTypeStatementContext(p *TypeStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_typeStatement
}

func (*TypeStatementContext) IsTypeStatementContext() {}

func NewTypeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStatementContext {
	var p = new(TypeStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_typeStatement

	return p
}

func (s *TypeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStatementContext) TYPE() antlr.TerminalNode {
	return s.GetToken(tesseractParserTYPE, 0)
}

func (s *TypeStatementContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *TypeStatementContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *TypeStatementContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *TypeStatementContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *TypeStatementContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserIDENT)
}

func (s *TypeStatementContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, i)
}

func (s *TypeStatementContext) AllLCB() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserLCB)
}

func (s *TypeStatementContext) LCB(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserLCB, i)
}

func (s *TypeStatementContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeStatementContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeStatementContext) AllRCB() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserRCB)
}

func (s *TypeStatementContext) RCB(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserRCB, i)
}

func (s *TypeStatementContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserSEMI)
}

func (s *TypeStatementContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, i)
}

func (s *TypeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterTypeStatement(s)
	}
}

func (s *TypeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitTypeStatement(s)
	}
}

func (p *tesseractParser) TypeStatement() (localctx ITypeStatementContext) {
	localctx = NewTypeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tesseractParserRULE_typeStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(67)
			p.Attribute()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(tesseractParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT || _la == tesseractParserIDENT {
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == tesseractParserAT {
			{
				p.SetState(75)
				p.Attribute()
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)
			p.Match(tesseractParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(tesseractParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Field()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == tesseractParserT__0 {
			{
				p.SetState(84)
				p.Match(tesseractParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(85)
				p.Field()
			}

			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(91)
			p.Match(tesseractParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserSEMI {
		{
			p.SetState(99)
			p.Match(tesseractParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstStatementContext is an interface to support dynamic dispatch.
type IConstStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllEscapedString() []antlr.TerminalNode
	EscapedString(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsConstStatementContext differentiates from other interfaces.
	IsConstStatementContext()
}

type ConstStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstStatementContext() *ConstStatementContext {
	var p = new(ConstStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_constStatement
	return p
}

func InitEmptyConstStatementContext(p *ConstStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_constStatement
}

func (*ConstStatementContext) IsConstStatementContext() {}

func NewConstStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstStatementContext {
	var p = new(ConstStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_constStatement

	return p
}

func (s *ConstStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstStatementContext) CONST() antlr.TerminalNode {
	return s.GetToken(tesseractParserCONST, 0)
}

func (s *ConstStatementContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *ConstStatementContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *ConstStatementContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserIDENT)
}

func (s *ConstStatementContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, i)
}

func (s *ConstStatementContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserEQ)
}

func (s *ConstStatementContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserEQ, i)
}

func (s *ConstStatementContext) AllEscapedString() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserEscapedString)
}

func (s *ConstStatementContext) EscapedString(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserEscapedString, i)
}

func (s *ConstStatementContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserNUMBER)
}

func (s *ConstStatementContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserNUMBER, i)
}

func (s *ConstStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterConstStatement(s)
	}
}

func (s *ConstStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitConstStatement(s)
	}
}

func (p *tesseractParser) ConstStatement() (localctx IConstStatementContext) {
	localctx = NewConstStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tesseractParserRULE_constStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(tesseractParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserIDENT {
		{
			p.SetState(107)
			p.Match(tesseractParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(tesseractParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tesseractParserEscapedString || _la == tesseractParserNUMBER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceStatementContext is an interface to support dynamic dispatch.
type IServiceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllRpc() []IRpcContext
	Rpc(i int) IRpcContext

	// IsServiceStatementContext differentiates from other interfaces.
	IsServiceStatementContext()
}

type ServiceStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceStatementContext() *ServiceStatementContext {
	var p = new(ServiceStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_serviceStatement
	return p
}

func InitEmptyServiceStatementContext(p *ServiceStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_serviceStatement
}

func (*ServiceStatementContext) IsServiceStatementContext() {}

func NewServiceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceStatementContext {
	var p = new(ServiceStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_serviceStatement

	return p
}

func (s *ServiceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceStatementContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(tesseractParserSERVICE, 0)
}

func (s *ServiceStatementContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *ServiceStatementContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *ServiceStatementContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *ServiceStatementContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ServiceStatementContext) AllRpc() []IRpcContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRpcContext); ok {
			len++
		}
	}

	tst := make([]IRpcContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRpcContext); ok {
			tst[i] = t.(IRpcContext)
			i++
		}
	}

	return tst
}

func (s *ServiceStatementContext) Rpc(i int) IRpcContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterServiceStatement(s)
	}
}

func (s *ServiceStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitServiceStatement(s)
	}
}

func (p *tesseractParser) ServiceStatement() (localctx IServiceStatementContext) {
	localctx = NewServiceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tesseractParserRULE_serviceStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(117)
			p.Attribute()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
		p.Match(tesseractParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT || _la == tesseractParserIDENT {
		{
			p.SetState(125)
			p.Rpc()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGatewayStatementContext is an interface to support dynamic dispatch.
type IGatewayStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GATEWAY() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllApi() []IApiContext
	Api(i int) IApiContext

	// IsGatewayStatementContext differentiates from other interfaces.
	IsGatewayStatementContext()
}

type GatewayStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGatewayStatementContext() *GatewayStatementContext {
	var p = new(GatewayStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_gatewayStatement
	return p
}

func InitEmptyGatewayStatementContext(p *GatewayStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_gatewayStatement
}

func (*GatewayStatementContext) IsGatewayStatementContext() {}

func NewGatewayStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GatewayStatementContext {
	var p = new(GatewayStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_gatewayStatement

	return p
}

func (s *GatewayStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GatewayStatementContext) GATEWAY() antlr.TerminalNode {
	return s.GetToken(tesseractParserGATEWAY, 0)
}

func (s *GatewayStatementContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *GatewayStatementContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *GatewayStatementContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *GatewayStatementContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *GatewayStatementContext) AllApi() []IApiContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApiContext); ok {
			len++
		}
	}

	tst := make([]IApiContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApiContext); ok {
			tst[i] = t.(IApiContext)
			i++
		}
	}

	return tst
}

func (s *GatewayStatementContext) Api(i int) IApiContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApiContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApiContext)
}

func (s *GatewayStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GatewayStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GatewayStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterGatewayStatement(s)
	}
}

func (s *GatewayStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitGatewayStatement(s)
	}
}

func (p *tesseractParser) GatewayStatement() (localctx IGatewayStatementContext) {
	localctx = NewGatewayStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tesseractParserRULE_gatewayStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(133)
			p.Attribute()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(139)
		p.Match(tesseractParserGATEWAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68751982592) != 0 {
		{
			p.SetState(141)
			p.Api()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBackendStatementContext is an interface to support dynamic dispatch.
type IBackendStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BACKEND() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	Call() ICallContext
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsBackendStatementContext differentiates from other interfaces.
	IsBackendStatementContext()
}

type BackendStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBackendStatementContext() *BackendStatementContext {
	var p = new(BackendStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_backendStatement
	return p
}

func InitEmptyBackendStatementContext(p *BackendStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_backendStatement
}

func (*BackendStatementContext) IsBackendStatementContext() {}

func NewBackendStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BackendStatementContext {
	var p = new(BackendStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_backendStatement

	return p
}

func (s *BackendStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BackendStatementContext) BACKEND() antlr.TerminalNode {
	return s.GetToken(tesseractParserBACKEND, 0)
}

func (s *BackendStatementContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *BackendStatementContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *BackendStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *BackendStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(tesseractParserEQ, 0)
}

func (s *BackendStatementContext) GT() antlr.TerminalNode {
	return s.GetToken(tesseractParserGT, 0)
}

func (s *BackendStatementContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *BackendStatementContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *BackendStatementContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *BackendStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackendStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BackendStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterBackendStatement(s)
	}
}

func (s *BackendStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitBackendStatement(s)
	}
}

func (p *tesseractParser) BackendStatement() (localctx IBackendStatementContext) {
	localctx = NewBackendStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tesseractParserRULE_backendStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(149)
			p.Attribute()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(tesseractParserBACKEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(157)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(tesseractParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(tesseractParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Call()
	}

	{
		p.SetState(162)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INT() antlr.TerminalNode
	LONG() antlr.TerminalNode
	SHORT() antlr.TerminalNode
	BYTE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	DATETIME() antlr.TerminalNode
	UNKNOWN() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(tesseractParserINT, 0)
}

func (s *TypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(tesseractParserLONG, 0)
}

func (s *TypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(tesseractParserSHORT, 0)
}

func (s *TypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(tesseractParserBYTE, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(tesseractParserBOOL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(tesseractParserSTRING, 0)
}

func (s *TypeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(tesseractParserDATETIME, 0)
}

func (s *TypeContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(tesseractParserUNKNOWN, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *tesseractParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tesseractParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137439997952) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *tesseractParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tesseractParserRULE_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(tesseractParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	List() IListContext

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_fieldType
	return p
}

func InitEmptyFieldTypeContext(p *FieldTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_fieldType
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldTypeContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *tesseractParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tesseractParserRULE_fieldType)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tesseractParserINT, tesseractParserLONG, tesseractParserSHORT, tesseractParserBYTE, tesseractParserBOOL, tesseractParserSTRING, tesseractParserDATETIME, tesseractParserUNKNOWN, tesseractParserIDENT:
		{
			p.SetState(169)
			p.Type_()
		}

	case tesseractParserT__1:
		{
			p.SetState(170)
			p.List()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	FieldType() IFieldTypeContext
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *tesseractParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tesseractParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(173)
			p.Attribute()
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.FieldType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	EQ() antlr.TerminalNode
	EscapedString() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Call() ICallContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(tesseractParserIDENT)
}

func (s *AssignmentContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, i)
}

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(tesseractParserEQ, 0)
}

func (s *AssignmentContext) EscapedString() antlr.TerminalNode {
	return s.GetToken(tesseractParserEscapedString, 0)
}

func (s *AssignmentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(tesseractParserNUMBER, 0)
}

func (s *AssignmentContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *tesseractParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tesseractParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(tesseractParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(184)
			p.Match(tesseractParserEscapedString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(185)
			p.Match(tesseractParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:

	case 4:
		{
			p.SetState(187)
			p.Match(tesseractParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(188)
			p.Call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	EscapedString() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	DATA() antlr.TerminalNode
	Call() ICallContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(tesseractParserNUMBER, 0)
}

func (s *ArgContext) EscapedString() antlr.TerminalNode {
	return s.GetToken(tesseractParserEscapedString, 0)
}

func (s *ArgContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *ArgContext) DATA() antlr.TerminalNode {
	return s.GetToken(tesseractParserDATA, 0)
}

func (s *ArgContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *tesseractParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tesseractParserRULE_arg)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.Match(tesseractParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(tesseractParserEscapedString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Match(tesseractParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(tesseractParserDATA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.Call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	AllArg() []IArgContext
	Arg(i int) IArgContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *CallContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *CallContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *CallContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *CallContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *tesseractParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tesseractParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412350414856) != 0 {
		{
			p.SetState(200)
			p.Arg()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserT__0 {
		{
			p.SetState(206)
			p.Match(tesseractParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Arg()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	Call() ICallContext

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) AT() antlr.TerminalNode {
	return s.GetToken(tesseractParserAT, 0)
}

func (s *AttributeContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *tesseractParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tesseractParserRULE_attribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(tesseractParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Call()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_request
	return p
}

func InitEmptyRequestContext(p *RequestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_request
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *RequestContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *RequestContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *tesseractParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tesseractParserRULE_request)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(218)
			p.Attribute()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(224)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResponseContext is an interface to support dynamic dispatch.
type IResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsResponseContext differentiates from other interfaces.
	IsResponseContext()
}

type ResponseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponseContext() *ResponseContext {
	var p = new(ResponseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_response
	return p
}

func InitEmptyResponseContext(p *ResponseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_response
}

func (*ResponseContext) IsResponseContext() {}

func NewResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponseContext {
	var p = new(ResponseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_response

	return p
}

func (s *ResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponseContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *ResponseContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *ResponseContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterResponse(s)
	}
}

func (s *ResponseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitResponse(s)
	}
}

func (p *tesseractParser) Response() (localctx IResponseContext) {
	localctx = NewResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tesseractParserRULE_response)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(226)
			p.Attribute()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(232)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	LP() antlr.TerminalNode
	Request() IRequestContext
	RP() antlr.TerminalNode
	Response() IResponseContext
	SEMI() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_rpc
	return p
}

func InitEmptyRpcContext(p *RpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_rpc
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *RpcContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *RpcContext) Request() IRequestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RpcContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *RpcContext) Response() IResponseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResponseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResponseContext)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, 0)
}

func (s *RpcContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *tesseractParser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tesseractParserRULE_rpc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(234)
			p.Attribute()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(240)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Request()
	}
	{
		p.SetState(243)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Response()
	}
	{
		p.SetState(245)
		p.Match(tesseractParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	GET() antlr.TerminalNode
	POST() antlr.TerminalNode
	PUT() antlr.TerminalNode
	PATCH() antlr.TerminalNode
	DELETE() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllRequest() []IRequestContext
	Request(i int) IRequestContext
	AllResponse() []IResponseContext
	Response(i int) IResponseContext

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_api
	return p
}

func InitEmptyApiContext(p *ApiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tesseractParserRULE_api
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tesseractParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) IDENT() antlr.TerminalNode {
	return s.GetToken(tesseractParserIDENT, 0)
}

func (s *ApiContext) LP() antlr.TerminalNode {
	return s.GetToken(tesseractParserLP, 0)
}

func (s *ApiContext) RP() antlr.TerminalNode {
	return s.GetToken(tesseractParserRP, 0)
}

func (s *ApiContext) SEMI() antlr.TerminalNode {
	return s.GetToken(tesseractParserSEMI, 0)
}

func (s *ApiContext) GET() antlr.TerminalNode {
	return s.GetToken(tesseractParserGET, 0)
}

func (s *ApiContext) POST() antlr.TerminalNode {
	return s.GetToken(tesseractParserPOST, 0)
}

func (s *ApiContext) PUT() antlr.TerminalNode {
	return s.GetToken(tesseractParserPUT, 0)
}

func (s *ApiContext) PATCH() antlr.TerminalNode {
	return s.GetToken(tesseractParserPATCH, 0)
}

func (s *ApiContext) DELETE() antlr.TerminalNode {
	return s.GetToken(tesseractParserDELETE, 0)
}

func (s *ApiContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ApiContext) AllRequest() []IRequestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRequestContext); ok {
			len++
		}
	}

	tst := make([]IRequestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRequestContext); ok {
			tst[i] = t.(IRequestContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Request(i int) IRequestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *ApiContext) AllResponse() []IResponseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResponseContext); ok {
			len++
		}
	}

	tst := make([]IResponseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResponseContext); ok {
			tst[i] = t.(IResponseContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Response(i int) IResponseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResponseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResponseContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.EnterApi(s)
	}
}

func (s *ApiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tesseractListener); ok {
		listenerT.ExitApi(s)
	}
}

func (p *tesseractParser) Api() (localctx IApiContext) {
	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tesseractParserRULE_api)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT {
		{
			p.SetState(247)
			p.Attribute()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32505856) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(254)
		p.Match(tesseractParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(tesseractParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT || _la == tesseractParserIDENT {
		{
			p.SetState(256)
			p.Request()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
		p.Match(tesseractParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tesseractParserAT || _la == tesseractParserIDENT {
		{
			p.SetState(263)
			p.Response()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Match(tesseractParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
