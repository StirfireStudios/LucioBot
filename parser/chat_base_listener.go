// Code generated from Chat.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Chat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChatListener is a complete listener for a parse tree produced by ChatParser.
type BaseChatListener struct{}

var _ ChatListener = &BaseChatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChat is called when production chat is entered.
func (s *BaseChatListener) EnterChat(ctx *ChatContext) {}

// ExitChat is called when production chat is exited.
func (s *BaseChatListener) ExitChat(ctx *ChatContext) {}

// EnterLine is called when production line is entered.
func (s *BaseChatListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseChatListener) ExitLine(ctx *LineContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseChatListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseChatListener) ExitCommand(ctx *CommandContext) {}

// EnterParamCommand is called when production paramCommand is entered.
func (s *BaseChatListener) EnterParamCommand(ctx *ParamCommandContext) {}

// ExitParamCommand is called when production paramCommand is exited.
func (s *BaseChatListener) ExitParamCommand(ctx *ParamCommandContext) {}

// EnterSingleCommand is called when production singleCommand is entered.
func (s *BaseChatListener) EnterSingleCommand(ctx *SingleCommandContext) {}

// ExitSingleCommand is called when production singleCommand is exited.
func (s *BaseChatListener) ExitSingleCommand(ctx *SingleCommandContext) {}

// EnterMention is called when production mention is entered.
func (s *BaseChatListener) EnterMention(ctx *MentionContext) {}

// ExitMention is called when production mention is exited.
func (s *BaseChatListener) ExitMention(ctx *MentionContext) {}
