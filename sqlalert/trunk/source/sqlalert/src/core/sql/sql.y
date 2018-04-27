%{
package sql
%}

%union {
	empty    interface{}
	value    interface{}
	token    Token
	list     []Token
	operator TokenType
}

%token<empty> Y_ERR
%token<value> Y_IDENT  Y_STR   Y_INT   Y_FLOAT   Y_BOOL   Y_NULL
%token<empty> Y_SELECT Y_FROM  Y_WHERE Y_GROUP   Y_ORDER  Y_LIMIT Y_HAVING
%token<empty> Y_AS     Y_IN    Y_BY    Y_LIKE    Y_UNLIKE Y_DESC  Y_ASC    Y_INTERVAL Y_UNIQUE

%token<empty>    '=' '?' ':' ';' ',' '$' '(' ')' '[' ']' '{' '}'
%token<operator> '+' '-' '*' '/' '%' '<' '>' Y_EQ Y_NE Y_LE Y_GE Y_AND Y_OR Y_NOT Y_INC Y_DEC
%type<operator>  ADDSUB MULDIV ANDOR COMP

%type<token> all

%type<token> stmt
%type<list>  stmt_list
%type<token> stmt_item
%type<token> stmt_select
%type<token> stmt_select_item
%type<list>  stmt_select_list
%type<token> stmt_from
%type<token> stmt_from_item
%type<list>  stmt_from_list
%type<token> stmt_where
%type<token> stmt_whereby
%type<token> stmt_where_item
%type<token> stmt_groupby
%type<token> stmt_groupby_item
%type<list>  stmt_groupby_list
%type<token> stmt_order
%type<token> stmt_order_item
%type<operator> stmt_order_order
%type<token> stmt_orderby
%type<list>  stmt_orderby_list
%type<token> stmt_limit
%type<token> stmt_limit_item
%type<list>  stmt_limit_list
%type<token> stmt_having
%type<token> stmt_having_item
%type<token> stmt_as
%type<token> stmt_unique


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
%type<token> expr_var
%type<token> expr_ident
%type<token> expr_str
%type<token> expr_num
%type<token> expr_int
%type<token> expr_float
%type<token> expr_bool
%type<token> expr_null

%type<token> expr_star
%type<token> expr_numunit

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
all: stmt { setToken(yylex, $1) };

stmt:       stmt_select stmt_list { $$ = (&TokenStmts{}).Init($1, $2) };
stmt_item:  stmt_from | stmt_where | stmt_whereby | stmt_groupby | stmt_orderby | stmt_limit | stmt_having;
stmt_list:  stmt_item { $$ = []Token{$1} } | stmt_list stmt_item { $$ = append($1, $2) };

stmt_select_item: expr_all | stmt_as;
stmt_select_list: stmt_select_item { $$ = []Token{$1} } | stmt_select_list ',' stmt_select_item { $$ = append($1, $3) };
stmt_select:      Y_SELECT stmt_select_list { $$ = (&TokenSelect{}).Init($2) };

stmt_from_item: expr_ident | expr_str | expr_var;
stmt_from_list: stmt_from_item { $$ = []Token{$1} } | stmt_from_list ',' stmt_from_item { $$ = append($1, $3) };
stmt_from:      Y_FROM stmt_from_list { $$ = (&TokenFrom{}).Init($2) };

stmt_where_item: expr_all;
stmt_where:   Y_WHERE stmt_where_item { $$ = (&TokenWhere{}).Init($2) };
stmt_whereby: Y_WHERE Y_BY stmt_groupby_list { $$ = (&TokenWhereBy{}).Init($3) };

stmt_groupby_item: expr_ident | expr_str | stmt_as;
stmt_groupby_list: stmt_groupby_item { $$ = []Token{$1} } | stmt_groupby_list ',' stmt_groupby_item { $$ = append($1, $3) };
stmt_groupby:      Y_GROUP Y_BY stmt_groupby_list { $$ = (&TokenGroupBy{}).Init($3) };

stmt_order_item:   expr_ident | expr_str;
stmt_order_order:  { $$ = T_ILL } | Y_DESC { $$ = T_DESC } | Y_ASC { $$ = T_ASC };
stmt_order:        stmt_order_item stmt_order_order { $$ = (&TokenOrder{}).Init($1, $2) };
stmt_orderby_list: stmt_order { $$ = []Token{$1} } | stmt_orderby_list ',' stmt_order { $$ = append($1, $3) };
stmt_orderby:      Y_ORDER Y_BY stmt_orderby_list { $$ = (&TokenOrderBy{}).Init($3) };

stmt_limit_item: expr_num;
stmt_limit_list: stmt_limit_item { $$ = []Token{$1} } | stmt_limit_list ',' stmt_limit_item { $$ = append($1, $3) };
stmt_limit:      Y_LIMIT stmt_limit_list { $$ = (&TokenLimit{}).Init($2) };

stmt_having_item: expr_all;
stmt_having: Y_HAVING stmt_having_item { $$ = (&TokenHaving{}).Init($2) };

stmt_as:     expr_all Y_AS expr_ident { $$ = (&TokenAs{}).Init($1, $3) } | expr_all Y_AS expr_str { $$ = (&TokenAs{}).Init($1, $3) };
stmt_unique: Y_UNIQUE expr_all { $$ = (&TokenUnique{}).Init($2) };

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
expr_base:  expr_ident | expr_str | expr_num | expr_bool | expr_null | expr_star | expr_numunit | expr_var;
expr_num:   expr_int   | expr_float;
expr_ident: Y_IDENT { $$ = (&TokenIdent{}).Init($1) };
expr_str:   Y_STR   { $$ = (&TokenStr{}).Init($1) };
expr_int:   Y_INT   { $$ = (&TokenInt{}).Init($1) };
expr_float: Y_FLOAT { $$ = (&TokenFloat{}).Init($1) };
expr_bool:  Y_BOOL  { $$ = (&TokenBool{}).Init($1) };
expr_null:  Y_NULL  { $$ = (&TokenNull{}).Init() };
expr_star:    '*'                  { $$ = (&TokenStar{}).Init() };
expr_numunit:  expr_num expr_ident { $$ = (&TokenNumUnit{}).Init($1, $2) };
expr_var:
	  '$'     '(' expr_ident ')' { $$ = (&TokenVar{}).Init($3, false) }
	| '$' '$' '(' expr_ident ')' { $$ = (&TokenVar{}).Init($4, true) }
;

expr_func:      expr_ident '(' expr_func_args_list ')' { $$ = (&TokenFunc{}).Init($1, $3) };
expr_func_args: expr_all | stmt_unique | stmt_as;
expr_func_args_list:
	                                         { $$ = []Token{} }
	| expr_func_args                         { $$ = []Token{$1} }
	| expr_func_args_list ',' expr_func_args { $$ = append($1, $3) }
;

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
expr_addsub_right:   expr_agg_level0 | expr_muldiv | expr_modular;
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
