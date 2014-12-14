//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package planner

import (
	"github.com/couchbaselabs/query/datastore"
	"github.com/couchbaselabs/query/expression"
)

type Range struct {
	Low       expression.Expressions
	High      expression.Expressions
	Inclusion datastore.Inclusion
}

type Ranges []*Range

type Span struct {
	Seek  expression.Expressions
	Range Range
}

type Spans []*Span

func SargFor(expr1, expr2 expression.Expression) Spans {
	if expr2.Value() != nil {
		return nil
	}

	s := newSarg(expr1)
	result, _ := expr2.Accept(s)
	if result != nil {
		return result.(Spans)
	}

	return nil
}

func newSarg(expr expression.Expression) expression.Visitor {
	s, _ := expr.Accept(_SARG_FACTORY)
	return s.(expression.Visitor)
}

type sargBase struct {
	sarg sargFunc
}

type sargFunc func(expression.Expression) (Spans, error)

// Arithmetic

func (this *sargBase) VisitAdd(expr *expression.Add) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitDiv(expr *expression.Div) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitMod(expr *expression.Mod) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitMult(expr *expression.Mult) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitNeg(expr *expression.Neg) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitSub(expr *expression.Sub) (interface{}, error) {
	return this.sarg(expr)
}

// Case

func (this *sargBase) VisitSearchedCase(expr *expression.SearchedCase) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitSimpleCase(expr *expression.SimpleCase) (interface{}, error) {
	return this.sarg(expr)
}

// Collection

func (this *sargBase) VisitAny(expr *expression.Any) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitArray(expr *expression.Array) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitEvery(expr *expression.Every) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitExists(expr *expression.Exists) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitFirst(expr *expression.First) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIn(expr *expression.In) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitWithin(expr *expression.Within) (interface{}, error) {
	return this.sarg(expr)
}

// Comparison

func (this *sargBase) VisitBetween(expr *expression.Between) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitEq(expr *expression.Eq) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitLE(expr *expression.LE) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitLike(expr *expression.Like) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitLT(expr *expression.LT) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsMissing(expr *expression.IsMissing) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsNotMissing(expr *expression.IsNotMissing) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsNotNull(expr *expression.IsNotNull) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsNotValued(expr *expression.IsNotValued) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsNull(expr *expression.IsNull) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitIsValued(expr *expression.IsValued) (interface{}, error) {
	return this.sarg(expr)
}

// Concat
func (this *sargBase) VisitConcat(expr *expression.Concat) (interface{}, error) {
	return this.sarg(expr)
}

// Constant
func (this *sargBase) VisitConstant(expr *expression.Constant) (interface{}, error) {
	return this.sarg(expr)
}

// Identifier
func (this *sargBase) VisitIdentifier(expr *expression.Identifier) (interface{}, error) {
	return this.sarg(expr)
}

// Construction

func (this *sargBase) VisitArrayConstruct(expr *expression.ArrayConstruct) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitObjectConstruct(expr *expression.ObjectConstruct) (interface{}, error) {
	return this.sarg(expr)
}

// Logic

func (this *sargBase) VisitAnd(expr *expression.And) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitNot(expr *expression.Not) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitOr(expr *expression.Or) (interface{}, error) {
	return this.sarg(expr)
}

// Navigation

func (this *sargBase) VisitElement(expr *expression.Element) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitField(expr *expression.Field) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitFieldName(expr *expression.FieldName) (interface{}, error) {
	return this.sarg(expr)
}

func (this *sargBase) VisitSlice(expr *expression.Slice) (interface{}, error) {
	return this.sarg(expr)
}

// Function
func (this *sargBase) VisitFunction(expr expression.Function) (interface{}, error) {
	return this.sarg(expr)
}

// Subquery
func (this *sargBase) VisitSubquery(expr expression.Subquery) (interface{}, error) {
	return this.sarg(expr)
}

// NamedParameter
func (this *sargBase) VisitNamedParameter(expr expression.NamedParameter) (interface{}, error) {
	return this.sarg(expr)
}

// PositionalParameter
func (this *sargBase) VisitPositionalParameter(expr expression.PositionalParameter) (interface{}, error) {
	return this.sarg(expr)
}
