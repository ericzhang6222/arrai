// AUTOGENERATED. DO NOT EDIT.
package syntax

import (
	"strings"

	"github.com/arr-ai/wbnf/wbnf"
)

func unfakeBackquote(s string) string {
	return strings.ReplaceAll(s, "‵", "`")
}

var arraiParsers = wbnf.MustCompile(unfakeBackquote(`
expr   -> C* amp="&"* @ C* arrow=(
              nest |
              unnest |
              ARROW @ |
              binding="->" C* "\\" C* IDENT C* %%bind C* @ |
              binding="->" C* %%bind @
          )* C*
        > C* @:binop=("without" | "with") C*
        > C* @:binop="||" C*
        > C* @:binop="&&" C*
        > C* @:compare=/{!?(?:<:|=|<=?|>=?|\((?:<=?|>=?|<>=?)\))} C*
        > C* @ if=("if" t=expr ("else" f=expr)?)* C*
        > C* @ cond=("cond" "{" (condition=(("(" expr:",", ")") | "("expr")" | pattern) ":" value=expr):",",? "}")? C*
        > C* @:binop=/{\+\+|[+|]|-%?} C*
        > C* @:binop=/{&~|&|~~?|[-<][-&][->]} C*
        > C* @:binop=/{//|[*/%]|\\} C*
        > C* @:rbinop="^" C*
        > C* unop=/{:>|=>|>>|[-+!*^]}* @ C*
        > C* @:binop=">>>" C*
        > C* @ postfix=/{count|single}? C* touch? C*
        > C* (get | @) tail_op=(safe_tail | tail)* C*
        > %!patternterms(expr)
        | C* cond=("cond" "{" (key="(" @ ")" ":" value=@):",",? ("_" ":" f=expr ","?)? "}") C*
        | C* "{:" C* embed=(grammar=@ ":" subgrammar=%%ast) ":}" C*
        | C* op="\\\\" @ C*
        | C* fn="\\" IDENT @ C*
        | C* "//" pkg=( "{" dot="."? PKGPATH "}" | std=IDENT?)
        | C* "(" @ ")" C*
        | C* let=("let" C* pattern C* "=" C* @ %%bind C* ";" C* @) C*
        | C* xstr C*
        | C* IDENT C*
        | C* STR C*
        | C* NUM C*
        | C* CHAR C*;
nest   -> C* "nest" names IDENT C*;
unnest -> C* "unnest" IDENT C*;
touch  -> C* ("->*" ("&"? IDENT | STR))+ "(" expr:"," ","? ")" C*;
get    -> C* dot="." ("&"? IDENT | STR | "*") C*;
names  -> C* "|" C* IDENT:"," C* "|" C*;
name   -> C* IDENT C* | C* STR C*;
xstr   -> C* quote=/{\$"\s*} part=( sexpr | fragment=/{(?: \\. | \$[^{"] | [^\\"$] )+} )* '"' C*
        | C* quote=/{\$'\s*} part=( sexpr | fragment=/{(?: \\. | \$[^{'] | [^\\'$] )+} )* "'" C*
        | C* quote=/{\$‵\s*} part=( sexpr | fragment=/{(?: ‵‵  | \$[^{‵] | [^‵  $] )+} )* "‵" C*;
sexpr  -> "${"
          C* expr C*
          control=/{ (?: : [-+#*\.\_0-9a-z]* (?: : (?: \\. | [^\\:}] )* ){0,2} )? }
          close=/{\}\s*};
tail   -> get
          | call=("("
                arg=(
                    expr (":" end=expr? (":" step=expr)?)?
                    |     ":" end=expr  (":" step=expr)?
                ):",",
            ")");
safe_tail -> first_safe=(tail "?") ops=(safe=(tail "?") | tail)* ":" fall=expr;
pattern -> extra | %!patternterms(pattern|expr) | IDENT | NUM;
extra -> ("..." ident=IDENT?);

ARROW  -> /{:>|=>|>>|orderby|order|where|sum|max|mean|median|min};
IDENT  -> /{ \. | [$@A-Za-z_][0-9$@A-Za-z_]* };
PKGPATH -> /{ (?: \\ | [^\\}] )* };
STR    -> /{ " (?: \\. | [^\\"] )* "
           | ' (?: \\. | [^\\'] )* '
           | ‵ (?: ‵‵  | [^‵  ] )* ‵
           };
NUM    -> /{ (?: \d+(?:\.\d*)? | \.\d+ ) (?: [Ee][-+]?\d+ )? };
CHAR   -> /{%(\\.|.)};
C      -> /{ # .* $ };

.wrapRE -> /{\s*()\s*};

.macro patternterms(top) {
    C* "{" C* rel=(names tuple=("(" v=top:",", ")"):",",?) "}" C*
  | C* "{" C* set=(elt=top:",",?) "}" C*
  | C* "{" C* dict=((ext=extra|key=expr ":" value=top):",",?) "}" C*
  | C* "[" C* array=(item=top:",",?) C* "]" C*
  | C* "<<" C* bytes=(item=top:",",?) C* ">>" C*
  | C* "(" tuple=(pairs=(extra|name? ":" v=top):",",?) ")" C*
  | C* "(" identpattern=IDENT ")" C*
};

`), nil)
