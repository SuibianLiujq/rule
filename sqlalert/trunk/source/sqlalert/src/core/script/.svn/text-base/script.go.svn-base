package script

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
const Y_VALUE = 57348
const Y_STR = 57349
const Y_INT = 57350
const Y_FLOAT = 57351
const Y_BOOL = 57352
const Y_NULL = 57353
const Y_EQ = 57354
const Y_NE = 57355
const Y_LE = 57356
const Y_GE = 57357
const Y_AND = 57358
const Y_OR = 57359
const Y_NOT = 57360
const Y_INC = 57361
const Y_DEC = 57362
const Y_IF = 57363
const Y_ELSE = 57364
const Y_FOR = 57365
const Y_IN = 57366
const Y_CONTINUE = 57367
const Y_BREAK = 57368
const Y_RETURN = 57369
const Y_DEF = 57370
const Y_INCLUDE = 57371
const Y_IMPORT = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Y_ERR",
	"Y_IDENT",
	"Y_VALUE",
	"Y_STR",
	"Y_INT",
	"Y_FLOAT",
	"Y_BOOL",
	"Y_NULL",
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
	"Y_IF",
	"Y_ELSE",
	"Y_FOR",
	"Y_IN",
	"Y_CONTINUE",
	"Y_BREAK",
	"Y_RETURN",
	"Y_DEF",
	"Y_INCLUDE",
	"Y_IMPORT",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	15, 53,
	-2, 1,
	-1, 4,
	13, 183,
	-2, 4,
	-1, 6,
	20, 85,
	24, 77,
	25, 77,
	26, 77,
	27, 77,
	28, 77,
	29, 79,
	30, 79,
	31, 79,
	32, 79,
	33, 79,
	34, 79,
	35, 79,
	36, 79,
	43, 77,
	-2, 59,
	-1, 20,
	29, 80,
	30, 80,
	31, 80,
	32, 80,
	33, 80,
	34, 80,
	35, 80,
	36, 80,
	-2, 72,
	-1, 21,
	35, 173,
	36, 173,
	-2, 73,
	-1, 22,
	35, 174,
	36, 174,
	-2, 74,
	-1, 32,
	24, 78,
	25, 78,
	26, 78,
	27, 78,
	28, 78,
	43, 78,
	-2, 68,
	-1, 33,
	24, 148,
	25, 148,
	43, 163,
	-2, 69,
	-1, 34,
	20, 89,
	-2, 70,
	-1, 47,
	20, 88,
	-2, 64,
	-1, 48,
	38, 122,
	39, 122,
	-2, 65,
	-1, 50,
	12, 51,
	20, 87,
	38, 121,
	39, 121,
	-2, 67,
	-1, 52,
	26, 153,
	27, 153,
	-2, 145,
	-1, 53,
	28, 157,
	-2, 146,
	-1, 57,
	35, 172,
	36, 172,
	-2, 167,
	-1, 60,
	12, 50,
	20, 86,
	38, 120,
	39, 120,
	-2, 102,
	-1, 72,
	26, 152,
	27, 152,
	28, 156,
	43, 162,
	-2, 147,
	-1, 89,
	12, 50,
	14, 138,
	20, 86,
	38, 120,
	39, 120,
	-2, 102,
	-1, 90,
	14, 139,
	-2, 103,
	-1, 92,
	13, 59,
	15, 59,
	20, 85,
	24, 77,
	25, 77,
	26, 77,
	27, 77,
	28, 77,
	43, 77,
	-2, 79,
	-1, 94,
	20, 87,
	38, 121,
	39, 121,
	-2, 67,
	-1, 95,
	20, 86,
	38, 120,
	39, 120,
	-2, 102,
	-1, 104,
	22, 32,
	-2, 22,
	-1, 109,
	12, 50,
	20, 86,
	38, 120,
	39, 120,
	-2, 102,
	-1, 120,
	15, 57,
	-2, 183,
	-1, 129,
	29, 167,
	30, 167,
	31, 167,
	32, 167,
	33, 167,
	34, 167,
	-2, 178,
	-1, 132,
	20, 85,
	24, 77,
	25, 77,
	26, 77,
	27, 77,
	28, 77,
	43, 77,
	-2, 79,
	-1, 139,
	20, 85,
	-2, 77,
	-1, 152,
	13, 183,
	-2, 132,
	-1, 162,
	14, 184,
	-2, 183,
	-1, 181,
	20, 89,
	-2, 165,
	-1, 183,
	29, 167,
	30, 167,
	31, 167,
	32, 167,
	33, 167,
	34, 167,
	-2, 175,
	-1, 188,
	13, 183,
	-2, 116,
	-1, 190,
	21, 128,
	26, 152,
	27, 152,
	28, 156,
	-2, 147,
	-1, 191,
	21, 129,
	-2, 148,
	-1, 193,
	26, 152,
	27, 152,
	28, 156,
	-2, 149,
	-1, 194,
	26, 153,
	27, 153,
	-2, 150,
	-1, 195,
	28, 157,
	-2, 151,
	-1, 197,
	28, 156,
	-2, 154,
	-1, 198,
	28, 157,
	-2, 155,
	-1, 205,
	13, 183,
	-2, 140,
	-1, 225,
	20, 85,
	24, 77,
	25, 77,
	26, 77,
	27, 77,
	28, 77,
	29, 79,
	30, 79,
	31, 79,
	32, 79,
	33, 79,
	34, 79,
	35, 79,
	36, 79,
	43, 77,
	-2, 81,
	-1, 240,
	20, 127,
	-2, 30,
}

const yyPrivate = 57344

const yyLast = 477

var yyAct = [...]int{

	95, 60, 14, 47, 6, 208, 60, 24, 79, 187,
	89, 8, 25, 94, 50, 92, 34, 82, 61, 50,
	87, 151, 33, 50, 20, 99, 52, 57, 90, 38,
	21, 109, 111, 102, 4, 7, 106, 105, 110, 53,
	231, 83, 2, 104, 50, 93, 91, 32, 124, 166,
	101, 149, 72, 135, 136, 113, 114, 117, 118, 115,
	116, 73, 121, 122, 132, 126, 127, 165, 147, 148,
	249, 139, 144, 145, 10, 157, 71, 120, 84, 245,
	123, 247, 156, 141, 133, 153, 60, 129, 234, 131,
	130, 229, 228, 227, 178, 202, 167, 82, 221, 50,
	201, 233, 220, 154, 232, 219, 152, 142, 109, 212,
	211, 173, 92, 134, 140, 110, 160, 132, 159, 138,
	171, 50, 134, 174, 230, 172, 164, 97, 96, 139,
	132, 206, 158, 162, 163, 98, 119, 133, 74, 73,
	176, 181, 223, 91, 170, 169, 161, 139, 139, 23,
	133, 139, 15, 183, 139, 185, 184, 186, 153, 141,
	141, 177, 73, 141, 74, 191, 141, 60, 60, 188,
	194, 65, 140, 210, 215, 217, 154, 180, 203, 60,
	50, 50, 48, 195, 218, 214, 198, 189, 209, 5,
	140, 140, 50, 205, 140, 190, 193, 140, 85, 197,
	207, 67, 200, 179, 56, 35, 128, 182, 58, 226,
	60, 225, 37, 22, 175, 36, 137, 49, 199, 70,
	196, 82, 69, 50, 222, 192, 68, 51, 60, 86,
	235, 60, 242, 236, 246, 244, 241, 152, 239, 204,
	240, 50, 88, 238, 50, 243, 60, 55, 141, 150,
	54, 248, 64, 63, 76, 188, 75, 62, 73, 50,
	74, 80, 81, 77, 78, 46, 11, 224, 26, 39,
	28, 10, 27, 71, 216, 84, 155, 19, 66, 103,
	237, 213, 108, 18, 100, 17, 29, 13, 12, 16,
	59, 3, 1, 45, 112, 30, 125, 40, 41, 42,
	31, 43, 44, 73, 146, 74, 80, 81, 77, 78,
	143, 0, 0, 0, 0, 0, 10, 0, 71, 0,
	9, 0, 0, 66, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 45, 0,
	30, 0, 40, 41, 42, 31, 43, 44, 73, 0,
	74, 80, 81, 77, 78, 0, 0, 0, 0, 0,
	0, 10, 0, 71, 0, 84, 0, 0, 66, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 45, 0, 30, 0, 40, 41, 42,
	31, 43, 44, 73, 0, 74, 80, 81, 77, 78,
	0, 0, 0, 0, 0, 0, 107, 168, 71, 0,
	84, 0, 0, 66, 73, 0, 74, 80, 81, 77,
	78, 0, 0, 0, 0, 59, 0, 10, 0, 71,
	0, 84, 0, 0, 66, 73, 0, 74, 80, 81,
	77, 78, 0, 0, 0, 0, 59, 0, 107, 0,
	71, 0, 84, 0, 0, 66, 73, 0, 74, 80,
	81, 77, 78, 0, 0, 0, 0, 59, 0, 10,
	0, 71, 0, 84, 0, 0, 66,
}
var yyPact = [...]int{

	298, -1000, -1000, -1000, -1000, 343, -1000, -1000, -1000, 298,
	409, -1000, 113, -1000, -1000, -1000, 112, -1000, -1000, -1000,
	-1000, -1000, -1000, 122, -1000, -1000, -1000, -1000, -1000, 9,
	430, 134, -1000, -1000, -1000, -1000, 26, -1000, -1000, 124,
	-1000, -1000, 409, 131, 131, 409, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 5, -1000, 30, 409,
	104, -1000, -1000, -1000, -1000, 15, 451, 87, 48, 42,
	23, 409, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 157, 253, 59, -1000, 118, 104,
	-1000, 99, 97, -1000, -1000, 104, -1000, -1000, 409, -1000,
	9, 27, -1000, 74, -1000, -1000, -1000, 388, 110, 95,
	-1000, 105, 451, -1000, -1000, -1000, -1000, -1000, -1000, 409,
	-1000, -1000, -1000, 72, 451, 409, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 409, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 451, 451, -1000, -1000, 451, -1000, -1000, 451,
	79, -1000, -1000, -1000, -1000, -1000, -1000, 157, 409, -1000,
	-1000, 117, -1000, -1000, -1000, 409, 298, 298, -1000, 91,
	90, -1000, 409, 134, 134, -1000, -1000, -1000, 298, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 86, -1000, -1000, 77,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 409, -1000, -1000, -1000, 409, 71, 69, 343,
	68, -1000, -1000, 109, -1000, -3, 85, -1000, 65, -1000,
	409, -1000, -1000, -1000, -1000, -1000, -1000, 298, -1000, -1000,
	409, 56, 57, 134, -1000, -1000, 58, -1000, -1000, -1000,
	-1000, -1000, 104, -1000, -1000, 298, -1000, -1000, 47, -1000,
}
var yyPgo = [...]int{

	0, 310, 304, 296, 294, 292, 41, 291, 11, 289,
	288, 287, 188, 5, 286, 285, 33, 284, 25, 283,
	37, 282, 281, 280, 279, 36, 277, 274, 272, 270,
	7, 269, 12, 268, 35, 34, 4, 52, 27, 267,
	47, 24, 2, 266, 8, 265, 0, 18, 257, 256,
	254, 253, 252, 16, 250, 21, 249, 247, 20, 242,
	239, 229, 22, 227, 226, 225, 26, 222, 220, 39,
	219, 218, 217, 216, 30, 215, 214, 213, 212, 208,
	207, 29, 206, 205, 204, 203, 13, 201, 187, 182,
	171, 3, 9, 157, 152, 149, 146, 142,
}
var yyR1 = [...]int{

	0, 5, 5, 7, 6, 9, 9, 9, 9, 9,
	10, 10, 11, 11, 11, 8, 8, 12, 12, 12,
	13, 13, 21, 21, 22, 23, 23, 20, 20, 20,
	25, 25, 24, 24, 24, 19, 14, 16, 18, 18,
	17, 17, 15, 15, 26, 27, 27, 27, 28, 29,
	31, 31, 30, 32, 33, 33, 33, 33, 34, 35,
	35, 36, 36, 40, 40, 40, 40, 40, 41, 41,
	41, 41, 42, 42, 42, 43, 43, 37, 37, 38,
	38, 39, 39, 53, 53, 44, 44, 44, 44, 44,
	1, 1, 2, 2, 3, 3, 4, 4, 4, 4,
	4, 4, 45, 45, 45, 45, 45, 48, 48, 46,
	47, 49, 50, 51, 52, 91, 92, 93, 93, 93,
	90, 90, 90, 89, 89, 73, 72, 87, 88, 88,
	86, 54, 55, 56, 56, 56, 57, 58, 59, 59,
	60, 61, 61, 61, 62, 62, 62, 64, 64, 65,
	65, 65, 67, 67, 68, 68, 70, 70, 71, 63,
	66, 69, 84, 84, 85, 85, 83, 75, 76, 74,
	77, 77, 79, 79, 79, 80, 80, 80, 82, 82,
	82, 78, 81, 95, 96, 97, 94,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 1, 1, 1, 1, 1, 2, 3,
	0, 1, 1, 1, 1, 1, 1, 2, 3, 5,
	5, 3, 1, 1, 1, 5, 5, 6, 0, 4,
	1, 2, 2, 3, 8, 0, 1, 3, 2, 2,
	1, 1, 3, 1, 1, 1, 1, 2, 1, 1,
	1, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 4, 1, 0, 1, 3,
	1, 1, 1, 2, 2, 1, 2, 1, 1, 1,
	4, 3, 1, 0, 1, 3, 3, 3, 1, 1,
	1, 0, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	3, 3, 1, 1, 1, 1, 3, 1, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 1, 1, 5,
}
var yyChk = [...]int{

	-1000, -5, -6, -7, -35, -12, -36, -34, -8, 22,
	18, -43, -10, -11, -42, -94, -9, -15, -19, -26,
	-41, -74, -77, -95, -30, -32, -33, -28, -29, -14,
	42, 47, -40, -62, -53, -83, -75, -78, -81, -31,
	44, 45, 46, 48, 49, 40, -45, -91, -89, -72,
	-86, -63, -66, -69, -54, -57, -84, -38, -79, 37,
	-46, -47, -48, -51, -52, -90, 25, -87, -64, -67,
	-70, 20, -37, 5, 7, -49, -50, 10, 11, -44,
	8, 9, -8, -6, 22, -12, -61, -58, -59, -46,
	-47, -34, -36, -35, -86, -46, 15, 15, 13, -18,
	-17, 41, -16, -24, -32, -20, -25, 18, -21, -46,
	-30, -46, -4, 29, 30, 33, 34, 31, 32, 12,
	-35, -47, -47, -35, 43, -3, 35, 36, -82, -38,
	-74, -81, -36, -41, 18, 38, 39, -73, -37, -36,
	-40, -53, 20, -1, 24, 25, -2, 26, 27, 28,
	-56, -55, -35, -46, -47, 23, 23, 16, 14, 19,
	19, -96, -35, -16, -18, 40, 22, 22, 19, -20,
	-25, -32, 15, 16, 18, -76, -38, -6, 22, -85,
	-37, -53, -80, -38, -74, -81, -93, -92, -35, -88,
	-37, -62, -65, -37, -66, -69, -68, -37, -69, -71,
	-37, 21, 16, -58, -60, -35, 14, -35, -13, -12,
	-13, 19, 19, -22, -32, -46, -27, -46, -13, 19,
	16, 21, -55, -97, -39, -36, -42, 22, 23, 23,
	15, 43, 19, 16, 23, -92, -13, -23, -32, -30,
	-44, -36, -46, -86, -91, 22, -46, 23, -13, 23,
}
var yyDef = [...]int{

	0, -2, -2, 2, -2, 3, -2, 60, 17, 141,
	0, 58, 15, 16, 75, 76, 0, 12, 13, 14,
	-2, -2, -2, 0, 5, 6, 7, 8, 9, 38,
	0, 0, -2, -2, -2, 71, 0, 170, 171, 0,
	54, 55, 56, 0, 0, 0, 63, -2, -2, 66,
	-2, 144, -2, -2, 83, 84, 0, -2, 0, 0,
	-2, 103, 104, 105, 106, 0, 0, 0, 0, 0,
	0, 133, -2, 109, 110, 107, 108, 113, 114, 127,
	111, 112, 18, 53, 141, 0, 0, 142, 0, -2,
	-2, 60, -2, 183, -2, -2, 11, 10, 0, 42,
	38, 0, 40, 0, -2, 33, 34, 0, 0, -2,
	23, 0, 0, 96, 97, 98, 99, 100, 101, 0,
	-2, 48, 49, 183, 0, 0, 94, 95, 182, -2,
	179, 180, -2, 80, 117, 123, 124, 126, 125, -2,
	78, 89, 0, 0, 90, 91, 0, 92, 93, 0,
	0, 134, -2, 138, 139, 19, 136, 0, 0, 61,
	62, 0, -2, 41, 43, 0, 20, 20, 27, 0,
	0, 22, 0, 0, 45, 169, 168, 52, 20, 166,
	164, -2, 181, -2, 176, 177, 0, 118, -2, 0,
	-2, -2, 159, -2, -2, -2, 160, -2, -2, 161,
	158, 131, 0, 143, 137, -2, 0, 183, 0, 21,
	0, 28, 31, 0, 24, 0, 0, 46, 0, 115,
	0, 130, 135, 186, 185, -2, 82, 20, 39, 35,
	0, 0, 0, 0, 36, 119, 0, 29, 25, 26,
	-2, 85, 86, 87, 88, 20, 47, 37, 0, 44,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 17, 28, 3, 3,
	18, 19, 26, 24, 16, 25, 3, 27, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 14, 15,
	29, 12, 30, 13, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 20, 3, 21, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 23,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			setToken(yylex, yyDollar[1].token)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenStmts{}).Init(yyDollar[1].list)
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = yyDollar[1].token
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = yyDollar[1].token
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].token)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = yyDollar[2].list
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = yyDollar[1].list
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenForIter{}).Init(nil, nil, nil)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenForIter{}).Init(yyDollar[1].token, yyDollar[3].token, yyDollar[5].token)
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenForinIter{}).Init(yyDollar[1].token, yyDollar[3].token, yyDollar[5].token)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if yyDollar[1].token.(*TokenExpr).Token.Type() == T_IN {
				t := yyDollar[1].token.(*TokenExpr).Token.(*TokenIn)
				yyVAL.token = (&TokenForinIter{}).Init(nil, t.Key, t.Object)
			} else {
				yyVAL.token = (&TokenForIter{}).Init(nil, yyDollar[1].token.(*TokenExpr).Token, nil)
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = yyDollar[1].token
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = yyDollar[1].token
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyDollar[2].token.Type() == T_FOR_ITER {
				yyVAL.token = (&TokenFor{}).Init(yyDollar[2].token, yyDollar[4].list)
			} else {
				yyVAL.token = (&TokenForin{}).Init(yyDollar[2].token, yyDollar[4].list)
			}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenIf{}).Init(yyDollar[2].token, yyDollar[4].list)
		}
	case 37:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.token = (&TokenElseIf{}).Init(yyDollar[3].token, yyDollar[5].list)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.token = nil
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenElse{}).Init(yyDollar[3].list)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].token)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenIfElse{}).Init(yyDollar[1].token, nil, yyDollar[2].token)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenIfElse{}).Init(yyDollar[1].token, yyDollar[2].list, yyDollar[3].token)
		}
	case 44:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.token = (&TokenDefine{}).Init(yyDollar[2].token, yyDollar[4].list, yyDollar[7].list)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenInclude{}).Init(yyDollar[2].token)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenImport{}).Init(yyDollar[2].token)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenAssign{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenExpr{}).Init(yyDollar[1].token)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenContinue{}).Init()
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenBreak{}).Init()
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenReturn{}).Init(nil)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenReturn{}).Init(yyDollar[2].token)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = yyDollar[2].token
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenIdent{}).Init(yyDollar[1].value, getSrc(yylex))
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenStr{}).Init(yyDollar[1].value, getSrc(yylex))
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenInt{}).Init(yyDollar[1].value, getSrc(yylex))
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenFloat{}).Init(yyDollar[1].value, getSrc(yylex))
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenBool{}).Init(yyDollar[1].value, getSrc(yylex))
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.token = (&TokenNull{}).Init(getSrc(yylex))
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenFunc{}).Init(yyDollar[1].token, yyDollar[3].list)
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, nil)
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, nil)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(nil, yyDollar[1].operator, yyDollar[2].token)
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.token = (&TokenIndex{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenList{}).Init(yyDollar[2].list)
		}
	case 133:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenDict{}).Init(yyDollar[2].list)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenPair{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 141:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.list = []Token{}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.list = []Token{yyDollar[1].token}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].token)
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenOper{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
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
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenIn{}).Init(yyDollar[1].token, yyDollar[3].token)
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenComp{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.token = (&TokenLogical{}).Init(yyDollar[1].token, yyDollar[2].operator, yyDollar[3].token)
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.token = (&TokenLogical{}).Init(nil, yyDollar[1].operator, yyDollar[2].token)
		}
	case 186:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.token = (&TokenCond{}).Init(yyDollar[1].token, yyDollar[3].token, yyDollar[5].token)
		}
	}
	goto yystack /* stack new state and value */
}
