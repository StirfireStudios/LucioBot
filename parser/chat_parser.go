// Code generated from Chat.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Chat

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 51, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 6, 2, 16, 10, 2, 13, 2, 14, 2, 17, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7,
	3, 25, 10, 3, 12, 3, 14, 3, 28, 11, 3, 3, 3, 3, 3, 7, 3, 32, 10, 3, 12,
	3, 14, 3, 35, 11, 3, 3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12,
	2, 3, 3, 2, 6, 9, 2, 48, 2, 15, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 38, 3,
	2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 44, 3, 2, 2, 2, 12, 46, 3, 2, 2, 2, 14,
	16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15, 3, 2, 2,
	2, 17, 18, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 20, 7, 2, 2, 3, 20, 3, 3,
	2, 2, 2, 21, 22, 5, 12, 7, 2, 22, 23, 7, 12, 2, 2, 23, 25, 3, 2, 2, 2,
	24, 21, 3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3,
	2, 2, 2, 27, 29, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 33, 5, 6, 4, 2, 30,
	32, 7, 13, 2, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2,
	2, 2, 33, 34, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 39,
	5, 8, 5, 2, 37, 39, 5, 10, 6, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2,
	39, 7, 3, 2, 2, 2, 40, 41, 7, 5, 2, 2, 41, 42, 7, 12, 2, 2, 42, 43, 7,
	11, 2, 2, 43, 9, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2, 45, 11, 3, 2, 2, 2, 46,
	47, 7, 3, 2, 2, 47, 48, 7, 10, 2, 2, 48, 49, 7, 4, 2, 2, 49, 13, 3, 2,
	2, 2, 6, 17, 26, 33, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'",
}
var symbolicNames = []string{
	"", "", "", "PLAY", "LIST", "SAY", "STOP", "STATUS", "USERIDENTIFIER",
	"WORD", "WHITESPACE", "NEWLINE", "TEXT",
}

var ruleNames = []string{
	"chat", "line", "command", "paramCommand", "singleCommand", "mention",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ChatParser struct {
	*antlr.BaseParser
}

func NewChatParser(input antlr.TokenStream) *ChatParser {
	this := new(ChatParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chat.g4"

	return this
}

// ChatParser tokens.
const (
	ChatParserEOF            = antlr.TokenEOF
	ChatParserT__0           = 1
	ChatParserT__1           = 2
	ChatParserPLAY           = 3
	ChatParserLIST           = 4
	ChatParserSAY            = 5
	ChatParserSTOP           = 6
	ChatParserSTATUS         = 7
	ChatParserUSERIDENTIFIER = 8
	ChatParserWORD           = 9
	ChatParserWHITESPACE     = 10
	ChatParserNEWLINE        = 11
	ChatParserTEXT           = 12
)

// ChatParser rules.
const (
	ChatParserRULE_chat          = 0
	ChatParserRULE_line          = 1
	ChatParserRULE_command       = 2
	ChatParserRULE_paramCommand  = 3
	ChatParserRULE_singleCommand = 4
	ChatParserRULE_mention       = 5
)

// IChatContext is an interface to support dynamic dispatch.
type IChatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChatContext differentiates from other interfaces.
	IsChatContext()
}

type ChatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChatContext() *ChatContext {
	var p = new(ChatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_chat
	return p
}

func (*ChatContext) IsChatContext() {}

func NewChatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChatContext {
	var p = new(ChatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_chat

	return p
}

func (s *ChatContext) GetParser() antlr.Parser { return s.parser }

func (s *ChatContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChatParserEOF, 0)
}

func (s *ChatContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ChatContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ChatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterChat(s)
	}
}

func (s *ChatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitChat(s)
	}
}

func (p *ChatParser) Chat() (localctx IChatContext) {
	localctx = NewChatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChatParserRULE_chat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChatParserT__0)|(1<<ChatParserPLAY)|(1<<ChatParserLIST)|(1<<ChatParserSAY)|(1<<ChatParserSTOP)|(1<<ChatParserSTATUS))) != 0) {
		{
			p.SetState(12)
			p.Line()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(17)
		p.Match(ChatParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *LineContext) AllMention() []IMentionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMentionContext)(nil)).Elem())
	var tst = make([]IMentionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMentionContext)
		}
	}

	return tst
}

func (s *LineContext) Mention(i int) IMentionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMentionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMentionContext)
}

func (s *LineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ChatParserWHITESPACE)
}

func (s *LineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ChatParserWHITESPACE, i)
}

func (s *LineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ChatParserNEWLINE)
}

func (s *LineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ChatParserNEWLINE, i)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *ChatParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChatParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChatParserT__0 {
		{
			p.SetState(19)
			p.Mention()
		}
		{
			p.SetState(20)
			p.Match(ChatParserWHITESPACE)
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
		p.Command()
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChatParserNEWLINE {
		{
			p.SetState(28)
			p.Match(ChatParserNEWLINE)
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) ParamCommand() IParamCommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamCommandContext)
}

func (s *CommandContext) SingleCommand() ISingleCommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleCommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *ChatParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChatParserRULE_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChatParserPLAY:
		{
			p.SetState(34)
			p.ParamCommand()
		}

	case ChatParserLIST, ChatParserSAY, ChatParserSTOP, ChatParserSTATUS:
		{
			p.SetState(35)
			p.SingleCommand()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamCommandContext is an interface to support dynamic dispatch.
type IParamCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamCommandContext differentiates from other interfaces.
	IsParamCommandContext()
}

type ParamCommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamCommandContext() *ParamCommandContext {
	var p = new(ParamCommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_paramCommand
	return p
}

func (*ParamCommandContext) IsParamCommandContext() {}

func NewParamCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamCommandContext {
	var p = new(ParamCommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_paramCommand

	return p
}

func (s *ParamCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamCommandContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ChatParserWHITESPACE, 0)
}

func (s *ParamCommandContext) WORD() antlr.TerminalNode {
	return s.GetToken(ChatParserWORD, 0)
}

func (s *ParamCommandContext) PLAY() antlr.TerminalNode {
	return s.GetToken(ChatParserPLAY, 0)
}

func (s *ParamCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterParamCommand(s)
	}
}

func (s *ParamCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitParamCommand(s)
	}
}

func (p *ChatParser) ParamCommand() (localctx IParamCommandContext) {
	localctx = NewParamCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChatParserRULE_paramCommand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ChatParserPLAY)
	}

	{
		p.SetState(39)
		p.Match(ChatParserWHITESPACE)
	}
	{
		p.SetState(40)
		p.Match(ChatParserWORD)
	}

	return localctx
}

// ISingleCommandContext is an interface to support dynamic dispatch.
type ISingleCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleCommandContext differentiates from other interfaces.
	IsSingleCommandContext()
}

type SingleCommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleCommandContext() *SingleCommandContext {
	var p = new(SingleCommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_singleCommand
	return p
}

func (*SingleCommandContext) IsSingleCommandContext() {}

func NewSingleCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleCommandContext {
	var p = new(SingleCommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_singleCommand

	return p
}

func (s *SingleCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleCommandContext) SAY() antlr.TerminalNode {
	return s.GetToken(ChatParserSAY, 0)
}

func (s *SingleCommandContext) STOP() antlr.TerminalNode {
	return s.GetToken(ChatParserSTOP, 0)
}

func (s *SingleCommandContext) LIST() antlr.TerminalNode {
	return s.GetToken(ChatParserLIST, 0)
}

func (s *SingleCommandContext) STATUS() antlr.TerminalNode {
	return s.GetToken(ChatParserSTATUS, 0)
}

func (s *SingleCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterSingleCommand(s)
	}
}

func (s *SingleCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitSingleCommand(s)
	}
}

func (p *ChatParser) SingleCommand() (localctx ISingleCommandContext) {
	localctx = NewSingleCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChatParserRULE_singleCommand)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChatParserLIST)|(1<<ChatParserSAY)|(1<<ChatParserSTOP)|(1<<ChatParserSTATUS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMentionContext is an interface to support dynamic dispatch.
type IMentionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMentionContext differentiates from other interfaces.
	IsMentionContext()
}

type MentionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMentionContext() *MentionContext {
	var p = new(MentionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_mention
	return p
}

func (*MentionContext) IsMentionContext() {}

func NewMentionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MentionContext {
	var p = new(MentionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_mention

	return p
}

func (s *MentionContext) GetParser() antlr.Parser { return s.parser }

func (s *MentionContext) USERIDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ChatParserUSERIDENTIFIER, 0)
}

func (s *MentionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MentionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MentionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterMention(s)
	}
}

func (s *MentionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitMention(s)
	}
}

func (p *ChatParser) Mention() (localctx IMentionContext) {
	localctx = NewMentionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChatParserRULE_mention)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(ChatParserT__0)
	}
	{
		p.SetState(45)
		p.Match(ChatParserUSERIDENTIFIER)
	}
	{
		p.SetState(46)
		p.Match(ChatParserT__1)
	}

	return localctx
}
