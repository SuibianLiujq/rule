package sql

import __yyfmt__ "fmt"

type yySymType struct {
	yys      int
	empty    interface{}
	value    interface{}
	token    Token
	list     []Token
	operator TokenType
}

const Y_ERR = 57346
const Y_IDENT = 57347
const Y_STR = 57348
const Y_INT = 57349
const Y_FLOAT = 57350
const Y_BOOL = 57351
const Y_NULL = 57352
const Y_SELECT = 57353
const Y_FROM = 57354
const Y_WHERE = 57355
const Y_GROUP = 57356
const Y_ORDER = 57357
const Y_LIMIT = 57358
const Y_HAVING = 57359
const Y_AS = 57360
const Y_IN = 57361
const Y_BY = 57362
const Y_LIKE = 57363
const Y_UNLIKE = 57364
const Y_DESC = 57365
const Y_ASC = 57366
const Y_INTERVAL = 57367
const Y_UNIQUE = 57368
const Y_EQ = 57369
const Y_NE = 57370
const Y_LE = 57371
const Y_GE = 57372
const Y_AND = 57373
const Y_OR = 57374
const Y_NOT = 57375
const Y_INC = 57376
const Y_DEC = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Y_ERR",
	"Y_IDENT",
	"Y_STR",
	"Y_INT",
	"Y_FLOAT",
	"Y_BOOL",
	"Y_NULL",
	"Y_SELECT",
	"Y_FROM",
	"Y_WHERE",
	"Y_GROUP",
	"Y_ORDER",
	"Y_LIMIT",
	"Y_HAVING",
	"Y_AS",
	"Y_IN",
	"Y_BY",
	"Y_LIKE",
	"Y_UNLIKE",
	"Y_DESC",
	"Y_ASC",
	"Y_INTERVAL",
	"Y_UNIQUE",
	"'='",
	"'?'",
	"':'",
	"';'",
	"','",
	"'$'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'<'",
	"'>'",
	"Y_EQ",
	"Y_NE",
	"Y_LE",
	"Y_GE",
	"Y_AND",
	"Y_OR",
	"Y_NOT",
	"Y_INC",
	"Y_DEC",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 22,
	28, 184,
	-2, 12,
	-1, 24,
	19, 69,
	35, 77,
	39, 69,
	40, 69,
	41, 69,
	42, 69,
	43, 69,
	44, 71,
	45, 71,
	46, 71,
	47, 71,
	48, 71,
	49, 71,
	50, 71,
	51, 71,
	-2, 51,
	-1, 30,
	44, 72,
	45, 72,
	46, 72,
	47, 72,
	48, 72,
	49, 72,
	50, 72,
	51, 72,
	-2, 64,
	-1, 31,
	50, 174,
	51, 174,
	-2, 65,
	-1, 32,
	50, 175,
	51, 175,
	-2, 66,
	-1, 34,
	19, 70,
	39, 70,
	40, 70,
	41, 70,
	42, 70,
	43, 70,
	-2, 60,
	-1, 35,
	19, 164,
	39, 149,
	40, 149,
	-2, 61,
	-1, 36,
	35, 81,
	-2, 62,
	-1, 42,
	35, 80,
	-2, 56,
	-1, 43,
	53, 123,
	54, 123,
	-2, 57,
	-1, 45,
	35, 79,
	53, 122,
	54, 122,
	-2, 59,
	-1, 47,
	41, 154,
	42, 154,
	-2, 146,
	-1, 48,
	43, 158,
	-2, 147,
	-1, 52,
	50, 173,
	51, 173,
	-2, 168,
	-1, 55,
	35, 78,
	53, 121,
	54, 121,
	-2, 94,
	-1, 71,
	19, 163,
	41, 153,
	42, 153,
	43, 157,
	-2, 148,
	-1, 91,
	28, 184,
	-2, 23,
	-1, 98,
	28, 184,
	-2, 45,
	-1, 102,
	19, 69,
	28, 51,
	35, 77,
	39, 69,
	40, 69,
	41, 69,
	42, 69,
	43, 69,
	-2, 71,
	-1, 117,
	44, 168,
	45, 168,
	46, 168,
	47, 168,
	48, 168,
	49, 168,
	-2, 179,
	-1, 120,
	19, 69,
	35, 77,
	39, 69,
	40, 69,
	41, 69,
	42, 69,
	43, 69,
	-2, 71,
	-1, 128,
	35, 77,
	-2, 69,
	-1, 141,
	28, 184,
	-2, 133,
	-1, 152,
	1, 26,
	12, 26,
	13, 26,
	14, 26,
	15, 26,
	16, 26,
	17, 26,
	31, 26,
	35, 78,
	53, 121,
	54, 121,
	-2, 94,
	-1, 153,
	1, 27,
	12, 27,
	13, 27,
	14, 27,
	15, 27,
	16, 27,
	17, 27,
	31, 27,
	-2, 95,
	-1, 169,
	29, 185,
	-2, 184,
	-1, 174,
	35, 81,
	-2, 166,
	-1, 176,
	44, 168,
	45, 168,
	46, 168,
	47, 168,
	48, 168,
	49, 168,
	-2, 176,
	-1, 181,
	28, 184,
	-2, 115,
	-1, 186,
	36, 129,
	41, 153,
	42, 153,
	43, 157,
	-2, 148,
	-1, 187,
	36, 130,
	-2, 149,
	-1, 189,
	41, 153,
	42, 153,
	43, 157,
	-2, 150,
	-1, 190,
	41, 154,
	42, 154,
	-2, 151,
	-1, 191,
	43, 158,
	-2, 152,
	-1, 193,
	43, 157,
	-2, 155,
	-1, 194,
	43, 158,
	-2, 156,
	-1, 214,
	28, 184,
	-2, 49,
	-1, 219,
	28, 184,
	-2, 141,
	-1, 226,
	19, 69,
	35, 77,
	39, 69,
	40, 69,
	41, 69,
	42, 69,
	43, 69,
	44, 71,
	45, 71,
	46, 71,
	47, 71,
	48, 71,
	49, 71,
	50, 71,
	51, 71,
	-2, 73,
}

const yyPrivate = 57344

const yyLast = 319

var yyAct = [...]int{

	155, 24, 55, 180, 154, 22, 151, 28, 158, 23,
	143, 140, 57, 95, 62, 56, 91, 86, 85, 35,
	98, 21, 36, 48, 47, 30, 150, 103, 102, 88,
	87, 96, 124, 125, 34, 52, 114, 115, 138, 40,
	106, 107, 110, 111, 108, 109, 215, 71, 131, 72,
	73, 81, 82, 76, 77, 31, 120, 136, 137, 229,
	123, 133, 134, 198, 200, 220, 128, 122, 197, 213,
	141, 199, 212, 145, 72, 73, 79, 26, 203, 69,
	121, 70, 167, 166, 64, 78, 146, 130, 148, 147,
	117, 206, 205, 152, 119, 152, 160, 162, 149, 129,
	22, 79, 99, 164, 23, 169, 153, 120, 153, 161,
	118, 211, 127, 201, 128, 120, 165, 104, 25, 156,
	93, 163, 92, 181, 208, 209, 112, 183, 100, 4,
	6, 121, 72, 128, 128, 174, 83, 128, 224, 121,
	128, 171, 81, 82, 168, 101, 33, 129, 29, 176,
	202, 187, 86, 178, 130, 130, 191, 190, 130, 194,
	173, 130, 72, 73, 88, 87, 129, 129, 204, 177,
	129, 179, 42, 129, 63, 96, 210, 43, 185, 186,
	189, 65, 45, 193, 172, 214, 196, 14, 15, 16,
	17, 18, 19, 51, 37, 116, 175, 53, 39, 141,
	32, 170, 219, 145, 38, 126, 221, 44, 152, 160,
	216, 217, 222, 226, 181, 223, 146, 228, 183, 227,
	195, 153, 161, 72, 73, 81, 82, 76, 77, 68,
	192, 67, 188, 72, 73, 81, 82, 76, 77, 66,
	46, 142, 218, 144, 184, 50, 139, 49, 90, 61,
	79, 26, 60, 69, 59, 70, 58, 75, 64, 78,
	79, 26, 74, 69, 41, 70, 80, 27, 64, 78,
	54, 72, 73, 81, 82, 76, 77, 225, 182, 97,
	54, 13, 94, 12, 157, 11, 207, 159, 10, 89,
	9, 8, 84, 7, 20, 3, 5, 2, 79, 26,
	1, 69, 105, 70, 113, 135, 64, 78, 132, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 54,
}
var yyPact = [...]int{

	118, -1000, -1000, 175, 266, 175, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 69, 228, 102, 100, 135, 266,
	71, -1000, 110, -1000, -1000, -1000, 266, -1000, -1000, -1000,
	-1000, -1000, -1000, 89, -1000, -1000, -1000, -1000, -4, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 107, -1000, -14, 266, 34, -1000, 127, -1000, -1000,
	-1000, -1000, -1000, -21, 44, 13, 22, 16, -5, 266,
	157, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 56,
	-1000, -1000, -1000, -1000, 67, -1000, -1000, -1000, -1000, -1000,
	266, -1000, 266, 157, 66, -1000, -1000, -1000, -1000, 266,
	157, 49, 48, -1000, 266, 44, -1000, -1000, -1000, -1000,
	-1000, -1000, 44, 266, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 218, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 44, 44, -1000, -1000, 44, -1000, -1000, 44, 32,
	-1000, -1000, 33, -1000, 84, -1000, -1000, 127, 45, 69,
	61, -1000, 34, -1000, -1000, 110, 61, 60, -1000, 101,
	-1000, -1000, 135, -1000, -1000, -1000, -1000, -1000, 82, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 38,
	-1000, 110, -1000, -1000, 266, 10, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 266, -1000,
	157, 266, 31, 127, -1000, 266, 157, -1000, -1000, -1000,
	-1000, 266, -1000, 218, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 25, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 308, 305, 304, 302, 300, 297, 296, 130, 295,
	21, 294, 293, 18, 292, 291, 290, 289, 288, 6,
	26, 8, 287, 286, 285, 284, 283, 13, 282, 281,
	279, 4, 278, 118, 0, 1, 47, 35, 277, 34,
	25, 7, 267, 266, 264, 14, 2, 15, 12, 262,
	257, 256, 254, 252, 249, 22, 247, 11, 246, 245,
	10, 243, 242, 241, 19, 240, 239, 232, 24, 231,
	230, 23, 229, 220, 207, 205, 55, 204, 201, 200,
	198, 197, 196, 39, 195, 194, 193, 184, 182, 181,
	178, 177, 174, 172, 3, 171, 148, 146, 144, 138,
}
var yyR1 = [...]int{

	0, 5, 6, 8, 8, 8, 8, 8, 8, 8,
	7, 7, 10, 10, 11, 11, 9, 13, 13, 13,
	14, 14, 12, 17, 15, 16, 19, 19, 19, 20,
	20, 18, 22, 22, 23, 23, 23, 21, 25, 25,
	24, 27, 28, 28, 26, 30, 29, 31, 31, 32,
	33, 34, 34, 35, 35, 39, 39, 39, 39, 39,
	40, 40, 40, 40, 41, 41, 41, 42, 42, 36,
	36, 37, 37, 38, 38, 55, 55, 43, 43, 43,
	43, 43, 1, 1, 2, 2, 3, 3, 4, 4,
	4, 4, 4, 4, 44, 44, 44, 44, 44, 44,
	44, 44, 48, 48, 46, 47, 49, 50, 51, 52,
	53, 54, 45, 45, 93, 94, 94, 94, 95, 95,
	95, 92, 92, 92, 91, 91, 75, 74, 89, 90,
	90, 88, 56, 57, 58, 58, 58, 59, 60, 61,
	61, 62, 63, 63, 63, 64, 64, 64, 66, 66,
	67, 67, 67, 69, 69, 70, 70, 72, 72, 73,
	65, 68, 71, 86, 86, 87, 87, 85, 77, 78,
	76, 79, 79, 81, 81, 81, 82, 82, 82, 84,
	84, 84, 80, 83, 97, 98, 99, 96,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 1, 3, 2, 1, 1, 1,
	1, 3, 2, 1, 2, 3, 1, 1, 1, 1,
	3, 3, 1, 1, 0, 1, 1, 2, 1, 3,
	3, 1, 1, 3, 2, 1, 2, 3, 3, 2,
	1, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 4, 5, 4, 1, 1, 1, 0, 1,
	3, 1, 1, 1, 2, 2, 1, 2, 1, 1,
	1, 4, 3, 1, 0, 1, 3, 3, 3, 1,
	1, 1, 0, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 3, 1, 1, 1, 1, 3, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 2, 1, 1, 1, 5,
}
var yyChk = [...]int{

	-1000, -5, -6, -9, 11, -7, -8, -12, -15, -16,
	-18, -24, -26, -29, 12, 13, 14, 15, 16, 17,
	-11, -10, -34, -31, -35, -33, 33, -42, -41, -96,
	-40, -76, -79, -97, -39, -64, -55, -85, -77, -80,
	-83, -44, -93, -91, -74, -88, -65, -68, -71, -56,
	-59, -86, -37, -81, 52, -46, -47, -48, -51, -52,
	-53, -54, -45, -92, 40, -89, -66, -69, -72, 35,
	37, -36, 5, 6, -49, -50, 9, 10, 41, 32,
	-43, 7, 8, -8, -14, -13, -46, -47, -45, -17,
	20, -34, 20, 20, -28, -27, -48, -30, -34, 31,
	18, -33, -35, -34, 28, -4, 44, 45, 48, 49,
	46, 47, 19, -3, 50, 51, -84, -37, -76, -83,
	-35, -40, 33, -46, 53, 54, -75, -36, -35, -39,
	-55, 35, -1, 39, 40, -2, 41, 42, 43, -58,
	-57, -34, -63, -60, -61, -46, -47, 33, 32, 31,
	-20, -19, -46, -47, -31, -34, -20, -25, -21, -22,
	-46, -47, 31, -10, -46, -47, 34, 34, -98, -34,
	-78, -37, -87, -36, -55, -82, -37, -76, -83, -95,
	-94, -34, -32, -31, 26, -90, -36, -64, -67, -36,
	-68, -71, -70, -36, -71, -73, -36, 36, 31, 38,
	31, 29, -46, 33, -13, 31, 31, -23, 23, 24,
	-27, 29, 34, 31, -34, 36, -57, -60, -62, -34,
	34, -46, -19, -21, -99, -38, -35, -41, -94, 34,
}
var yyDef = [...]int{

	0, -2, 1, 0, 0, 2, 10, 3, 4, 5,
	6, 7, 8, 9, 0, 0, 0, 0, 0, 0,
	16, 14, -2, 13, -2, 52, 0, 50, 67, 68,
	-2, -2, -2, 0, -2, -2, -2, 63, 0, 171,
	172, 55, -2, -2, 58, -2, 145, -2, -2, 75,
	76, 0, -2, 0, 0, -2, 95, 96, 97, 98,
	99, 100, 101, 0, 0, 0, 0, 0, 0, 134,
	142, -2, 104, 105, 102, 103, 108, 109, 110, 0,
	128, 106, 107, 11, 22, 20, 17, 18, 19, 24,
	0, -2, 0, 0, 44, 42, 41, 46, -2, 0,
	0, 52, -2, 184, 0, 0, 88, 89, 90, 91,
	92, 93, 0, 0, 86, 87, 183, -2, 180, 181,
	-2, 72, 118, 111, 124, 125, 127, 126, -2, 70,
	81, 0, 0, 82, 83, 0, 84, 85, 0, 0,
	135, -2, 0, 143, 0, 139, 140, 0, 0, 0,
	25, 29, -2, -2, 28, 184, 31, 40, 38, 34,
	32, 33, 0, 15, 47, 48, 53, 54, 0, -2,
	170, 169, 167, 165, -2, 182, -2, 177, 178, 0,
	119, -2, 116, 117, 0, 0, -2, -2, 160, -2,
	-2, -2, 161, -2, -2, 162, 159, 132, 0, 137,
	0, 0, 0, 0, 21, 0, 0, 37, 35, 36,
	43, 0, 114, 0, -2, 131, 136, 144, 138, -2,
	112, 0, 30, 39, 187, 186, -2, 74, 120, 113,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 32, 43, 3, 3,
	33, 34, 41, 39, 31, 40, 3, 42, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 30,
	44, 27, 45, 28, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 35, 3, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 3, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 46, 47, 48, 49, 50,
	51, 52, 53, 54,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			setToken(yylex, yyDollar[1].token)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenStmts{}).Init(yyDollar[1].token, yyDollar[2].list)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].token)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenSelect{}).Init(yyDollar[2].list)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenFrom{}).Init(yyDollar[2].list)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenWhere{}).Init(yyDollar[2].token)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenWhereBy{}).Init(yyDollar[3].list)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenGroupBy{}).Init(yyDollar[3].list)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.operator = T_ILL
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.operator = T_DESC
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.operator = T_ASC
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOrder{}).Init(yyDollar[1].token, yyDollar[2].operator)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenOrderBy{}).Init(yyDollar[3].list)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenLimit{}).Init(yyDollar[2].list)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenHaving{}).Init(yyDollar[2].token)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenAs{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenAs{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenUnique{}).Init(yyDollar[2].token)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenIdent{}).Init(yyDollar[1].value)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenStr{}).Init(yyDollar[1].value)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenInt{}).Init(yyDollar[1].value)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenFloat{}).Init(yyDollar[1].value)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenBool{}).Init(yyDollar[1].value)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenNull{}).Init()
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenStar{}).Init()
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenNumUnit{}).Init(yyDollar[1].token, yyDollar[2].token)
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenVar{}).Init(yyDollar[3].token, false)
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenVar{}).Init(yyDollar[4].token, true)
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenFunc{}).Init(yyDollar[1].token, yyDollar[3].list)
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, nil)
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, nil)
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(nil, yyDollar[1].operator, yyDollar[2].token)
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenIndex{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenList{}).Init(yyDollar[2].list)
		}
	case 134:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenDict{}).Init(yyDollar[2].list)
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenPair{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 142:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenIn{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenComp{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenLogical{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenLogical{}).Init(nil, yyDollar[1].operator, yyDollar[2].token)
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenCond{}).Init(yyDollar[1].token, yyDollar[3].token, yyDollar[5].token)
		}
	}
	goto yystack /* stack new state and value */
}
