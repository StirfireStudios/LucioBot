grammar Chat;

/*
 * Lexer Rules
 */

fragment A : [aA]; // match either an 'a' or 'A'
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

fragment LOWERCASE  : [a-z] ;
fragment UPPERCASE  : [A-Z] ;

// Commands

PLAY    : P L A Y ;
LIST    : L I S T ;
SAY     : S A Y ; 
STOP    : S T O P;
STATUS  : S T A T U S;

// fragments

WORD            : (LOWERCASE | UPPERCASE | '_')+ ;
WHITESPACE      : (' ' | '\t') ;
NEWLINE         : ('\r'? '\n' | '\r')+ ;
TEXT            : ('['|'(') .*? (']'|')');

/*
 * Parser Rules
 */

chat            : line+ EOF ;
line            : (mention WHITESPACE)* command NEWLINE;
command         : (paramCommand | singleCommand);
paramCommand    : (PLAY) WHITESPACE WORD;
singleCommand   : (SAY | STOP | LIST | STATUS);
mention         : '@' WORD ;