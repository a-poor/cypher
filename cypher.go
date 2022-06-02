package cypher

// Cypher
//
// EBNF Grammar:
//   Cypher = [SP], Statement, [[SP], ';'], [SP], EOI ;
type Cypher struct {
	Statement *Statement `@@ ";"?`
}

// Statement
//
// EBNF Grammar:
//   Statement = Query ;
type Statement struct {
	Query *Query `@@`
}

// Query
//
// EBNF Grammar:
//   Query = RegularQuery
//         | StandaloneCall
//         ;
type Query struct {
	RegularQuery   *RegularQuery   `  @@`
	StandaloneCall *StandaloneCall `| @@`
}

// RegularQuery
//
// EBNF Grammar:
//   RegularQuery = SingleQuery, { [SP], Union } ;
type RegularQuery struct {
	SingleQuery *SingleQuery `@@`
	Union       *Union       `@@*`
}

// Union
//
// EBNF Grammar:
//   Union = ((U,N,I,O,N), SP, (A,L,L), [SP], SingleQuery)
//         | ((U,N,I,O,N), [SP], SingleQuery)
//         ;
type Union struct {
	All         bool         `(("U"|"u") ("N"|"n") ("I"|"i") ("O"|"o") ("N"|"n")) @(("A"|"a") ("L"|"l") ("L"|"l"))?`
	SingleQuery *SingleQuery `@@`
}

// SingleQuery
//
// EBNF Grammar:
//   SingleQuery = SinglePartQuery
//               | MultiPartQuery
//               ;
type SingleQuery struct {
	SinglePart *SinglePartQuery `@@`
	MultiPart  *MultiPartQuery  `| @@`
}

// SinglePartQuery
//
// EBNF Grammar:
//   SinglePartQuery = ({ ReadingClause, [SP] }, Return)
//                   | ({ ReadingClause, [SP] }, UpdatingClause, { [SP], UpdatingClause }, [[SP], Return])
//                   ;
type SinglePartQuery struct {
	ReadingClauses  []*ReadingClause  `@@*`
	UpdatingClauses []*UpdatingClause `@@*`
	Return          *Return           `@@`
}

// MultiPartQuery
//
// EBNF Grammar:
//   MultiPartQuery = { { ReadingClause, [SP] }, { UpdatingClause, [SP] }, With, [SP] }-, SinglePartQuery ;
type MultiPartQuery struct{}

// UpdatingClause
//
// EBNF Grammar:
//   UpdatingClause = Create
//                  | Merge
//                  | Delete
//                  | Set
//                  | Remove
//                  ;
type UpdatingClause struct {
	Create *Create `@@`
	Merge  *Merge  `| @@`
	Delete *Delete `| @@`
	Set    *Set    `| @@`
	Remove *Remove `| @@`
}

// ReadingClause
//
// EBNF Grammar:
//   ReadingClause = Match
//                 | Unwind
//                 | InQueryCall
//                 ;
type ReadingClause struct {
	Match   *Match       `@@`
	Unwind  *Unwind      `| @@`
	InQuery *InQueryCall `| @@`
}

// Match
//
// EBNF Grammar:
//   Match = [(O,P,T,I,O,N,A,L), SP], (M,A,T,C,H), [SP], Pattern, [[SP], Where] ;
type Match struct {
	Optional bool     `@(("O"|"o") ("P"|"p") ("T"|"t") ("I"|"i") ("O"|"o") ("N"|"n") ("A"|"a") ("L"|"l"))?`
	Pattern  *Pattern `(("M"|"m") ("A"|"a") ("T"|"t") ("C"|"c") ("H"|"h")) @@`
	Where    *Where   `@@*`
}

// Unwind
//
// EBNF Grammar:
//   Unwind = (U,N,W,I,N,D), [SP], Expression, SP, (A,S), SP, Variable ;
type Unwind struct {
	Expression *Expression `(("U"|"u") ("N"|"n") ("W"|"w") ("I"|"i") ("N"|"n") ("D"|"d")) @@`
	Variable   *Variable   `(("A"|"a") ("S"|"s")) @@`
}

// Merge
//
// EBNF Grammar:
//   Merge = (M,E,R,G,E), [SP], PatternPart, { SP, MergeAction } ;
type Merge struct {
	PatternPart  *PatternPart   `(("M"|"m") ("E"|"e") ("R"|"r") ("G"|"g") ("E"|"e")) @@`
	MergeActions []*MergeAction `@@*`
}

// MergeAction
//
// EBNF Grammar:
//   MergeAction = ((O,N), SP, (M,A,T,C,H), SP, Set)
//               | ((O,N), SP, (C,R,E,A,T,E), SP, Set)
//               ;
type MergeAction struct{}

// Create
//
// EBNF Grammar:
//   Create = (C,R,E,A,T,E), [SP], Pattern ;
type Create struct{}

// Set
//
// EBNF Grammar:
//   Set = (S,E,T), [SP], SetItem, { [SP], ',', [SP], SetItem } ;
type Set struct{}

// SetItem
//
// EBNF Grammar:
//   SetItem = (PropertyExpression, [SP], '=', [SP], Expression)
//           | (Variable, [SP], '=', [SP], Expression)
//           | (Variable, [SP], '+=', [SP], Expression)
//           | (Variable, [SP], NodeLabels)
//           ;
type SetItem struct{}

// Delete
//
// EBNF Grammar:
//   Delete = [(D,E,T,A,C,H), SP], (D,E,L,E,T,E), [SP], Expression, { [SP], ',', [SP], Expression } ;
type Delete struct{}

// Remove
//
// EBNF Grammar:
//   Remove = (R,E,M,O,V,E), SP, RemoveItem, { [SP], ',', [SP], RemoveItem } ;
type Remove struct{}

// RemoveItem
//
// EBNF Grammar:
//   RemoveItem = (Variable, NodeLabels)
//              | PropertyExpression
//              ;
type RemoveItem struct{}

// InQueryCall
//
// EBNF Grammar:
//   InQueryCall = (C,A,L,L), SP, ExplicitProcedureInvocation, [[SP], (Y,I,E,L,D), SP, YieldItems] ;
type InQueryCall struct{}

// StandaloneCall
//
// EBNF Grammar:
//   StandaloneCall = (C,A,L,L), SP, (ExplicitProcedureInvocation | ImplicitProcedureInvocation), [[SP], (Y,I,E,L,D), SP, ('*' | YieldItems)] ;
type StandaloneCall struct{}

// YieldItems
//
// EBNF Grammar:
//   YieldItems = YieldItem, { [SP], ',', [SP], YieldItem }, [[SP], Where] ;
type YieldItems struct{}

// YieldItem
//
// EBNF Grammar:
//   YieldItem = [ProcedureResultField, SP, (A,S), SP], Variable ;
type YieldItem struct{}

// With
//
// EBNF Grammar:
//   With = (W,I,T,H), ProjectionBody, [[SP], Where] ;
type With struct{}

// Return
//
// EBNF Grammar:
//   Return = (R,E,T,U,R,N), ProjectionBody ;
type Return struct{}

// ProjectionBody
//
// EBNF Grammar:
//   ProjectionBody = [[SP], (D,I,S,T,I,N,C,T)], SP, ProjectionItems, [SP, Order], [SP, Skip], [SP, Limit] ;
type ProjectionBody struct{}

// ProjectionItems
//
// EBNF Grammar:
//   ProjectionItems = ('*', { [SP], ',', [SP], ProjectionItem })
//                   | (ProjectionItem, { [SP], ',', [SP], ProjectionItem })
//                   ;
type ProjectionItems struct{}

// ProjectionItem
//
// EBNF Grammar:
//   ProjectionItem = (Expression, SP, (A,S), SP, Variable)
//                  | Expression
//                  ;
type ProjectionItem struct{}

// Order
//
// EBNF Grammar:
//   Order = (O,R,D,E,R), SP, (B,Y), SP, SortItem, { ',', [SP], SortItem } ;
type Order struct{}

// Skip
//
// EBNF Grammar:
//   Skip = (S,K,I,P), SP, Expression ;
type Skip struct{}

// Limit
//
// EBNF Grammar:
//   Limit = (L,I,M,I,T), SP, Expression ;
type Limit struct{}

// SortItem
//
// EBNF Grammar:
//   SortItem = Expression, [[SP], ((A,S,C,E,N,D,I,N,G) | (A,S,C) | (D,E,S,C,E,N,D,I,N,G) | (D,E,S,C))] ;
type SortItem struct{}

// Where
//
// EBNF Grammar:
//   Where = (W,H,E,R,E), SP, Expression ;
type Where struct{}

// Pattern
//
// EBNF Grammar:
//   Pattern = PatternPart, { [SP], ',', [SP], PatternPart } ;
type Pattern struct{}

// PatternPart
//
// EBNF Grammar:
//   PatternPart = (Variable, [SP], '=', [SP], AnonymousPatternPart)
//               | AnonymousPatternPart
//               ;
type PatternPart struct{}

// AnonymousPatternPart
//
// EBNF Grammar:
//   AnonymousPatternPart = PatternElement ;
type AnonymousPatternPart struct{}

// PatternElement
//
// EBNF Grammar:
//   PatternElement = (NodePattern, { [SP], PatternElementChain })
//                  | ('(', PatternElement, ')')
//                  ;
type PatternElement struct{}

// NodePattern
//
// EBNF Grammar:
//   NodePattern = '(', [SP], [Variable, [SP]], [NodeLabels, [SP]], [Properties, [SP]], ')' ;
type NodePattern struct{}

// PatternElementChain
//
// EBNF Grammar:
//   PatternElementChain = RelationshipPattern, [SP], NodePattern ;
type PatternElementChain struct{}

// RelationshipPattern
//
// EBNF Grammar:
//   RelationshipPattern = (LeftArrowHead, [SP], Dash, [SP], [RelationshipDetail], [SP], Dash, [SP], RightArrowHead)
//                       | (LeftArrowHead, [SP], Dash, [SP], [RelationshipDetail], [SP], Dash)
//                       | (Dash, [SP], [RelationshipDetail], [SP], Dash, [SP], RightArrowHead)
//                       | (Dash, [SP], [RelationshipDetail], [SP], Dash)
//                       ;
type RelationshipPattern struct{}

// RelationshipDetail
//
// EBNF Grammar:
//   RelationshipDetail = '[', [SP], [Variable, [SP]], [RelationshipTypes, [SP]], [RangeLiteral], [Properties, [SP]], ']' ;
type RelationshipDetail struct{}

// Properties
//
// EBNF Grammar:
//   Properties = MapLiteral
//              | Parameter
//              ;
type Properties struct{}

// RelationshipTypes
//
// EBNF Grammar:
//   RelationshipTypes = ':', [SP], RelTypeName, { [SP], '|', [':'], [SP], RelTypeName } ;
type RelationshipTypes struct{}

// NodeLabels
//
// EBNF Grammar:
//   NodeLabels = NodeLabel, { [SP], NodeLabel } ;
type NodeLabels struct{}

// NodeLabel
//
// EBNF Grammar:
//   NodeLabel = ':', [SP], LabelName ;
type NodeLabel struct{}

// RangeLiteral
//
// EBNF Grammar:
//   RangeLiteral = '*', [SP], [IntegerLiteral, [SP]], ['..', [SP], [IntegerLiteral, [SP]]] ;
type RangeLiteral struct{}

// LabelName
//
// EBNF Grammar:
//   LabelName = SchemaName ;
type LabelName struct{}

// RelTypeName
//
// EBNF Grammar:
//   RelTypeName = SchemaName ;
type RelTypeName struct{}

// Expression
//
// EBNF Grammar:
//   Expression = OrExpression ;
type Expression struct{}

// OrExpression
//
// EBNF Grammar:
//   OrExpression = XorExpression, { SP, (O,R), SP, XorExpression } ;
type OrExpression struct{}

// XorExpression
//
// EBNF Grammar:
//   XorExpression = AndExpression, { SP, (X,O,R), SP, AndExpression } ;
type XorExpression struct{}

// AndExpression
//
// EBNF Grammar:
//   AndExpression = NotExpression, { SP, (A,N,D), SP, NotExpression } ;
type AndExpression struct{}

// NotExpression
//
// EBNF Grammar:
//   NotExpression = { (N,O,T), [SP] }, ComparisonExpression ;
type NotExpression struct{}

// ComparisonExpression
//
// EBNF Grammar:
//   ComparisonExpression = AddOrSubtractExpression, { [SP], PartialComparisonExpression } ;
type ComparisonExpression struct{}

// AddOrSubtractExpression
//
// EBNF Grammar:
//   AddOrSubtractExpression = MultiplyDivideModuloExpression, { ([SP], '+', [SP], MultiplyDivideModuloExpression) | ([SP], '-', [SP], MultiplyDivideModuloExpression) } ;
type AddOrSubtractExpression struct{}

// MultiplyDivideModuloExpression
//
// EBNF Grammar:
//   MultiplyDivideModuloExpression = PowerOfExpression, { ([SP], '*', [SP], PowerOfExpression) | ([SP], '/', [SP], PowerOfExpression) | ([SP], '%', [SP], PowerOfExpression) } ;
type MultiplyDivideModuloExpression struct{}

// PowerOfExpression
//
// EBNF Grammar:
//   PowerOfExpression = UnaryAddOrSubtractExpression, { [SP], '^', [SP], UnaryAddOrSubtractExpression } ;
type PowerOfExpression struct{}

// UnaryAddOrSubtractExpression
//
// EBNF Grammar:
//   UnaryAddOrSubtractExpression = { ('+' | '-'), [SP] }, StringListNullOperatorExpression ;
type UnaryAddOrSubtractExpression struct{}

// StringListNullOperatorExpression
//
// EBNF Grammar:
//   StringListNullOperatorExpression = PropertyOrLabelsExpression, { StringOperatorExpression | ListOperatorExpression | NullOperatorExpression } ;
type StringListNullOperatorExpression struct{}

// ListOperatorExpression
//
// EBNF Grammar:
//   ListOperatorExpression = (SP, (I,N), [SP], PropertyOrLabelsExpression)
//                          | ([SP], '[', Expression, ']')
//                          | ([SP], '[', [Expression], '..', [Expression], ']')
//                          ;
type ListOperatorExpression struct{}

// StringOperatorExpression
//
// EBNF Grammar:
//   StringOperatorExpression = ((SP, (S,T,A,R,T,S), SP, (W,I,T,H)) | (SP, (E,N,D,S), SP, (W,I,T,H)) | (SP, (C,O,N,T,A,I,N,S))), [SP], PropertyOrLabelsExpression ;
type StringOperatorExpression struct{}

// NullOperatorExpression
//
// EBNF Grammar:
//   NullOperatorExpression = (SP, (I,S), SP, (N,U,L,L))
//                          | (SP, (I,S), SP, (N,O,T), SP, (N,U,L,L))
//                          ;
type NullOperatorExpression struct{}

// PropertyOrLabelsExpression
//
// EBNF Grammar:
//   PropertyOrLabelsExpression = Atom, { [SP], PropertyLookup }, [[SP], NodeLabels] ;
type PropertyOrLabelsExpression struct{}

// Atom
//
// EBNF Grammar:
//   Atom = Literal
//        | Parameter
//        | CaseExpression
//        | ((C,O,U,N,T), [SP], '(', [SP], '*', [SP], ')')
//        | ListComprehension
//        | PatternComprehension
//        | ((A,L,L), [SP], '(', [SP], FilterExpression, [SP], ')')
//        | ((A,N,Y), [SP], '(', [SP], FilterExpression, [SP], ')')
//        | ((N,O,N,E), [SP], '(', [SP], FilterExpression, [SP], ')')
//        | ((S,I,N,G,L,E), [SP], '(', [SP], FilterExpression, [SP], ')')
//        | RelationshipsPattern
//        | ParenthesizedExpression
//        | FunctionInvocation
//        | ExistentialSubquery
//        | Variable
//        ;
type Atom struct{}

// Literal
//
// EBNF Grammar:
//   Literal = NumberLiteral
//           | StringLiteral
//           | BooleanLiteral
//           | (N,U,L,L)
//           | MapLiteral
//           | ListLiteral
//           ;
type Literal struct{}

// BooleanLiteral
//
// EBNF Grammar:
//   BooleanLiteral = (T,R,U,E)
//                  | (F,A,L,S,E)
//                  ;
type BooleanLiteral struct{}

// ListLiteral
//
// EBNF Grammar:
//   ListLiteral = '[', [SP], [Expression, [SP], { ',', [SP], Expression, [SP] }], ']' ;
type ListLiteral struct{}

// PartialComparisonExpression
//
// EBNF Grammar:
//   PartialComparisonExpression = ('=', [SP], AddOrSubtractExpression)
//                               | ('<>', [SP], AddOrSubtractExpression)
//                               | ('<', [SP], AddOrSubtractExpression)
//                               | ('>', [SP], AddOrSubtractExpression)
//                               | ('<=', [SP], AddOrSubtractExpression)
//                               | ('>=', [SP], AddOrSubtractExpression)
//                               ;
type PartialComparisonExpression struct{}

// ParenthesizedExpression
//
// EBNF Grammar:
//   ParenthesizedExpression = '(', [SP], Expression, [SP], ')' ;
type ParenthesizedExpression struct{}

// RelationshipsPattern
//
// EBNF Grammar:
//   RelationshipsPattern = NodePattern, { [SP], PatternElementChain }- ;
type RelationshipsPattern struct{}

// FilterExpression
//
// EBNF Grammar:
//   FilterExpression = IdInColl, [[SP], Where] ;
type FilterExpression struct{}

// IdInColl
//
// EBNF Grammar:
//   IdInColl = Variable, SP, (I,N), SP, Expression ;
type IdInColl struct{}

// FunctionInvocation
//
// EBNF Grammar:
//   FunctionInvocation = FunctionName, [SP], '(', [SP], [(D,I,S,T,I,N,C,T), [SP]], [Expression, [SP], { ',', [SP], Expression, [SP] }], ')' ;
type FunctionInvocation struct{}

// FunctionName
//
// EBNF Grammar:
//   FunctionName = Namespace, SymbolicName ;
type FunctionName struct{}

// ExistentialSubquery
//
// EBNF Grammar:
//   ExistentialSubquery = (E,X,I,S,T,S), [SP], '{', [SP], (RegularQuery | (Pattern, [[SP], Where])), [SP], '}' ;
type ExistentialSubquery struct{}

// ExplicitProcedureInvocation
//
// EBNF Grammar:
//   ExplicitProcedureInvocation = ProcedureName, [SP], '(', [SP], [Expression, [SP], { ',', [SP], Expression, [SP] }], ')' ;
type ExplicitProcedureInvocation struct{}

// ImplicitProcedureInvocation
//
// EBNF Grammar:
//   ImplicitProcedureInvocation = ProcedureName ;
type ImplicitProcedureInvocation struct{}

// ProcedureResultField
//
// EBNF Grammar:
//   ProcedureResultField = SymbolicName ;
type ProcedureResultField struct{}

// ProcedureName
//
// EBNF Grammar:
//   ProcedureName = Namespace, SymbolicName ;
type ProcedureName struct{}

// Namespace
//
// EBNF Grammar:
//   Namespace = { SymbolicName, '.' } ;
type Namespace struct{}

// ListComprehension
//
// EBNF Grammar:
//   ListComprehension = '[', [SP], FilterExpression, [[SP], '|', [SP], Expression], [SP], ']' ;
type ListComprehension struct{}

// PatternComprehension
//
// EBNF Grammar:
//   PatternComprehension = '[', [SP], [Variable, [SP], '=', [SP]], RelationshipsPattern, [SP], [Where, [SP]], '|', [SP], Expression, [SP], ']' ;
type PatternComprehension struct{}

// PropertyLookup
//
// EBNF Grammar:
//   PropertyLookup = '.', [SP], (PropertyKeyName) ;
type PropertyLookup struct{}

// CaseExpression
//
// EBNF Grammar:
//   CaseExpression = (((C,A,S,E), { [SP], CaseAlternative }-) | ((C,A,S,E), [SP], Expression, { [SP], CaseAlternative }-)), [[SP], (E,L,S,E), [SP], Expression], [SP], (E,N,D) ;
type CaseExpression struct{}

// CaseAlternative
//
// EBNF Grammar:
//   CaseAlternative = (W,H,E,N), [SP], Expression, [SP], (T,H,E,N), [SP], Expression ;
type CaseAlternative struct{}

// Variable
//
// EBNF Grammar:
//   Variable = SymbolicName ;
type Variable struct{}

// StringLiteral
//
// EBNF Grammar:
//   StringLiteral = ('"', { ANY - ('"' | '\') | EscapedChar }, '"')
//                 | ("'", { ANY - ("'" | '\') | EscapedChar }, "'")
//                 ;
type StringLiteral struct{}

// EscapedChar
//
// EBNF Grammar:
//   EscapedChar = '\', ('\' | "'" | '"' | (B) | (F) | (N) | (R) | (T) | ((U), 4 * HexDigit) | ((U), 8 * HexDigit)) ;
type EscapedChar struct{}

// NumberLiteral
//
// EBNF Grammar:
//   NumberLiteral = DoubleLiteral
//                 | IntegerLiteral
//                 ;
type NumberLiteral struct{}

// MapLiteral
//
// EBNF Grammar:
//   MapLiteral = '{', [SP], [PropertyKeyName, [SP], ':', [SP], Expression, [SP], { ',', [SP], PropertyKeyName, [SP], ':', [SP], Expression, [SP] }], '}' ;
type MapLiteral struct{}

// Parameter
//
// EBNF Grammar:
//   Parameter = '$', (SymbolicName | DecimalInteger) ;
type Parameter struct{}

// PropertyExpression
//
// EBNF Grammar:
//   PropertyExpression = Atom, { [SP], PropertyLookup }- ;
type PropertyExpression struct{}

// PropertyKeyName
//
// EBNF Grammar:
//   PropertyKeyName = SchemaName ;
type PropertyKeyName struct{}

// IntegerLiteral
//
// EBNF Grammar:
//   IntegerLiteral = HexInteger
//                  | OctalInteger
//                  | DecimalInteger
//                  ;
type IntegerLiteral struct{}

// HexInteger
//
// EBNF Grammar:
//   HexInteger = '0x', { HexDigit }- ;
type HexInteger struct{}

// DecimalInteger
//
// EBNF Grammar:
//   DecimalInteger = ZeroDigit
//                  | (NonZeroDigit, { Digit })
//                  ;
type DecimalInteger struct{}

// OctalInteger
//
// EBNF Grammar:
//   OctalInteger = ZeroDigit, { OctDigit }- ;
type OctalInteger struct{}

// HexLetter
//
// EBNF Grammar:
//   HexLetter = (A)
//             | (B)
//             | (C)
//             | (D)
//             | (E)
//             | (F)
//             ;
type HexLetter struct{}

// HexDigit
//
// EBNF Grammar:
//   HexDigit = Digit
//            | HexLetter
//            ;
type HexDigit struct{}

// Digit
//
// EBNF Grammar:
//   Digit = ZeroDigit
//         | NonZeroDigit
//         ;
type Digit struct{}

// NonZeroDigit
//
// EBNF Grammar:
//   NonZeroDigit = NonZeroOctDigit
//                | '8'
//                | '9'
//                ;
type NonZeroDigit struct{}

// NonZeroOctDigit
//
// EBNF Grammar:
//   NonZeroOctDigit = '1'
//                   | '2'
//                   | '3'
//                   | '4'
//                   | '5'
//                   | '6'
//                   | '7'
//                   ;
type NonZeroOctDigit struct{}

// OctDigit
//
// EBNF Grammar:
//   OctDigit = ZeroDigit
//            | NonZeroOctDigit
//            ;
type OctDigit struct{}

// ZeroDigit
//
// EBNF Grammar:
//   ZeroDigit = '0' ;
type ZeroDigit struct{}

// DoubleLiteral
//
// EBNF Grammar:
//   DoubleLiteral = ExponentDecimalReal
//                 | RegularDecimalReal
//                 ;
type DoubleLiteral struct{}

// ExponentDecimalReal
//
// EBNF Grammar:
//   ExponentDecimalReal = ({ Digit }- | ({ Digit }-, '.', { Digit }-) | ('.', { Digit }-)), (E), ['-'], { Digit }- ;
type ExponentDecimalReal struct{}

// RegularDecimalReal
//
// EBNF Grammar:
//   RegularDecimalReal = { Digit }, '.', { Digit }- ;
type RegularDecimalReal struct{}

// SchemaName
//
// EBNF Grammar:
//   SchemaName = SymbolicName
//              | ReservedWord
//              ;
type SchemaName struct{}

// ReservedWord
//
// EBNF Grammar:
//   ReservedWord = (A,L,L)
//                | (A,S,C)
//                | (A,S,C,E,N,D,I,N,G)
//                | (B,Y)
//                | (C,R,E,A,T,E)
//                | (D,E,L,E,T,E)
//                | (D,E,S,C)
//                | (D,E,S,C,E,N,D,I,N,G)
//                | (D,E,T,A,C,H)
//                | (E,X,I,S,T,S)
//                | (L,I,M,I,T)
//                | (M,A,T,C,H)
//                | (M,E,R,G,E)
//                | (O,N)
//                | (O,P,T,I,O,N,A,L)
//                | (O,R,D,E,R)
//                | (R,E,M,O,V,E)
//                | (R,E,T,U,R,N)
//                | (S,E,T)
//                | (S,K,I,P)
//                | (W,H,E,R,E)
//                | (W,I,T,H)
//                | (U,N,I,O,N)
//                | (U,N,W,I,N,D)
//                | (A,N,D)
//                | (A,S)
//                | (C,O,N,T,A,I,N,S)
//                | (D,I,S,T,I,N,C,T)
//                | (E,N,D,S)
//                | (I,N)
//                | (I,S)
//                | (N,O,T)
//                | (O,R)
//                | (S,T,A,R,T,S)
//                | (X,O,R)
//                | (F,A,L,S,E)
//                | (T,R,U,E)
//                | (N,U,L,L)
//                | (C,O,N,S,T,R,A,I,N,T)
//                | (D,O)
//                | (F,O,R)
//                | (R,E,Q,U,I,R,E)
//                | (U,N,I,Q,U,E)
//                | (C,A,S,E)
//                | (W,H,E,N)
//                | (T,H,E,N)
//                | (E,L,S,E)
//                | (E,N,D)
//                | (M,A,N,D,A,T,O,R,Y)
//                | (S,C,A,L,A,R)
//                | (O,F)
//                | (A,D,D)
//                | (D,R,O,P)
//                ;
type ReservedWord struct{}

// SymbolicName
//
// EBNF Grammar:
//   SymbolicName = UnescapedSymbolicName
//                | EscapedSymbolicName
//                | HexLetter
//                | (C,O,U,N,T)
//                | (F,I,L,T,E,R)
//                | (E,X,T,R,A,C,T)
//                | (A,N,Y)
//                | (N,O,N,E)
//                | (S,I,N,G,L,E)
//                ;
type SymbolicName struct{}

// UnescapedSymbolicName
//
// EBNF Grammar:
//   UnescapedSymbolicName = IdentifierStart, { IdentifierPart } ;
type UnescapedSymbolicName struct{}

// IdentifierStart
//
// EBNF Grammar:
//   IdentifierStart = ID_Start
//                   | Pc
//                   ;
type IdentifierStart struct{}

// IdentifierPart
//
// EBNF Grammar:
//   IdentifierPart = ID_Continue
//                  | Sc
//                  ;
type IdentifierPart struct{}

// EscapedSymbolicName
//
// EBNF Grammar:
//   EscapedSymbolicName = { '`', { ANY - ('`') }, '`' }- ;
type EscapedSymbolicName struct{}

// SP
//
// EBNF Grammar:
//   SP = { whitespace }- ;
type SP struct{}

// whitespace
//
// EBNF Grammar:
//   whitespace = SPACE
//              | TAB
//              | LF
//              | VT
//              | FF
//              | CR
//              | FS
//              | GS
//              | RS
//              | US
//              | ' '
//              | '᠎'
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | ' '
//              | '　'
//              | ' '
//              | ' '
//              | ' '
//              | Comment
//              ;
type whitespace struct{}

// Comment
//
// EBNF Grammar:
//   Comment = ('/*', { ANY - ('*') | ('*', ANY - ('/')) }, '*/')
//           | ('//', { ANY - (LF | CR) }, [CR], (LF | EOI))
//           ;
type Comment struct{}

// LeftArrowHead
//
// EBNF Grammar:
//   LeftArrowHead = '<'
//                 | '⟨'
//                 | '〈'
//                 | '﹤'
//                 | '＜'
//                 ;
type LeftArrowHead struct{}

// RightArrowHead
//
// EBNF Grammar:
//   RightArrowHead = '>'
//                  | '⟩'
//                  | '〉'
//                  | '﹥'
//                  | '＞'
//                  ;
type RightArrowHead struct{}

// Dash
//
// EBNF Grammar:
//   Dash = '-'
//        | '­'
//        | '‐'
//        | '‑'
//        | '‒'
//        | '–'
//        | '—'
//        | '―'
//        | '−'
//        | '﹘'
//        | '﹣'
//        | '－'
//        ;
type Dash struct{}

// A
//
// EBNF Grammar:
//   A = 'A' | 'a' ;
type A struct{}

// B
//
// EBNF Grammar:
//   B = 'B' | 'b' ;
type B struct{}

// C
//
// EBNF Grammar:
//   C = 'C' | 'c' ;
type C struct{}

// D
//
// EBNF Grammar:
//   D = 'D' | 'd' ;
type D struct{}

// E
//
// EBNF Grammar:
//   E = 'E' | 'e' ;
type E struct{}

// F
//
// EBNF Grammar:
//   F = 'F' | 'f' ;
type F struct{}

// G
//
// EBNF Grammar:
//   G = 'G' | 'g' ;
type G struct{}

// H
//
// EBNF Grammar:
//   H = 'H' | 'h' ;
type H struct{}

// I
//
// EBNF Grammar:
//   I = 'I' | 'i' ;
type I struct{}

// K
//
// EBNF Grammar:
//   K = 'K' | 'k' ;
type K struct{}

// L
//
// EBNF Grammar:
//   L = 'L' | 'l' ;
type L struct{}

// M
//
// EBNF Grammar:
//   M = 'M' | 'm' ;
type M struct{}

// N
//
// EBNF Grammar:
//   N = 'N' | 'n' ;
type N struct{}

// O
//
// EBNF Grammar:
//   O = 'O' | 'o' ;
type O struct{}

// P
//
// EBNF Grammar:
//   P = 'P' | 'p' ;
type P struct{}

// Q
//
// EBNF Grammar:
//   Q = 'Q' | 'q' ;
type Q struct{}

// R
//
// EBNF Grammar:
//   R = 'R' | 'r' ;
type R struct{}

// S
//
// EBNF Grammar:
//   S = 'S' | 's' ;
type S struct{}

// T
//
// EBNF Grammar:
//   T = 'T' | 't' ;
type T struct{}

// U
//
// EBNF Grammar:
//   U = 'U' | 'u' ;
type U struct{}

// V
//
// EBNF Grammar:
//   V = 'V' | 'v' ;
type V struct{}

// W
//
// EBNF Grammar:
//   W = 'W' | 'w' ;
type W struct{}

// X
//
// EBNF Grammar:
//   X = 'X' | 'x' ;
type X struct{}

// Y
//
// EBNF Grammar:
//   Y = 'Y' | 'y' ;
type Y struct{}
