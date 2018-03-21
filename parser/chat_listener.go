// Code generated from Chat.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Chat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChatListener is a complete listener for a parse tree produced by ChatParser.
type ChatListener interface {
	antlr.ParseTreeListener

	// EnterChat is called when entering the chat production.
	EnterChat(c *ChatContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterParamCommand is called when entering the paramCommand production.
	EnterParamCommand(c *ParamCommandContext)

	// EnterSingleCommand is called when entering the singleCommand production.
	EnterSingleCommand(c *SingleCommandContext)

	// EnterMention is called when entering the mention production.
	EnterMention(c *MentionContext)

	// ExitChat is called when exiting the chat production.
	ExitChat(c *ChatContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitParamCommand is called when exiting the paramCommand production.
	ExitParamCommand(c *ParamCommandContext)

	// ExitSingleCommand is called when exiting the singleCommand production.
	ExitSingleCommand(c *SingleCommandContext)

	// ExitMention is called when exiting the mention production.
	ExitMention(c *MentionContext)
}
