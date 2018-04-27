%{
package json
type __pair struct {
	key   string
	value interface{}
}
%}

%union {
	empty interface{}
	value interface{}
	list  []interface{}
	pair  *__pair
}

%token<empty> Y_ERR '-' '[' ']' ':' ',' '{' '}'
%token<value> Y_STR Y_INT Y_FLOAT Y_BOOL Y_NULL

%type<value> object
%type<value> base
%type<value> negative
%type<value> list
%type<list>  list_list
%type<value> dict
%type<pair>  pair
%type<list>  pair_list

%%
all:
	  dict { setValue(yylex, $1) }
	| list { setValue(yylex, $1) }
;

object: base | dict | list;

pair: Y_STR ':' object { $$ = &__pair{key: $1.(string), value: $3} };
pair_list:
	                     { $$ = []interface{}{} }
	| pair               { $$ = []interface{}{$1} }
	| pair_list ',' pair { $$ = append($1, $3) }
;

dict: '{' pair_list '}' {
	dict := map[string]interface{}{}
	for _, item := range $2 {
		dict[item.(*__pair).key] = item.(*__pair).value
	}
	$$ = dict
};

list: '[' list_list ']' { $$ = $2 };
list_list:
	                       { $$ = []interface{}{} }
	| object               { $$ = []interface{}{$1} }
	| list_list ',' object { $$ = append($1, $3) }
;

base: Y_STR | Y_BOOL | Y_NULL | Y_INT | Y_FLOAT | negative;
negative:
	  '-' Y_INT   { $$ = -$2.(int64) }
	| '-' Y_FLOAT { $$ = -$2.(float64) }
;

%%
