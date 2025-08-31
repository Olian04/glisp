# glisp

Lisp like go scripting

A Looks like lisp, but all Operations on data is done using functions that hook back out into go. The lisp code is only used to construct data flows on arbitrary lists and dicts.

Should result in fast execution with little to no overhead.

## MVP:

* Lexer(io.Reader) -> Token[]
* Parser(Token[]) -> Ast, error
* StaticalAnalyser(Ast) -> Program, error
