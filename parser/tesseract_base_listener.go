// Code generated from ./grammar/tesseract.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // tesseract

import "github.com/antlr4-go/antlr/v4"

// BasetesseractListener is a complete listener for a parse tree produced by tesseractParser.
type BasetesseractListener struct{}

var _ tesseractListener = &BasetesseractListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetesseractListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetesseractListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetesseractListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetesseractListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetesseractListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetesseractListener) ExitProgram(ctx *ProgramContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BasetesseractListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BasetesseractListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterUseStatement is called when production useStatement is entered.
func (s *BasetesseractListener) EnterUseStatement(ctx *UseStatementContext) {}

// ExitUseStatement is called when production useStatement is exited.
func (s *BasetesseractListener) ExitUseStatement(ctx *UseStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BasetesseractListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BasetesseractListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterTypeStatement is called when production typeStatement is entered.
func (s *BasetesseractListener) EnterTypeStatement(ctx *TypeStatementContext) {}

// ExitTypeStatement is called when production typeStatement is exited.
func (s *BasetesseractListener) ExitTypeStatement(ctx *TypeStatementContext) {}

// EnterConstStatement is called when production constStatement is entered.
func (s *BasetesseractListener) EnterConstStatement(ctx *ConstStatementContext) {}

// ExitConstStatement is called when production constStatement is exited.
func (s *BasetesseractListener) ExitConstStatement(ctx *ConstStatementContext) {}

// EnterServiceStatement is called when production serviceStatement is entered.
func (s *BasetesseractListener) EnterServiceStatement(ctx *ServiceStatementContext) {}

// ExitServiceStatement is called when production serviceStatement is exited.
func (s *BasetesseractListener) ExitServiceStatement(ctx *ServiceStatementContext) {}

// EnterGatewayStatement is called when production gatewayStatement is entered.
func (s *BasetesseractListener) EnterGatewayStatement(ctx *GatewayStatementContext) {}

// ExitGatewayStatement is called when production gatewayStatement is exited.
func (s *BasetesseractListener) ExitGatewayStatement(ctx *GatewayStatementContext) {}

// EnterBackendStatement is called when production backendStatement is entered.
func (s *BasetesseractListener) EnterBackendStatement(ctx *BackendStatementContext) {}

// ExitBackendStatement is called when production backendStatement is exited.
func (s *BasetesseractListener) ExitBackendStatement(ctx *BackendStatementContext) {}

// EnterType is called when production type is entered.
func (s *BasetesseractListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasetesseractListener) ExitType(ctx *TypeContext) {}

// EnterList is called when production list is entered.
func (s *BasetesseractListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasetesseractListener) ExitList(ctx *ListContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BasetesseractListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BasetesseractListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterField is called when production field is entered.
func (s *BasetesseractListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasetesseractListener) ExitField(ctx *FieldContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasetesseractListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasetesseractListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterArg is called when production arg is entered.
func (s *BasetesseractListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BasetesseractListener) ExitArg(ctx *ArgContext) {}

// EnterCall is called when production call is entered.
func (s *BasetesseractListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BasetesseractListener) ExitCall(ctx *CallContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasetesseractListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasetesseractListener) ExitAttribute(ctx *AttributeContext) {}

// EnterRequest is called when production request is entered.
func (s *BasetesseractListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BasetesseractListener) ExitRequest(ctx *RequestContext) {}

// EnterResponse is called when production response is entered.
func (s *BasetesseractListener) EnterResponse(ctx *ResponseContext) {}

// ExitResponse is called when production response is exited.
func (s *BasetesseractListener) ExitResponse(ctx *ResponseContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BasetesseractListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BasetesseractListener) ExitRpc(ctx *RpcContext) {}

// EnterApi is called when production api is entered.
func (s *BasetesseractListener) EnterApi(ctx *ApiContext) {}

// ExitApi is called when production api is exited.
func (s *BasetesseractListener) ExitApi(ctx *ApiContext) {}
