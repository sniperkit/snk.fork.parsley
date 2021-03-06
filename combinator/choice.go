/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package combinator

import (
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
)

// Choice tries to apply the given parsers until one of them succeeds
func Choice(name string, parsers ...parsley.Parser) *parser.NamedFunc {
	if parsers == nil {
		panic("No parsers were given")
	}

	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		cp := data.EmptyIntSet
		var err parsley.Error
		for _, p := range parsers {
			h.RegisterCall()
			node, err2, cp2 := p.Parse(h, leftRecCtx, r, pos)
			cp = cp.Union(cp2)
			if err2 != nil && (err == nil || err2.Pos() >= err.Pos()) {
				err = err2
			}
			if node != nil {
				return node, err, cp
			}
		}
		return nil, err, cp
	}).WithName(name)
}
