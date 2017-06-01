package terminal

import (
	"fmt"
	"strconv"

	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/reader"
	"github.com/opsidian/parsley/text"
)

// Integer matches all integer numbers and zero with an optional -/+ sign
func Integer() parser.Func {
	return parser.Func(func(ctx data.IntMap, r reader.Reader) (data.IntSet, parser.ResultSet) {
		tr := r.(*text.Reader)
		if matches, pos := tr.ReadMatch("^[-+]?0|[1-9][0-9]*"); matches != nil {
			intValue, err := strconv.Atoi(matches[0])
			if err != nil {
				panic(fmt.Sprintf("Could not convert %s to integer", matches[0]))
			}
			return parser.NoCurtailingParsers(), parser.NewResult(ast.NewTerminalNode("INT", pos, intValue), r).AsSet()
		}
		return parser.NoCurtailingParsers(), nil
	})
}