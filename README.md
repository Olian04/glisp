# glisp

Lisp like go scripting

A Looks like lisp, but all Operations on data is done using functions that hook back out into go. The lisp code is only used to construct data flows on arbitrary lists and dicts.

Should result in fast execution with little to no overhead.

## MVP

* Lexer(io.Reader) -> Token[]
* Parser(Token[]) -> Ast, error
* Compiler(Runtime, Ast) -> Program, error

Expected flow:

1. Write lisp like code and load it into go using an io.Reader interface
2. Lex the code into an array of Tokens
3. Parse the array of tokens into an Ast or a parse error
4. Prepare a Runtime with functions and data types
5. Compile the Ast into a Program by checking it against the Runtime through static type analysis
6. Call functions on the resulting Program
