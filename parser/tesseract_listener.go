// Code generated from ./grammar/tesseract.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // tesseract

import "github.com/antlr4-go/antlr/v4"

// tesseractListener is a complete listener for a parse tree produced by tesseractParser.
type tesseractListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterUseStatement is called when entering the useStatement production.
	EnterUseStatement(c *UseStatementContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterTypeStatement is called when entering the typeStatement production.
	EnterTypeStatement(c *TypeStatementContext)

	// EnterConstStatement is called when entering the constStatement production.
	EnterConstStatement(c *ConstStatementContext)

	// EnterServiceStatement is called when entering the serviceStatement production.
	EnterServiceStatement(c *ServiceStatementContext)

	// EnterGatewayStatement is called when entering the gatewayStatement production.
	EnterGatewayStatement(c *GatewayStatementContext)

	// EnterBackendStatement is called when entering the backendStatement production.
	EnterBackendStatement(c *BackendStatementContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterResponse is called when entering the response production.
	EnterResponse(c *ResponseContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterApi is called when entering the api production.
	EnterApi(c *ApiContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitUseStatement is called when exiting the useStatement production.
	ExitUseStatement(c *UseStatementContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitTypeStatement is called when exiting the typeStatement production.
	ExitTypeStatement(c *TypeStatementContext)

	// ExitConstStatement is called when exiting the constStatement production.
	ExitConstStatement(c *ConstStatementContext)

	// ExitServiceStatement is called when exiting the serviceStatement production.
	ExitServiceStatement(c *ServiceStatementContext)

	// ExitGatewayStatement is called when exiting the gatewayStatement production.
	ExitGatewayStatement(c *GatewayStatementContext)

	// ExitBackendStatement is called when exiting the backendStatement production.
	ExitBackendStatement(c *BackendStatementContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitResponse is called when exiting the response production.
	ExitResponse(c *ResponseContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitApi is called when exiting the api production.
	ExitApi(c *ApiContext)
}
