package main

import (
	"fmt"
	"strings"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nlopes/slack"
)

type command int

const (
	Play command = iota
	List
	Say
	Stop
	Status
)

var (
	myUserID string
)

type messageListener struct {
	*parser.BaseChatListener

	privateChannel bool
	userIdentifier string

	filename string
	command  command
	err      bool
}

func replyError(msg *slack.MessageEvent) {
	api := Config().SlackAPI()
	params := slack.PostMessageParameters{}
	api.PostMessage(msg.Channel, "Sorry, I couldn't understand you!", params)
}

func handleConnection(msg *slack.ConnectedEvent) {
	myUserID = msg.Info.User.ID
}

func handleMessage(msg *slack.MessageEvent) {
	if msg.Type != "message" {
		return
	}

	if msg.User == myUserID || len(msg.User) == 0 {
		return
	}

	if msg.Hidden {
		return
	}

	if len(msg.Text) == 0 {
		return
	}

	parsedMsg := messageListener{
		privateChannel: ChannelIsPrivate(msg.Channel),
	}

	chars := antlr.NewInputStream(msg.Text)
	lexer := parser.NewChatLexer(chars)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewChatParser(tokens)
	//	parser.RemoveErrorListeners()
	parser.BuildParseTrees = true
	tree := parser.Chat()

	antlr.ParseTreeWalkerDefault.Walk(&parsedMsg, tree)

	if !parsedMsg.ForMe() {
		fmt.Printf("Message '%s' was not for me\n", msg.Text)
		fmt.Printf("%#v\n", parsedMsg)
		return
	}

	if parsedMsg.err {
		replyError(msg)
	}
}

func (msg *messageListener) ForMe() bool {
	if msg.privateChannel {
		return true
	}

	if msg.userIdentifier == myUserID {
		return true
	}

	return false
}

func getCommand(node antlr.TerminalNode) command {
	switch node.GetSymbol().GetTokenType() {
	case parser.ChatLexerPLAY:
		return Play
	case parser.ChatLexerSAY:
		return Say
	case parser.ChatLexerSTOP:
		return Stop
	case parser.ChatLexerSTATUS:
		return Status
	case parser.ChatLexerLIST:
		return List
	}
	return Status
}

func (msg *messageListener) VisitErrorNode(node antlr.ErrorNode) {
	msg.err = true
}

func (msg *messageListener) EnterParamCommand(ctx *parser.ParamCommandContext) {
	msg.command = getCommand(ctx.GetChild(0).(antlr.TerminalNode))
	msg.filename = ctx.GetChild(2).(antlr.TerminalNode).GetText()
}

func (msg *messageListener) EnterSingleCommand(ctx *parser.SingleCommandContext) {
	msg.command = getCommand(ctx.GetChild(0).(antlr.TerminalNode))
}

func (msg *messageListener) EnterMention(ctx *parser.MentionContext) {
	msg.userIdentifier = strings.TrimLeft(ctx.GetChild(1).(antlr.TerminalNode).GetText(), "@")
}