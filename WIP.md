======

VariableAssignment Flag support:

- [x] Add Flag struct to VariableAssignment struct
- [x] Parse Flag keywords and push to VariableAssignment struct

CharsetStatement support

- [x] Add CharsetStatement struct
- [x] Add Encoding field
- [x] Add ParseCharsetStatement method to parser.

SymbolTable

- [x] Register parsed variable to the scope symbol table.
  - [x] RuleSet symbol table
  - [x] Global symbol table
- [ ] Add symbol table lookup method to the expression evaluator.
  - Add type switch case for ast.Variable struct

Optimizer

- [x] Constant Value elimination optimizer for VariableAssignment.
- [ ] Call IfStatementOptimizer after the if statement is parsed.

CSS Slash and Divide

- [x] Review Declaration String() interface.
- [ ] Test simple ruleset output.
- [ ] Add test utility function that accept: {input scss} and {output css}.
- [ ] Add expr stringer test case for `font: 12px/20px`.
- [ ] Add expr stringer test case for expressions like 12px/20px + 13px.

Nested properties

- [ ] Allow declaration block after the colon of property name.
- [ ] Allow declaration block after the property value.
- [ ] `lexPropertyValue` should check if there is another '{' token, then we should go to `lexStatement` state.

`@my` statement

- [ ] Declare variable in the specific scope.


