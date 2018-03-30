// Code generated by mockery v1.0.0
package mocks

import ast "github.com/opsidian/parsley/ast"
import mock "github.com/stretchr/testify/mock"
import reader "github.com/opsidian/parsley/reader"

// Interpreter is an autogenerated mock type for the Interpreter type
type Interpreter struct {
	mock.Mock
}

// Eval provides a mock function with given fields: ctx, nodes
func (_m *Interpreter) Eval(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
	ret := _m.Called(ctx, nodes)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}, []ast.Node) interface{}); ok {
		r0 = rf(ctx, nodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 reader.Error
	if rf, ok := ret.Get(1).(func(interface{}, []ast.Node) reader.Error); ok {
		r1 = rf(ctx, nodes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(reader.Error)
		}
	}

	return r0, r1
}