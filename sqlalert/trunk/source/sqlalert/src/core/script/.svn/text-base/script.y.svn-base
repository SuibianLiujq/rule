%{
package script
%}

%union {
	empty    interface{}
	value    interface{}
	token    Token
	list     []Token
	operator TokenType
}

%token<empty> Y_ERR
%token<value> Y_IDENT
%token<value> Y_VALUE
%token<value> Y_STR
%token<value> Y_INT
%token<value> Y_FLOAT
%token<value> Y_BOOL
%token<value> Y_NULL

%token<empty>    '=' '?' ':' ';' ',' '$' '(' ')' '[' ']' '{' '}'
%token<operator> '+' '-' '*' '/' '%' '<' '>' Y_EQ Y_NE Y_LE Y_GE Y_AND Y_OR Y_NOT Y_INC Y_DEC
%type<operator>  ADDSUB MULDIV ANDOR COMP

%token<empty> Y_IF
%token<empty> Y_ELSE
%token<empty> Y_FOR
%token<empty> Y_IN
%token<empty> Y_CONTINUE
%token<empty> Y_BREAK
%token<empty> Y_RETURN
%token<empty> Y_DEF
%token<empty> Y_INCLUDE
%token<empty> Y_IMPORT


%type<token> all
%type<token> expression
%type<token> statements

%type<token> stmt
%type<token> stmt_item
%type<token> stmt_single
%type<token> stmt_group
%type<list>  stmt_list
%type<list>  stmt_list_none
%type<token> stmt_if
%type<token> stmt_ifelse
%type<token> stmt_elseif
%type<list>  stmt_elseif_list
%type<token> stmt_else
%type<token> stmt_for
%type<token> stmt_for_iter
%type<token> stmt_for_start
%type<token> stmt_for_check
%type<token> stmt_for_next
%type<token> stmt_for_cond
%type<token> stmt_forin_iter
%type<token> stmt_def
%type<list>  stmt_def_args_list
%type<token> stmt_include
%type<token> stmt_import
%type<token> stmt_assign
%type<token> stmt_assign_left
%type<token> stmt_expr
%type<token> stmt_interrupt

%type<token> expr
%type<token> expr_all
%type<token> expr_agg
%type<token> expr_agg_level0
%type<token> expr_agg_level1
%type<token> expr_agg_level2
%type<token> expr_level0
%type<token> expr_level1
%type<token> expr_level2
%type<token> expr_level3
%type<token> expr_iterable

%type<token> expr_base
%type<token> expr_ident
%type<token> expr_str
%type<token> expr_num
%type<token> expr_int
%type<token> expr_float
%type<token> expr_bool
%type<token> expr_null

%type<token> expr_listdict
%type<token> expr_list
%type<token> expr_list_item
%type<list>  expr_list_item_list
%type<token> expr_dict
%type<token> expr_dict_pair
%type<token> expr_dict_pair_key
%type<token> expr_dict_pair_value
%type<list>  expr_dict_pair_list

%type<token> expr_arithmetic
%type<token> expr_addsub
%type<token> expr_addsub_left
%type<token> expr_addsub_right
%type<token> expr_muldiv
%type<token> expr_muldiv_left
%type<token> expr_muldiv_right
%type<token> expr_modular
%type<token> expr_modular_left
%type<token> expr_modular_right
%type<token> expr_negative
%type<token> expr_negative_right

%type<token> expr_comp
%type<token> expr_comp_left
%type<token> expr_comp_right

%type<token> expr_logical
%type<token> expr_logical_andor
%type<token> expr_logical_andor_left
%type<token> expr_logical_andor_right
%type<token> expr_logical_not
%type<token> expr_logical_not_right

%type<token> expr_in
%type<token> expr_in_left
%type<token> expr_in_right

%type<token> expr_index
%type<token> expr_index_left
%type<token> expr_index_right

%type<token> expr_self
%type<token> expr_self_left

%type<token> expr_func
%type<token> expr_func_args
%type<list>  expr_func_args_list

%type<token> expr_cond
%type<token> expr_cond_left
%type<token> expr_cond_middle
%type<token> expr_cond_right

%%
all: expression { setToken(yylex, $1) } | statements { setToken(yylex, $1) };
statements: stmt_list { $$ = (&TokenStmts{}).Init($1) };
expression: expr_all;


// statements.
stmt_item:   stmt_assign | stmt_expr   | stmt_interrupt | stmt_include | stmt_import;
stmt_single: stmt_item ';' { $$ = $1 } | stmt_single ';' { $$ = $1 };
stmt_group:  stmt_ifelse | stmt_for    | stmt_def;
stmt:        stmt_single | stmt_group;
stmt_list:   stmt { $$ = []Token{$1} } | stmt_list stmt { $$ = append($1, $2) } | '{' stmt_list '}' { $$ = $2 };

stmt_list_none: { $$ = []Token{} } | stmt_list { $$ = $1 };

stmt_for_start: stmt_expr | stmt_assign;
stmt_for_check: stmt_expr;
stmt_for_next:  stmt_expr | stmt_assign;
stmt_for_iter:
      '('               ')'                               { $$ = (&TokenForIter{}).Init(nil, nil, nil) }
	| '(' stmt_for_iter ')'                               { $$ = $2 }
	| stmt_for_start ';' stmt_for_check ';' stmt_for_next { $$ = (&TokenForIter{}).Init($1, $3, $5) }
;

stmt_forin_iter:
	  expr_ident ',' expr_ident Y_IN expr_iterable { $$ = (&TokenForinIter{}).Init($1, $3, $5) }
	| '(' stmt_forin_iter ')'                    { $$ = $2 }
;

stmt_for_cond:
	stmt_expr {
		if $1.(*TokenExpr).Token.Type() == T_IN {
			t := $1.(*TokenExpr).Token.(*TokenIn)
			$$ = (&TokenForinIter{}).Init(nil, t.Key, t.Object)
		} else {
			$$ = (&TokenForIter{}).Init(nil, $1.(*TokenExpr).Token, nil)
		}
	}
	| stmt_for_iter   { $$ = $1 }
	| stmt_forin_iter { $$ = $1 }
;

stmt_for: Y_FOR stmt_for_cond '{' stmt_list_none '}' {
	if $2.Type() == T_FOR_ITER {
		$$ = (&TokenFor{}).Init($2, $4)
	} else {
		$$ = (&TokenForin{}).Init($2, $4)
	}
};

stmt_if:          Y_IF         expr_all '{' stmt_list_none '}' { $$ = (&TokenIf{}).Init($2, $4) };
stmt_elseif:      Y_ELSE Y_IF  expr_all '{' stmt_list_none '}' { $$ = (&TokenElseIf{}).Init($3, $5) };
stmt_else:        { $$ = nil } | Y_ELSE '{' stmt_list_none '}' { $$ = (&TokenElse{}).Init($3) };
stmt_elseif_list: stmt_elseif { $$ = []Token{$1} } | stmt_elseif_list stmt_elseif { $$ = append($1, $2) };
stmt_ifelse:
	  stmt_if                  stmt_else { $$ = (&TokenIfElse{}).Init($1, nil, $2) }
	| stmt_if stmt_elseif_list stmt_else { $$ = (&TokenIfElse{}).Init($1, $2, $3) };
;

stmt_def: Y_DEF expr_ident '(' stmt_def_args_list ')' '{' stmt_list_none '}' { $$ = (&TokenDefine{}).Init($2, $4, $7) };
stmt_def_args_list:
	                                    { $$ = []Token{} }
	| expr_ident                        { $$ = []Token{$1} }
	| stmt_def_args_list ',' expr_ident { $$ = append($1, $3) }
;

stmt_include:     Y_INCLUDE expr_str    { $$ = (&TokenInclude{}).Init($2) };
stmt_import:      Y_IMPORT  expr_str    { $$ = (&TokenImport{}).Init($2)  };
stmt_assign_left: expr_ident | expr_index;
stmt_assign:      stmt_assign_left '=' expression { $$ = (&TokenAssign{}).Init($1, $3) };
stmt_expr:        expression                      { $$ = (&TokenExpr{}).Init($1) };
stmt_interrupt:
	  Y_CONTINUE        { $$ = (&TokenContinue{}).Init() }
	| Y_BREAK           { $$ = (&TokenBreak{}).Init() }
	| Y_RETURN          { $$ = (&TokenReturn{}).Init(nil) }
	| Y_RETURN expr_all { $$ = (&TokenReturn{}).Init($2) }
;


// expressions;
expr:     expr_level3;
expr_all: expr_agg | expr;
expr_agg: '(' expr ')' { $$ = $2 }  | '(' expr_agg ')' { $$ = $2 };

expr_level0:     expr_base   | expr_func       | expr_self     | expr_negative | expr_index;
expr_level1:     expr_level0 | expr_arithmetic | expr_listdict | expr_in;
expr_level2:     expr_level1 | expr_comp       | expr_logical;
expr_level3:     expr_level2 | expr_cond;

expr_agg_level0:     expr_agg | expr_level0;
expr_agg_level1:     expr_agg | expr_level1;
expr_agg_level2:     expr_agg | expr_level2;

expr_listdict:  expr_list       | expr_dict;
expr_iterable:  expr_agg        | expr_ident | expr_index | expr_func | expr_listdict;


// operatores.
ADDSUB: '+'   | '-';
MULDIV: '*'   | '/';
ANDOR:  Y_AND | Y_OR;
COMP:   '<'  | '>' | Y_LE | Y_GE | Y_EQ | Y_NE;


// level 0;
expr_base:  expr_ident | expr_str | expr_num | expr_bool | expr_null;
expr_num:   expr_int   | expr_float;
expr_ident: Y_IDENT { $$ = (&TokenIdent{}).Init($1, getSrc(yylex)) };
expr_str:   Y_STR   { $$ = (&TokenStr{}).Init($1,   getSrc(yylex)) };
expr_int:   Y_INT   { $$ = (&TokenInt{}).Init($1,   getSrc(yylex)) };
expr_float: Y_FLOAT { $$ = (&TokenFloat{}).Init($1, getSrc(yylex)) };
expr_bool:  Y_BOOL  { $$ = (&TokenBool{}).Init($1,  getSrc(yylex)) };
expr_null:  Y_NULL  { $$ = (&TokenNull{}).Init(     getSrc(yylex)) };

expr_func:      expr_ident '(' expr_func_args_list ')' { $$ = (&TokenFunc{}).Init($1, $3) };
expr_func_args: expr_all;
expr_func_args_list:
	                                         { $$ = []Token{} }
	| expr_func_args                         { $$ = []Token{$1} }
	| expr_func_args_list ',' expr_func_args { $$ = append($1, $3) };

expr_self_left: expr_ident | expr_index | expr_self;
expr_self:
	  expr_self_left Y_INC { $$ = (&TokenOper{}).Init($1, $2, nil) }
	| expr_self_left Y_DEC { $$ = (&TokenOper{}).Init($1, $2, nil) }
;

expr_negative_right: expr_agg_level0;
expr_negative:       '-' expr_negative_right { $$ = (&TokenOper{}).Init(nil, $1, $2) };


// level 1;
expr_index_left:  expr_iterable;
expr_index_right: expr_agg_level0 | expr_arithmetic;
expr_index:       expr_index_left '[' expr_index_right ']' { $$ = (&TokenIndex{}).Init($1, $3) };

expr_list:      '[' expr_list_item_list ']' { $$ = (&TokenList{}).Init($2) };
expr_list_item: expr_all;
expr_list_item_list:
	                                         { $$ = []Token{} }
	| expr_list_item                         { $$ = []Token{$1} }
	| expr_list_item_list ',' expr_list_item { $$ = append($1, $3) };

expr_dict:            '{' expr_dict_pair_list '}'                 { $$ = (&TokenDict{}).Init($2) }
expr_dict_pair:       expr_dict_pair_key ':' expr_dict_pair_value { $$ = (&TokenPair{}).Init($1, $3) }
expr_dict_pair_key:   expr_ident | expr_str;
expr_dict_pair_value: expr_all;
expr_dict_pair_list:
	                                         { $$ = []Token{} }
	| expr_dict_pair                         { $$ = []Token{$1} }
	| expr_dict_pair_list ',' expr_dict_pair { $$ = append($1, $3) }
;

expr_arithmetic:     expr_addsub     | expr_muldiv  | expr_modular;
expr_addsub_left:    expr_agg_level0 | expr_arithmetic;
expr_addsub_right:   expr_agg_level0 | expr_muldiv  | expr_modular;
expr_muldiv_left:    expr_agg_level0 | expr_muldiv;
expr_muldiv_right:   expr_agg_level0 | expr_modular;
expr_modular_left:   expr_agg_level0 | expr_modular;
expr_modular_right:  expr_agg_level0;
expr_addsub:         expr_addsub_left  ADDSUB expr_addsub_right   { $$ = (&TokenOper{}).Init($1, $2, $3) };
expr_muldiv:         expr_muldiv_left  MULDIV expr_muldiv_right   { $$ = (&TokenOper{}).Init($1, $2, $3) };
expr_modular:        expr_modular_left '%'    expr_modular_right  { $$ = (&TokenOper{}).Init($1, $2, $3) };

expr_in_left:  expr_agg_level0 | expr_arithmetic;
expr_in_right: expr_agg_level0 | expr_listdict;
expr_in:       expr_in_left Y_IN expr_in_right { $$ = (&TokenIn{}).Init($1, $3) };


// level 2;
expr_comp_left:  expr_agg_level1;
expr_comp_right: expr_agg_level1;
expr_comp:       expr_comp_left COMP expr_comp_right { $$ = (&TokenComp{}).Init($1, $2, $3) };

expr_logical:             expr_logical_andor | expr_logical_not;
expr_logical_andor_left:  expr_agg_level1 | expr_comp | expr_logical;
expr_logical_andor_right: expr_agg_level1 | expr_comp | expr_logical_not;
expr_logical_not_right:   expr_agg_level1 | expr_comp | expr_logical_not;
expr_logical_andor:       expr_logical_andor_left ANDOR expr_logical_andor_right { $$ = (&TokenLogical{}).Init($1,  $2, $3) };
expr_logical_not:                                 Y_NOT expr_logical_not_right   { $$ = (&TokenLogical{}).Init(nil, $1, $2) };


// level 3
expr_cond_left:   expr_all;
expr_cond_middle: expr_all;
expr_cond_right:  expr_agg_level2;
expr_cond:        expr_cond_left '?' expr_cond_middle ':' expr_cond_right { $$ = (&TokenCond{}).Init($1, $3, $5) };

%%
