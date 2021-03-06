//line parse.y:7
package xp

import __yyfmt__ "fmt"

//line parse.y:7
import (
	"clive/dbg"
	"math"
	"os"
	"time"
)

var (
	debugYacc bool
	yprintf   = dbg.FlagPrintf(os.Stderr, &debugYacc)
)

//line parse.y:23
type yySymType struct {
	yys  int
	ival int64
	uval uint64
	fval float64
	sval string
	tval time.Time
	vval interface{}
}

const INT = 57346
const UINT = 57347
const NUM = 57348
const FUNC = 57349
const NAME = 57350
const TIME = 57351
const OR = 57352
const AND = 57353
const EQN = 57354
const NEQ = 57355
const LE = 57356
const GE = 57357
const SLEFT = 57358
const SRIGHT = 57359
const UMINUS = 57360

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT",
	"UINT",
	"NUM",
	"FUNC",
	"NAME",
	"TIME",
	"OR",
	"AND",
	"'='",
	"EQN",
	"NEQ",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'&'",
	"'|'",
	"SLEFT",
	"SRIGHT",
	"UMINUS",
	"'!'",
	"'^'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parse.y:173

var funcs = map[string]func(float64) float64{
	"abs":   math.Abs,
	"acos":  math.Acos,
	"acosh": math.Acosh,
	"asin":  math.Asin,
	"asinh": math.Asinh,
	"atan":  math.Atan,
	"atanh": math.Atanh,
	"cbrt":  math.Cbrt,
	"cos":   math.Cos,
	"cosh":  math.Cosh,
	"exp":   math.Exp,
	"exp2":  math.Exp2,
	"floor": math.Floor,
	"Γ":     math.Gamma,
	"log":   math.Log,
	"log10": math.Log10,
	"log2":  math.Log2,
	"sin":   math.Sin,
	"sinh":  math.Sinh,
	"sqrt":  math.Sqrt,
	"tan":   math.Tan,
	"tanh":  math.Tanh,
	"trunc": math.Trunc,
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 30
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 147

var yyAct = [...]int{

	2, 1, 0, 0, 31, 32, 33, 0, 0, 0,
	0, 0, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 28, 27, 24, 25, 26, 20, 21, 22,
	23, 13, 14, 15, 16, 17, 29, 30, 18, 19,
	0, 0, 0, 0, 54, 28, 27, 24, 25, 26,
	20, 21, 22, 23, 13, 14, 15, 16, 17, 29,
	30, 18, 19, 27, 24, 25, 26, 20, 21, 22,
	23, 13, 14, 15, 16, 17, 29, 30, 18, 19,
	24, 25, 26, 20, 21, 22, 23, 13, 14, 15,
	16, 17, 29, 30, 18, 19, 7, 8, 6, 5,
	9, 10, 13, 14, 15, 16, 17, 29, 30, 18,
	19, 0, 3, 0, 15, 16, 17, 29, 30, 18,
	19, 11, 12, 4, 20, 21, 22, 23, 13, 14,
	15, 16, 17, 29, 30, 18, 19,
}
var yyPact = [...]int{

	102, -1000, 45, 102, 102, 102, -1000, -1000, -1000, -1000,
	-1000, 102, 102, 102, 102, 102, 102, 102, 102, 102,
	102, 102, 102, 102, 102, 102, 102, 102, 102, 102,
	102, -1000, 22, -1000, -1000, -1000, 103, 103, -1000, -1000,
	-1000, -1000, -1000, 93, 93, 93, 93, 119, 119, 119,
	78, 62, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 1,
}
var yyR1 = [...]int{

	0, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}
var yyR2 = [...]int{

	0, 1, 3, 3, 3, 2, 3, 3, 3, 3,
	3, 2, 1, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
}
var yyChk = [...]int{

	-1000, -2, -1, 20, 31, 7, 6, 4, 5, 8,
	9, 29, 30, 19, 20, 21, 22, 23, 26, 27,
	15, 16, 17, 18, 12, 13, 14, 11, 10, 24,
	25, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, 32,
}
var yyDef = [...]int{

	0, -2, 1, 0, 0, 0, 12, 13, 14, 15,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 5, 0, 11, 28, 29, 2, 3, 4, 6,
	7, 8, 9, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 10,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 29, 3, 3, 3, 23, 24, 3,
	31, 32, 21, 19, 3, 20, 3, 22, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 12, 16, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 30, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 25,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	13, 14, 17, 18, 26, 27, 28,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

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
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yychar = -1
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
		//line parse.y:46
		{
			x := yylex.(*lex)
			x.result = yyDollar[1].vval
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:53
		{
			yyVAL.vval = add(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:57
		{
			yyVAL.vval = sub(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:61
		{
			yyVAL.vval = mul(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:65
		{
			yyVAL.vval = minus(yyDollar[2].vval)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:69
		{
			yyVAL.vval = div(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:73
		{
			yyVAL.vval = mod(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:77
		{
			yyVAL.vval = shiftleft(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:81
		{
			yyVAL.vval = shiftright(yyDollar[1].vval, yyDollar[3].vval)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:85
		{
			yyVAL.vval = yyDollar[2].vval
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:89
		{
			if f, ok := funcs[yyDollar[1].sval]; ok {
				n := Nval(yyDollar[2].vval)
				yyVAL.vval = f(n)
			} else if v, err := fmtf(yyDollar[1].sval, yyDollar[2].vval); err == nil {
				yyVAL.vval = v
			} else if v, err := attr(yyDollar[1].sval, yyDollar[2].vval); err == nil {
				yyVAL.vval = v
			} else {
				panic("unknown function")
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:102
		{
			yyVAL.vval = value(yyDollar[1].fval)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:106
		{
			yyVAL.vval = value(yyDollar[1].ival)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:110
		{
			yyVAL.vval = value(yyDollar[1].uval)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:114
		{
			yyVAL.vval = value(yyDollar[1].sval)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:118
		{
			yyVAL.vval = value(yyDollar[1].tval)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:122
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) < 0)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:126
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) > 0)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:130
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) <= 0)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:134
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) >= 0)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:138
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) == 0)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:142
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) == 0)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:146
		{
			yyVAL.vval = value(cmp(yyDollar[1].vval, yyDollar[3].vval) != 0)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:150
		{
			yyVAL.vval = value(Bval(yyDollar[1].vval) && Bval(yyDollar[3].vval))
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:154
		{
			yyVAL.vval = value(Bval(yyDollar[1].vval) || Bval(yyDollar[3].vval))
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:158
		{
			yyVAL.vval = value(Ival(yyDollar[1].vval) & Ival(yyDollar[3].vval))
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:162
		{
			yyVAL.vval = value(Ival(yyDollar[1].vval) | Ival(yyDollar[3].vval))
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:166
		{
			yyVAL.vval = value(!Bval(yyDollar[2].vval))
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:170
		{
			yyVAL.vval = value(^Ival(yyDollar[2].vval))
		}
	}
	goto yystack /* stack new state and value */
}
