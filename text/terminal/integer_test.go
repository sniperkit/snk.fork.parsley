/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package terminal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
	"github.com/sniperkit/snk.fork.parsley/text/terminal"
)

var _ = Describe("Integer", func() {

	var p = terminal.Integer()

	It("should have a name", func() {
		Expect(p.Name()).ToNot(BeEmpty())
	})

	DescribeTable("should match",
		func(input string, startPos int, value interface{}, nodePos parsley.Pos, endPos int) {
			f := text.NewFile("textfile", []byte(input))
			r := text.NewReader(f)
			res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
			Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
			Expect(err).ToNot(HaveOccurred())
			node := res.(*ast.TerminalNode)
			Expect(node.Token()).To(Equal("INT"))
			Expect(node.Value(nil)).To(Equal(value))
			Expect(node.Pos()).To(Equal(nodePos))
			Expect(node.ReaderPos()).To(Equal(f.Pos(endPos)))
		},
		Entry("1 beginning", "1 ---", 0, 1, parsley.Pos(1), 1),
		Entry("1 middle", "--- 1 ---", 4, 1, parsley.Pos(5), 5),
		Entry("1 end", "--- 1", 4, 1, parsley.Pos(5), 5),
		Entry("0", "0", 0, 0, parsley.Pos(1), 1),
		Entry("+0", "+0", 0, 0, parsley.Pos(1), 2),
		Entry("-0", "-0", 0, 0, parsley.Pos(1), 2),
		Entry("1", "1", 0, 1, parsley.Pos(1), 1),
		Entry("+1", "+1", 0, 1, parsley.Pos(1), 2),
		Entry("-1", "-1", 0, -1, parsley.Pos(1), 2),
		Entry("1234567890", "1234567890", 0, 1234567890, parsley.Pos(1), 10),
		Entry("+1234567890", "+1234567890", 0, 1234567890, parsley.Pos(1), 11),
		Entry("-1234567890", "-1234567890", 0, -1234567890, parsley.Pos(1), 11),
		Entry("123abc", "123abc", 0, 123, parsley.Pos(1), 3),
		Entry("00", "00", 0, 0, parsley.Pos(1), 2), // this is a valid octal number
		Entry("01234567", "01234567", 0, 01234567, parsley.Pos(1), 8),
		Entry("+012", "+012", 0, 012, parsley.Pos(1), 4),
		Entry("-012", "-012", 0, -012, parsley.Pos(1), 4),
		Entry("08", "08", 0, 0, parsley.Pos(1), 1), // as 08 is not a valid octal number only 0 should be parsed
		Entry("0x0123456789abcdef", "0x0123456789abcdef", 0, 0x0123456789abcdef, parsley.Pos(1), 18),
		Entry("0X0123456789abcdef", "0X0123456789abcdef", 0, 0x0123456789abcdef, parsley.Pos(1), 18),
		Entry("+0x12", "+0x12", 0, 0x12, parsley.Pos(1), 5),
		Entry("-0x12", "-0x12", 0, -0x12, parsley.Pos(1), 5),
		Entry("0xg", "0xg", 0, 0, parsley.Pos(1), 1), // as 0xg is not a valid hexadecimal number only 0 should be parsed
	)

	DescribeTable("should not match",
		func(input string, startPos int) {
			f := text.NewFile("textfile", []byte(input))
			r := text.NewReader(f)
			res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
			Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(BeNil())
		},
		Entry("empty", "", 0),
		Entry("a", "a", 0),
		Entry("-", "-", 0),
		Entry("+", "+", 0),
		Entry("float 0.1", "0.1", 0),
		Entry("float 0.", "0.", 0),
	)
})
