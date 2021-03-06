/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package terminal

import (
	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
)

// Nil matches a nil literal
func Nil(nilStr string) *parser.NamedFunc {
	if nilStr == "" {
		panic("Nil() should not be called with an empty nil string")
	}

	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		tr := r.(*text.Reader)
		if readerPos, found := tr.MatchWord(pos, nilStr); found {
			return ast.NewTerminalNode("NIL", nil, pos, readerPos), nil, data.EmptyIntSet
		}

		return nil, nil, data.EmptyIntSet
	}).WithName(nilStr)
}
